---
layout: post
title: "alpine 更换阿里源"
date: 2022-12-19
tags: linux
---

```bash
sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories
```
