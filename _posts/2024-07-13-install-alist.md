---
layout: post
title: "安装 alist"
date: 2024-07-13
tags: linux
---



## Install

```bash
docker run -d --name="alist" --restart=unless-stopped \
  -e PUID="$(id -u)" \
  -e PGID="$(id -g)" \
  -v /home/debian/data/alist/data:/opt/alist/data \
  -v /home/debian/Public:/media/Public \
  -p 5244:5244 \
  -e PUID=0 -e PGID=0 -e UMASK=022 \
  xhofe/alist:v3.35.0-ffmpeg
```

设置密码

```bash
docker exec -it alist ./alist admin set NEW_PASSWORD
```