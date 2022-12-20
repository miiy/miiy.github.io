---
layout: post
title: "debian11 更换阿里源"
date: 2022-12-19
tags: linux
---

```bash
sed -i "s/deb.debian.org/mirrors.aliyun.com/g" /etc/apt/sources.list
sed -i "s/security.debian.org/mirrors.aliyun.com/g" /etc/apt/sources.list
cat /etc/apt/sources.list
apt-get update
```
