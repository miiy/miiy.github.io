---
layout: post
title: "docker 网络"
---

创建网络

```bash
docker network create my_network
```

查看网络

```bash
docker network inspect my_network
```

指定容器的网络

```bash
docker run -d --name my_container --network=my_network my_image
```