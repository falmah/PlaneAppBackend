#!/bin/bash

docker build -t postgres ./postgres
docker rm -f postgres &> /dev/null
docker run -dit --rm -p 5432:5432 --name postgres postgres
