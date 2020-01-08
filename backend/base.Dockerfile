FROM ubuntu:16.04

ARG GO_VERSION=1.13.5

RUN apt-get update && apt-get install -y wget

WORKDIR /home/go

RUN wget --quiet https://dl.google.com/go/go$GO_VERSION.linux-amd64.tar.gz && \ 
    tar -C /usr/local -xf ./go$GO_VERSION.linux-amd64.tar.gz && \
    echo "export PATH=$PATH:/usr/local/go/bin" >> /etc/profile && \
    rm go$GO_VERSION.linux-amd64.tar.gz
