const port = process.env.PORT || 1500;
const domain = process.env.DOMAIN || `localhost:${port}`;

const proxy = require('redbird')({ port: port });
const Pusher = require('pusher-js');

const pusherSocket = new Pusher(process.env.PUSHER_APP_KEY, {
  forceTLS: process.env.PUSHER_APP_SECURE === '1' ? true : false,
  cluster: process.env.PUSHER_APP_CLUSTER,
});

const channel = pusherSocket.subscribe('mapped-discovery');

channel.bind('register', data => {
  proxy.register(
    `${domain}${data.prefix}`,
    `http://${data.address}:${data.port}`
  );
});

channel.bind('exit', data => {
  proxy.unregister(
    `${domain}${data.prefix}`,
    `http://${data.address}:${data.port}`
  );
});
