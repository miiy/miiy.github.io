---
layout: post
title: "安装 Syncthing 文件同步程序"
date: 2024-07-13
tags: linux
---

Syncthing 是一个连续的文件同步程序。它在两台或多台计算机之间同步文件。

Syncthing: <https://github.com/syncthing/syncthing>

## Install

```bash
docker run --name syncthing -d --restart=unless-stopped \
    -p 22000:22000/tcp -p 22000:22000/udp -p 21027:21027/udp \
    -v /home/debian/data/syncthing:/var/syncthing \
    --network frontend \
    syncthing/syncthing:1.29
```

port

```txt
8384:8384 # Web UI
22000:22000/tcp # TCP file transfers
22000:22000/udp # QUIC file transfers
21027:21027/udp # Receive local discovery broadcasts
```
