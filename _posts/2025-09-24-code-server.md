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
docker run --privileged -it --name code-server -d --restart=always \
  --user "$(id -u)" \
  -v "/srv/docker/code-server/.config:/home/coder/.config" \
  -v "/srv/docker/code-server/.ssh:/home/coder/.ssh" \
  -v "$HOME/project:/home/coder/project" \
  -v "/srv/docker/code-server/workspace:/home/coder/workspace" \
  --net frontend \
  codercom/code-server:4.127.0
```

进入 web 端

init env

```bash
sudo sed -i 's/deb.debian.org/mirrors.ustc.edu.cn/g' /etc/apt/sources.list.d/debian.sources \
&& sudo sed -i 's/security.debian.org/mirrors.ustc.edu.cn/g' /etc/apt/sources.list.d/debian.sources \
&& cat /etc/apt/sources.list.d/debian.sources \
&& sudo apt update \
&& sudo apt install -y vim zip iputils-ping python3 pip python3-venv jq chromium \
&& curl --proto '=https' --tlsv1.2 -sSf https://sh.rustup.rs | sh \
&& echo 'export PATH="$HOME/workspace/dev/node/bin:$PATH"' >> ~/.bashrc \
&& echo 'export PATH="$HOME/workspace/dev/go/bin:$PATH"' >> ~/.bashrc \
&& source ~/.bashrc \
&& npm install -g @openai/codex
```

设置 dns

```bash
sudo sh -c 'echo "nameserver 8.8.8.8" > /etc/resolv.conf'
```