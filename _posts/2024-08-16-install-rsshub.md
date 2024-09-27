---
layout: post
title: "安装 RSSHub"
date: 2024-08-16
---

## Run

```bash
docker run -d --restart=unless-stopped --name rsshub \
  --add-host=host.docker.internal:host-gateway \
  -e "PROXY_URI=http://host.docker.internal:7890" \
  --network frontend \
  diygod/rsshub:chromium-bundled-2024-08-15
```

```txt
 -p 1200:1200
```

安装好后配合 RSS 阅读器订阅

Fluent Reader: <https://github.com/yang991178/fluent-reader>
