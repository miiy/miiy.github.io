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

<https://docs.gitlab.com/ee/install/docker/configuration.html>

<https://docs.gitlab.com/omnibus/settings/ssl/index.html>

```bash
GITLAB_HOME=/data/gitlab
sudo docker run --detach \
  --hostname gitlab.example.com \
  # --publish 443:443 --publish 80:80 --publish 2022:22 \
  --publish 2022:22 \
  --name gitlab \
  --restart always \
  --volume $GITLAB_HOME/config:/etc/gitlab \
  --volume $GITLAB_HOME/logs:/var/log/gitlab \
  --volume $GITLAB_HOME/data:/var/opt/gitlab \
  --shm-size 256m \
  gitlab/gitlab-ce:17.7.0-ce.0
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

```

修改配置文件

```bash
vi /data/gitlab/config/gitlab.rb
```

```text
external_url 'https://gitlab.example.com'
letsencrypt['enable'] = false
nginx['redirect_http_to_https'] = true
gitlab_rails['gitlab_shell_ssh_port'] = 2022
# 禁用 Puma 集群模式，减少内存占用
puma['worker_processes'] = 0
# 设置 Sidekiq 进程数量
sidekiq['concurrency'] = 10
```

```bash
docker restart gitlab
```

## Security

禁用注册

/help, /explore 暴漏项目，建议所有项目私有