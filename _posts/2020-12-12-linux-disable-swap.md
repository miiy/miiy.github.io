---
layout: post
title: "禁用 swap 交换分区"
date: 2020-12-12
tags: linux
---

永久禁用 swap

```bash
$ sudo sed -ri 's/.*swap.*/#&/' /etc/fstab
$ sudo reboot
$ free -mh

# 如果未起作用，执行下面的命令
# systemd 会根据 分区类型 自动生成 .swap unit
$ systemctl list-units --type=swap
  UNIT                                   LOAD   ACTIVE SUB    DESCRIPTION
  dev-nvme0n1p3.swap                     loaded active active Swap Partition
# mask 掉 systemd 生成的 swap unit
$ sudo systemctl mask dev-nvme0n1p3.swap
$ sudo reboot
$ free -mh
```

临时关闭 swap，启用 swap

```bash
$ free -mh
              total        used        free      shared  buff/cache   available
Mem:          987Mi        64Mi       831Mi       2.0Mi        92Mi       806Mi
Swap:         974Mi          0B       974Mi
$ sudo swapoff -a
$ free -mh
              total        used        free      shared  buff/cache   available
Mem:          987Mi        64Mi       831Mi       2.0Mi        92Mi       805Mi
Swap:            0B          0B          0B
$ sudo swapon -a
$ free -mh
              total        used        free      shared  buff/cache   available
Mem:          987Mi        64Mi       831Mi       2.0Mi        92Mi       806Mi
Swap:         974Mi          0B       974Mi
```
