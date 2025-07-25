---
layout: post
title: "配置 git 服务仓库"
date: 2023-02-24
tags: linux
---

服务端：

```bash
$ groupadd git
$ useradd git -M -g git -s /sbin/nologin
$ mkdir -p /data/git
$ cd /data/git
$ git init --bare testrepo.git
```

useradd -s 指定用户登录后使用的shell

useradd -M not create home 不创建home目录

客户端：

```bash
$ git clone root@aliyun:/data/git/testrepo.git
```
