---
layout: post
title: "Docker Compose"
---

## docker-compose

<https://docs.docker.com/compose/>

```bash
# Start services
docker-compose up -d nginx php-fpm mysql redis
# Stop all service
docker-compose stop
# build
docker-compose build php-fpm
# Lists containers
docker-compose ps
# mysql
docker-compose run --rm mysql bash
# php-cli
docker-compose run --rm php-cli bash
# redis
docker-compose run --rm redis redis-cli -h redis
# node
docker-compose run --rm node sh
# bind port
docker-compose run --rm -p 3001:3001 node sh
# golang
docker-compose run --rm golang bash
```
