---
layout: post
title: "安装 Calibre-Web"
date: 2024-07-13
tags: linux
---

<https://github.com/janeczku/calibre-web>

<https://github.com/linuxserver/docker-calibre-web>

# Install

```bash
docker run -d \
  --name=calibre-web \
  -e PUID="$(id -u)" \
  -e PGID="$(id -g)" \
  -e TZ=Etc/UTC \
  -e DOCKER_MODS=linuxserver/mods:universal-calibre `#optional` \
  -e OAUTHLIB_RELAX_TOKEN_SCOPE=1 `#optional` \
  -v /home/debian/data/www/calibre-web/config:/config \
  -v /home/debian/Calibre\ Library:/books \
  --restart unless-stopped \
  --network frontend \
  linuxserver/calibre-web:0.6.22
```

port

```txt
  -p 8083:8083 \
```
