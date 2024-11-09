#!/bin/sh

git pull
docker stop $(docker ps -a -q)
docker-compose up -d --build --force-recreate --remove-orphans