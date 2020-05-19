#!/bin/bash

docker network rm app_net &> /dev/null
docker network create app_net 

docker build --build-arg "DOCKER_HOST=$(ip -4 addr show docker0 | grep -Po 'inet \K[\d.]+')" -t postgres ./postgres 
docker rm -f postgres &> /dev/null
docker run -dit --net app_net --rm -p 5432:5432 --name postgres postgres

docker build --build-arg "SOME_HOST=$(ip -4 addr show docker0 | grep -Po 'inet \K[\d.]+')" -t proxy  ./nginx
docker rm -f proxy
docker run -dit --net app_net -p 81:80 --name proxy proxy

source ./backend/build.sh
