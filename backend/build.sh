#!/bin/bash

docker build -t base_image .

#docker build -t login_image -f ./backend/loginService/login.Dockerfile .
#docker rm -f login_service &> /dev/null
#docker run -dit --net app_net -p 5001:5000 --rm --name login_service login_image

#docker build -t customer_image -f ./backend/customerService/customer.Dockerfile .
#docker rm -f customer_service &> /dev/null
#docker run -dit --net app_net -p 5002:5000 --rm --name customer_service customer_image

#docker build -t operator_image -f ./backend/operatorService/operator.Dockerfile .
#docker rm -f operator_service &> /dev/null
#docker run -dit --net app_net -p 5003:5000 --rm --name operator_service operator_image
