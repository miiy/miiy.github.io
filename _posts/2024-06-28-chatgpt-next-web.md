---
layout: post
title: "使用 ChatGPT Next Web 部署私人 ChatGPT 网页应用"
date: 2024-06-28
tags: AI
---

ChatGPT Next Web: <https://github.com/ChatGPTNextWeb/ChatGPT-Next-Web>

```bash
docker run --name chatgpt-next-web -d --restart=always \
    --log-opt max-size=1g \
    --net frontend \
    yidadaa/chatgpt-next-web:v2.16.0
```

端口

```
    -p 3000:3000 \
```

代理：

```
    --add-host=host.docker.internal:host-gateway \
    -e PROXY_URL=http://172.17.0.1:7890 \
```

nginx stream: <https://github.com/Yidadaa/ChatGPT-Next-Web/blob/main/docs/faq-en.md>
