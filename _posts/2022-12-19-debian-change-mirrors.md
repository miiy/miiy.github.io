---
layout: post
title: "debian 更换国内源"
date: 2022-12-19
tags: linux
---

debian12

```bash
sed -i 's/deb.debian.org/mirrors.ustc.edu.cn/g' /etc/apt/sources.list.d/debian.sources \
&& sed -i 's/security.debian.org/mirrors.ustc.edu.cn/g' /etc/apt/sources.list.d/debian.sources \
&& cat /etc/apt/sources.list.d/debian.sources \
&& apt update
```

debian11

```bash
sed -i "s/deb.debian.org/mirrors.aliyun.com/g" /etc/apt/sources.list
sed -i "s/security.debian.org/mirrors.aliyun.com/g" /etc/apt/sources.list
cat /etc/apt/sources.list
apt-get update
```
