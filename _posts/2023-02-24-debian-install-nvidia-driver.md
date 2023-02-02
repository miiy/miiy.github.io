---
layout: post
title: "debian 安装 Nvidia 显卡驱动"
date: 2023-02-24
tags: linux
---

添加 non-free 源

```bash
vi /etc/apt/sources.list
```

```conf
deb http://mirrors.163.com/debian/ bullseye main contrib non-free
deb-src http://mirrors.163.com/debian/ bullseye main contrib non-free
```

安装Nvidia 驱动

```bash
sudo apt-get install firmware-misc-nonfree nvidia-detect  nvidia-driver  nvidia-settings  nvidia-opencl-dev
```

