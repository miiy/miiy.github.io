---
layout: post
title: "Redis"
date: 2016-06-13
tags: redis
---

## Install

```bash
cd /usr/local/src
wget http://download.redis.io/releases/redis-4.0.2.tar.gz
tar -zxvf redis-4.0.2.tar.gz
cd redis-4.0.2
make MALLOC=libc
make install
```

## 开机启动

```bash
cd /usr/local/src/redis-4.0.2
mkdir /etc/redis
cp redis.conf /etc/redis/6379.conf
sed -i 's/^daemonize no/daemonize yes/' /etc/redis/6379.conf
sed -i 's/^dir .\//dir \/var\/lib\/redis/' /etc/redis/6379.conf
sed -i 's/^appendonly no/appendonly yes/' /etc/redis/6379.conf
mkdir -p /var/lib/redis
cp utils/redis_init_script /etc/rc.d/init.d/redisd
chmod +x /etc/rc.d/init.d/redisd
sed -i '5i# chkconfig: 2345 90 10\n# description: Redis is a persistent key-value database' /etc/init.d/redisd
chkconfig redisd on
```

dir /var/redis/6379 设置持久化文件的存储位置
