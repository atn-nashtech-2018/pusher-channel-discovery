const express = require('express');
const bodyParser = require('body-parser');
const os = require('os');
const uuidv5 = require('uuid/v5');
const uuidv4 = require('uuid/v4');
const Pusher = require('pusher');
const internalIp = require('internal-ip');

const app = express();
const hostName = os.hostname();
const port = process.env.PORT || 3000;

let pusher = new Pusher({
  appId: process.env.PUSHER_APP_ID,
  key: process.env.PUSHER_APP_KEY,
  secret: process.env.PUSHER_APP_SECRET,
  encrypted: process.env.PUSHER_APP_SECURE,
  cluster: process.env.PUSHER_APP_CLUSTER,
});

let svc = {};

internalIp
  .v4()
  .then(ip => {
    svc = {
      prefix: '/v1',
      id: uuidv4(),
      name: 'Unique ID generator',
      host: hostName,
      port: port,
      address: ip,
      health: {
        endpoint: `/health`, // This would be appended to the ip and port address.
        method: 'GET',
      },
    };

    console.log('Registering service');

    pusher.trigger('mapped-discovery', 'register', svc);
  })
  .catch(err => {
    console.log(err);
    process.exit();
  });

process.stdin.resume();

process.on('SIGINT', () => {
  console.log('Deregistering service... ');

  // Send an exit signal on shutdown
  pusher.trigger('mapped-discovery', 'exit', svc);

  // Timeout to make sure the signal sent to
  // Pusher was successful before shutting down
  setTimeout(() => {
    process.exit();
  }, 1000);
});

app.use(bodyParser.json());

app.use(function(req, res, next) {
  res.header('X-Server', hostName);
  next();
});

app.get('/', function(req, res) {
  res.status(200).send({ service: 'ID generator' });
});

app.get('/health', function(req, res) {
  res.status(200).send({ status: 'ok' });
});

app.post('/generate', function(req, res) {
  const identifier = req.body.id;

  if (identifier === undefined) {
    res.status(400).send({
      message: 'Please provide an ID to use to generate your UUID V5',
    });
    return;
  }

  if (identifier.length === 0) {
    res.status(400).send({
      message: 'Please provide an ID to use to generate your UUID V5',
    });
    return;
  }

  res.status(200).send({
    id: uuidv5(identifier, uuidv5.URL),
    timestamp: new Date().getTime(),
    message: 'UUID was successfully generated',
  });
});

app.listen(port, function() {
  console.log(`Service is running at ${port} at ${hostName}`);
});
