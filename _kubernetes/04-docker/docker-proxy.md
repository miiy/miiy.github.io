---
layout: post
title: "docker 配置 http 代理"
---

方法1

```bash
sudo vi /etc/docker/daemon.json
```

```json
  "proxies": {
    "http-proxy": "http://127.0.0.1:37890",
    "https-proxy": "http://127.0.0.1:37890",
    "no-proxy": "*.test.example.com,.example.org,127.0.0.0/8"
  }
```

方法2 创建 http proxy service

```bash
sudo vim /etc/systemd/system/docker.service.d/http-proxy.conf
```

```conf
[Service]
Environment="HTTP_PROXY=172.1.1.1:7890"
Environment="HTTPS_PROXY=172.1.1.1:7890"
Environment="NO_PROXY=localhost,127.0.0.1,::1,10.0.0.0/8,.youdomain.com"
```

重启服务

```bash
sudo systemctl daemon-reload
sudo systemctl restart docker
sudo docker info
```

build 时使用代理

```bash
docker build -f Dockerfile_API --build-arg http_proxy="http://127.0.0.1:7890" \
--build-arg https_proxy="http://127.0.0.1:7890" \ -t apiserver:v0.0.1 .
```
