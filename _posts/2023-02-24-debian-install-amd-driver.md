---
layout: post
title: "debian 安装 AMD 显卡驱动"
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

安装 AMD 显卡驱动

```bash
sudo apt-get install firmware-amd-graphics libgl1-mesa-dri libglx-mesa0 mesa-vulkan-drivers xserver-xorg-video-all mesa-opencl-icd
```

查看显卡使用情况

```bash
sudo apt install radeontop
```

移除显卡驱动

```bash
sudo apt-get install firmware-amd-graphics
```
