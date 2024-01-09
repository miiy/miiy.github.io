---
layout: post
title: "安装 gitlab"
date: 2023-09-04
tags: gitlab
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
sudo docker run --detach \
  --hostname gitlab.example.com \
  --publish 443:443 --publish 80:80 --publish 8022:22 \
  --name gitlab \
  --restart always \
  --volume /data/gitlab/config:/etc/gitlab \
  --volume /data/gitlab/logs:/var/log/gitlab \
  --volume /data/gitlab/data:/var/opt/gitlab \
  --shm-size 256m \
  gitlab/gitlab-ee:latest
```
