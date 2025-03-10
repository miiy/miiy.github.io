---
layout: post
title: "安装 scrapyd"
date: 2025-01-23
---

<https://scrapyd.readthedocs.io/en/latest/deploy.html>

```bash
sudo docker run --name scrapyd -d --restart=unless-stopped \
    -v /data/scrapyd/config/:/etc/scrapyd/ \
    --network backend \
    registry.cn-hangzhou.aliyuncs.com/appx/scrapyd:latest
```