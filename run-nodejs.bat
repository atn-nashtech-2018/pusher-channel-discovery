@echo off
call SET AppId=802103
call SET AppKey=180a327c9905465f99ca
call SET SecretKey=f777847385b993fc346d
call SET AppCluster=ap1
call SET ProjectRootFolder=C:\Users\tunguyena1\Documents\Projects\pusher-channel-discovery

@echo on
call cd /d "%ProjectRootFolder%"
call cd nodejs
call npm i event-stream
call npm install
call docker build -t pusher-channel-node .
call docker run -p 127.0.0.1:3000:3000 -e PUSHER_APP_ID=%AppId% -e PUSHER_APP_KEY=%AppKey% -e PUSHER_APP_SECRET="%SecretKey%" -e PUSHER_APP_CLUSTER="%AppCluster%" -e PUSHER_APP_SECURE="1" pusher-channel-node
