#!/bin/bash
docker rm $(docker ps --filter status=exited -q)
docker network prune
docker compose build
docker compose up