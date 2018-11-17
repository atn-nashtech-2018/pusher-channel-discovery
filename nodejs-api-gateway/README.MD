## An api gateway using Pusher channels for service discovery.

Using pusher channels for both service registry and discovery. This repo
contains an implementation of an api gateway that uses Pusher channels as a backend for
service registration and discovery.

When services register themselves, they are added to the list of backends for a
given route, after which requests are loadbalanced amogst them.

So how do services register themselves ?
- Golang example https://github.com/adelowo/pusher-channel-discovery-go
- NodeJS example https://github.com/adelowo/pusher-channel-discovery-node

## Pre-requisites

- Node `>= 10`
- Docker

## Built With

- [Pusher Channels](https://pusher.com/channels) - APIs to enable devs building realtime features.
- [Redbird](https://github.com/OptimalBits/redbird) - A modern reverse proxy for node.

