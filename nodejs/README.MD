## Service discovery with Pusher channels

Using pusher channels for both service registry and discovery. This repo
contains an implementation of an HTTP server that uses Pusher channels to notify
others of it's state. This is a node implementation, you can find a Golang
implementation at https://github.com/adelowo/pusher-channel-discovery

An api gateway that uses the information published by this can be found at
https://github.com/adelowo/pusher-channels-api-gateway

## Pre-requisites

- Node `>= 10`

## Built With

- [Pusher Channels](https://pusher.com/channels) - APIs to enable devs building realtime features.


