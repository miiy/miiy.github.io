---
layout: post
title: "MacOS 如何查看硬盘健康度"
date: 2024-04-08
tags: mac
---

## 安装

下载 smartmontools

<https://www.smartmontools.org/wiki/Download>

直接下载dmg 或者编译

```bash
tar -zxvf smartmontools-7.4.tar.gz
cd smartmontools-7.4
./configure
make
```

## 使用

查看硬盘

```bash
diskutil list
```

查看硬盘状态

```bash
./smartctl -a disk0
```
