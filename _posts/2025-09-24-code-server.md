---
layout: post
title: "code server"
date: 2025-09-24
tags: linux
---

# code-server

```bash
#!/bin/sh
# This will start a code-server container and expose it at http://127.0.0.1:8080.
# It will also mount your current directory into the container as `/home/coder/project`
# and forward your UID/GID so that all file system operations occur as your user outside
# the container.
#
# Your $HOME/.config is mounted at $HOME/.config within the container to ensure you can
# easily access/modify your code-server config in $HOME/.config/code-server/config.json
# outside the container.
docker run -it --name code-server -d --restart=always \
  --user "$(id -u)" \
  -v "/srv/docker/code-server/.config:/home/coder/.config" \
  -v "/srv/docker/code-server/.ssh:/home/coder/.ssh" \
  -v "$HOME/project:/home/coder/project" \
  --net frontend \
  codercom/code-server:4.104.1
```

进入 web 端

安装常用工具

```bash
sudo apt update
sudo apt install iputils-ping
```

设置 dns

```bash
sudo sh -c 'echo "nameserver 8.8.8.8" > /etc/resolv.conf'
```