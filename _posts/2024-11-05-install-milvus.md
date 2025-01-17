---
layout: post
title: "install milvus"
date: 2024-11-05
---

## Install Milvus in Docker

<https://milvus.io/docs/install_standalone-docker.md>

```bash
# Download the installation script
$ curl -sfL https://raw.githubusercontent.com/milvus-io/milvus/master/scripts/standalone_embed.sh -o standalone_embed.sh

# Start the Docker container
$ bash standalone_embed.sh start
```

## Install Attu

<https://github.com/zilliztech/attu>

```bash
docker run --name attu -d \
  -p 127.0.0.1:3000:3000 \
  -e MILVUS_URL={milvus server IP}:19530 \
  zilliz/attu:v2.4
```
