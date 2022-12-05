---
layout: post
title: "hostname"
---

```bash
debian@debian:~# hostname
debian
debian@debian:~# sudo hostnamectl set-hostname server01
debian@debian:~#cat /etc/hostname
server01
debian@debian:~# sudo vi /etc/hosts
```

```text
127.0.0.1 server01
```
