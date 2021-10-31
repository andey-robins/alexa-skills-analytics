#!/bin/bash
docker rm $(docker ps --filter status=exited -q)
docker compose build
docker compose up