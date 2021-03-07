#!/usr/bin/env bash
git pull origin master
docker-compose up --build -d --force-recreate