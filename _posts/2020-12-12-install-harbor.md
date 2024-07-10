---
layout: post
title: "安装 Harbor"
date: 2020-12-12
tags: kubernetes
---

https://goharbor.io/

## Install

安装 docker, docker-compose

```bash
cd /usr/local/src
wget harbor-online-installer-v2.2.0.tgz
tar -zxvf harbor-online-installer-v2.2.0.tgz
mv harbor harbor-v2.2.0
mv harbor-v2.2.0 ../
cd ../
ln -s harbor-v2.2.0 harbor
```

生成证书

https://goharbor.io/docs/2.2.0/install-config/configure-https/

```bash

mkdir -p /data/harbor/certs
cd /data/harbor/harbor

openssl genrsa -out harbor.example.com.key 4096

openssl req -sha512 -new \
    -subj "/C=CN/ST=Beijing/L=Beijing/O=exampleOrg/CN=harbor.example.com" \
    -key harbor.example.com.key \
    -out harbor.example.com.csr

openssl x509 -req  -days 3650 \
    -in harbor.example.com.csr \
    -signkey harbor.example.com.key \
    -out harbor.example.com.crt
```

修改配置文件

```bash
cd /usr/local/harbor
cp harbor.yaml.tmpl harbor.yaml
vi harbor.yaml
```

```text
hostname: harbor.example.com

certificate: /data/harbor/certs/harbor.example.com.crt
private_key: /data/harbor/certs/harbor.example.com.key

data_volume: /data/harbor/data
```

安装

```bash
./install.sh
```

启动停止

``bash
cd /usr/local/harbor
docker-compose up -d
docker-compose stop
``

## 使用

```bash
vi /etc/hosts
```

```text
10.0.2.20       harbor.example.com
```

```bash
vi /etc/docker/daemon.json
{
    "registry-mirrors": ["http://hub-mirror.c.163.com"],
    "insecure-registries": ["https://harbor.example.com"]
}

systemctl restart docker

docker login -u admin harbor.example.com
```

在安装了 docker-compose 的机器上登录会报错，卸载 docker-compose 登录后在重新安装 docker-compose， 或使用其他机器登录

```text
Error saving credentials: error storing credentials - err: exit status 1, out: `Cannot autolaunch D-Bus without X11 $DISPLAY`
```

```bash
root@debian:~# docker tag nginx:latest harbor.example.com/library/nginx:v1
root@debian:~# docker images
REPOSITORY                                           TAG           IMAGE ID       CREATED         SIZE
nginx                                                latest        6084105296a9   3 hours ago     133MB
harbor.example.com/library/nginx                             v1            6084105296a9   3 hours ago     133MB
root@debian:~# docker push harbor.example.com/library/nginx:v1
```

## MacOS 下安装

docker 设置共享目录 /usr/local 并重启 docker

```text
修改 prepare
-v /:/hostfs \
换成
-v /usr/local:/hostfs/usr/local \
```

mkdir -p common/config
