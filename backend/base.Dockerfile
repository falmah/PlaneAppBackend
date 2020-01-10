FROM ubuntu:16.04

ARG GO_VERSION=1.13.5
ENV GOPATH=/home/go
ENV GOROOT=/usr/local/go

RUN apt-get update && apt-get install -y wget

WORKDIR $GOPATH/src

RUN wget --quiet https://dl.google.com/go/go$GO_VERSION.linux-amd64.tar.gz && \ 
    tar -C /usr/local -xf ./go$GO_VERSION.linux-amd64.tar.gz && \
    echo "export PATH=$PATH:$GOROOT/bin" >> /etc/profile && \
    rm go$GO_VERSION.linux-amd64.tar.gz
