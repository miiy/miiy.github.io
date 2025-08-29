---
layout: post
title: "Docker 安装"
---


## Install

<https://docs.docker.com/engine/install>

阿里云 docker-ce 镜像：<https://developer.aliyun.com/mirror/docker-ce>

### Debian

```bash
sudo apt-get update
sudo apt-get install \
    apt-transport-https \
    ca-certificates \
    curl \
    gnupg-agent \
    software-properties-common

curl -fsSL https://download.docker.com/linux/debian/gpg | sudo apt-key add -
apt-key fingerprint 0EBFCD88
sudo add-apt-repository \
   "deb [arch=amd64] https://download.docker.com/linux/debian \
   $(lsb_release -cs) \
   stable"
sudo apt-get update
sudo apt-get install docker-ce docker-ce-cli containerd.io

cat <<EOF | sudo tee /etc/docker/daemon.json
{
    "registry-mirrors": ["http://hub-mirror.c.163.com"]
}
EOF
```

### 迁移 docker 数据目录

/etc/docker/daemon.json

```conf
{
  "data-root": "/mnt/data/docker"
}
```

```bash
sudo systemctl stop docker
sudo systemctl stop docker.socket
sudo rsync -aP /var/lib/docker/ /mnt/data/docker/

sudo mv /var/lib/docker /var/lib/docker.old

sudo systemctl start docker

sudo systemctl start docker
sudo docker info
```
