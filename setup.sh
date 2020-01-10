#!/bin/bash

docker network rm test_net &> /dev/null
docker network create test_net

docker build --build-arg "DOCKER_HOST=$(ip -4 addr show docker0 | grep -Po 'inet \K[\d.]+')" -t postgres ./postgres 
docker rm -f postgres &> /dev/null
docker run -dit --net test_net --rm -p 5432:5432 --name postgres postgres
