---
layout: post
title: "install redis"
date: 2016-06-13
tags: linux
---

# Install

<https://docs.oracle.com/en/java/javase/15/install/installation-jdk-linux-platforms.html>

```bash
cd /usr/local/src
wget https://download.redis.io/releases/redis-6.2.0.tar.gz
tar -zxvf redis-6.2.0.tar.gz redis-6.2.0/
mv redis-6.2.0 /usr/local/
cd /usr/local/
ln -s redis-6.2.0/ redis
make
make PREFIX=/usr/local/redis install
```
