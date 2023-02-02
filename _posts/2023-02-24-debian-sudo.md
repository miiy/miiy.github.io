---
layout: post
title: "debian sudo"
date: 2023-02-24
tags: linux
---

方法1：为用户单独添加 sudo 权限

```bash
echo "debian        ALL=(ALL:ALL) ALL" > /etc/sudoers.d/debian
```

方法2：将用加入 sudo 用户组

```bash
root@debian:~# id debian
uid=1000(debian) gid=1000(debian) groups=1000(debian),24(cdrom),25(floppy),29(audio),30(dip),44(video),46(plugdev),109(netdev),111(bluetooth)
root@debian:~# usermod -a -G sudo debian
root@k8s-master:~# id debian
uid=1000(debian) gid=1000(debian) groups=1000(debian),24(cdrom),25(floppy),27(sudo),29(audio),30(dip),44(video),46(plugdev),109(netdev),111(bluetooth)
```
