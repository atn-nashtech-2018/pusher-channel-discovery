## Service discovery with Pusher channels

Using pusher channels for both service registry and discovery. This repo
contains an implementation of an HTTP server ( Golang and NodeJS implementations) 
that uses Pusher channels to notify others of it's state.

There is also an API gateway implementation in NodeJS that listens to the
messages published by the aforementioned HTTP servers in other to set up a
reverse proxy to them.

The tutorial can be found at https://pusher.com/tutorials/service-discovery-microservices

## Pre-requisites

- Node.
- Golang `>=1.11`
- Docker


## Getting started

You will have to clone this project to get started. That can be done by 
`git clone git@github.com:adelowo/pusher-channel-discovery.git`

- Api Gateway:
  - `cd nodejs-api-gateway`
  - `docker build -t pusher-channel-api-gateway .`
  - `docker run -p 127.0.0.1:1500:1500 -e PUSHER_APP_KEY="XXXXX" -e PUSHER_APP_SECURE="1" -e PUSHER_APP_CLUSTER="eu" pusher-channel-api-gateway`
  
- NodeJS HTTP server:
  - `cd nodejs`
  - `docker build -t pusher-channel-node .`
  - `docker run -p 127.0.0.1:3000:3000 -e PUSHER_APP_ID="XXXX" -e PUSHER_APP_KEY="XXXXXXX" -e PUSHER_APP_SECRET="XXXXXX" -e PUSHER_APP_CLUSTER="eu" -e PUSHER_APP_SECURE="1" pusher-channel-node`
 
- Golang HTTP server:
  - `cd golang`
  - `docker build -t pusher-channel-go .`
  - `docker run -p 127.0.0.1:2000:3000 -e PUSHER_APP_ID="XXXX" -e PUSHER_APP_KEY="XXXXX" -e PUSHER_APP_SECRET="XXXX" -e PUSHER_APP_CLUSTER="eu" -e PUSHER_APP_SECURE="1" pusher-channel-go`

## Built With

- [Pusher Channels](https://pusher.com/channels) - APIs to enable devs building realtime features.
- [Redbird](https://github.com/OptimalBits/redbird) - A modern reverse proxy for node.

