#!/bin/bash
docker build -t base_image ./backend
docker-compose up --build