---
layout: post
title: "禁用 swap 交换分区"
date: 2020-12-12
tags: linux
---

永久禁用 swap

```bash
debian@debian:~# sudo sed -ri 's/.*swap.*/#&/' /etc/fstab
debian@debian:~# sudo reboot
debian@debian:~# free -mh
```

临时关闭 swap，启用 swap

```bash
debian@debian:~# free -mh
              total        used        free      shared  buff/cache   available
Mem:          987Mi        64Mi       831Mi       2.0Mi        92Mi       806Mi
Swap:         974Mi          0B       974Mi
debian@debian:~# sudo swapoff -a
debian@debian:~# free -mh
              total        used        free      shared  buff/cache   available
Mem:          987Mi        64Mi       831Mi       2.0Mi        92Mi       805Mi
Swap:            0B          0B          0B
debian@debian:~# sudo swapon -a
debian@debian:~# free -mh
              total        used        free      shared  buff/cache   available
Mem:          987Mi        64Mi       831Mi       2.0Mi        92Mi       806Mi
Swap:         974Mi          0B       974Mi
```
