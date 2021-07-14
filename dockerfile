FROM golang:1.16.6-alpine

RUN apk add curl \
    && apk add git mercurial

RUN curl -fsSL -o /usr/local/bin/dep https://github.com/golang/dep/releases/download/v0.5.0/dep-linux-amd64 && chmod +x /usr/local/bin/dep

WORKDIR /go/src/github.com/xabi93/pig-translator

ADD . /go/src/github.com/xabi93/pig-translator

RUN dep ensure