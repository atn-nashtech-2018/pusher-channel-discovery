@echo off
call SET AppId=802103
call SET AppKey=180a327c9905465f99ca
call SET SecretKey=f777847385b993fc346d
call SET AppCluster=ap1
call SET ProjectRootFolder=C:\Users\tunguyena1\Documents\Projects\pusher-channel-discovery

@echo on
call cd /d "%ProjectRootFolder%"
call cd nodejs-api-gateway
call docker build -t pusher-channel-api-gateway .
call docker run -p 127.0.0.1:1500:1500 -e PUSHER_APP_KEY="%AppKey%" -e PUSHER_APP_SECURE="1" -e PUSHER_APP_CLUSTER="%AppCluster%" pusher-channel-api-gateway
