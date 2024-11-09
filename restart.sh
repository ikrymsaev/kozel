#!/bin/sh

docker stop $(docker ps -a -q) \
docker-compose up --build --force-recreate -d