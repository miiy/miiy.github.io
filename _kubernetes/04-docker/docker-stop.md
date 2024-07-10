---
layout: post
title: "停止 docker"
---

```bash
sudo systemctl stop docker
sudo systemctl status docker
```

在停止 docker 之后，查看状态发现 docker 又重新启动了。

Docker 使用了一个守护进程 (Docker Daemon)，它监听来自客户端的指令，并管理容器的生命周期。而 docker.socket 则是控制 Docker 守护进程的套接字文件。

正确的关闭 docker 服务

```bash
sudo systemctl stop docker.socket
sudo systemctl stop docker.service
sudo systemctl status docker
```
