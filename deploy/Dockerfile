FROM golang:alpine as builder
MAINTAINER Jack Murdock <jack_murdock@comcast.com>

WORKDIR /go/src/github.com/xmidt-org/themis

RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh

COPY . .
RUN GO111MODULE=on go build -o themis_linux_amd64

FROM alpine

RUN apk --no-cache add ca-certificates
RUN mkdir -p /etc/themis
VOLUME /etc/themis

EXPOSE 8080

COPY --from=builder /go/src/github.com/xmidt-org/themis/themis_linux_amd64 /
COPY themis.yaml /
ENTRYPOINT ["/themis_linux_amd64"]