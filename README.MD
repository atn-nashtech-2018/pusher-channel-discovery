## Service discovery with Pusher channels

Using pusher channels for both service registry and discovery. This repo
contains an implementation of an HTTP server ( Golang and NodeJS implementations) that uses Pusher channels to notify
others of it's state.

There is also an API gateway implementation in NodeJS that listens to the
messages published by the aforementioned HTTP servers in other to set up a
reverse proxy to them.

## Pre-requisites

- Node.
- Golang `>=1.11`
- Docker

## Built With

- [Pusher Channels](https://pusher.com/channels) - APIs to enable devs building realtime features.
- [Redbird](https://github.com/OptimalBits/redbird) - A modern reverse proxy for node.

