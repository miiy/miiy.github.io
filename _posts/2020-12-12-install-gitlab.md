---
layout: post
title: "安装 gitlab"
date: 2020-12-12
tags: linux
---

## Requirements

<https://docs.gitlab.com/ee/install/requirements.html>

CPU:

- 4 cores is the recommended minimum number of cores and supports up to 500 users
- 8 cores supports up to 1000 users

Memory:

- 4 GB RAM is the required minimum memory size and supports up to 500 users
- 8 GB RAM supports up to 1000 users

推荐 4核，4G


## Install Gitlab using Docker Engine

<https://docs.gitlab.com/ee/install/docker.html>

```bash
GITLAB_HOME=/data/gitlab
sudo docker run --detach \
  --hostname gitlab.example.com \
  --add-host jenkins.example.com:10.0.2.20 \
  --add-host harbor.example.com:10.0.2.20 \
  --publish 443:443 --publish 80:80 --publish 8022:22 \
  --name gitlab \
  --restart always \
  --volume $GITLAB_HOME/config:/etc/gitlab \
  --volume $GITLAB_HOME/logs:/var/log/gitlab \
  --volume $GITLAB_HOME/data:/var/opt/gitlab \
  --shm-size 256m \
  gitlab/gitlab-ee:latest
```

## 启用 https

```bash
mkdir -p /data/gitlab/config/ssl
cd /data/gitlab/config/ssl

openssl genrsa -out gitlab.example.com.key 4096

openssl req -sha512 -new \
-subj "/C=CN/ST=Beijing/L=Beijing/O=exampleOrg/CN=gitlab.example.com" \
-key gitlab.example.com.key \
-out gitlab.example.com.csr

openssl x509 -req  -days 3650 \
-in gitlab.example.com.csr \
-signkey gitlab.example.com.key \
-out gitlab.example.com.crt

vi /data/gitlab/config/gitlab.rb
```

```text
external_url 'https://gitlab.example.com'
nginx['redirect_http_to_https'] = true
```

```bash
docker restart gitlab
```
