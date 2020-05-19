FROM ubuntu:16.04

ARG GO_VERSION=1.14
ENV GOPATH=/home/go
ENV GOROOT=/usr/local/go

RUN apt-get update && apt-get install -y wget git

WORKDIR $GOPATH/src

COPY ./model/* ./model/

RUN wget --quiet https://dl.google.com/go/go$GO_VERSION.linux-amd64.tar.gz && \ 
    tar -C /usr/local -xf ./go$GO_VERSION.linux-amd64.tar.gz && \
    echo "export PATH=$PATH:$GOROOT/bin" >> /etc/profile && \
    rm go$GO_VERSION.linux-amd64.tar.gz

RUN ${GOROOT}/bin/go get -u github.com/jinzhu/gorm && \
    ${GOROOT}/bin/go get -u github.com/gorilla/mux && \
    ${GOROOT}/bin/go get -u github.com/sirupsen/logrus && \
    ${GOROOT}/bin/go get -u github.com/lib/pq

EXPOSE 5000
