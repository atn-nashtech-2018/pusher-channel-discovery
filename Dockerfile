FROM golang:1.11 as build-env

WORKDIR /go/src/github.com/adelowo/pusher-channel-discovery-go
ADD . /go/src/github.com/adelowo/pusher-channel-discovery-go

ENV GO111MODULE=on

RUN go mod download
RUN go mod verify
RUN go install ./cmd

## A better scratch
FROM gcr.io/distroless/base
COPY --from=build-env /go/bin/cmd /
CMD ["/cmd"]
