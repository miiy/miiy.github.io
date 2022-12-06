---
layout: post
title: "linux 挂载 qcow2 磁盘镜像"
date: 2022-12-06
tags: linux
---

加载nbd（网络块设备: Network Block Device）模块

```bash
sudo lsmod |grep nbd
sudo modprobe nbd max_part=8
```


连接 qemu-nbd

```bash
sudo qemu-nbd --connect=/dev/nbd0 /path/to/image.qcow2
```

查看分区信息

```bash
$ sudo fdisk -l /dev/nbd0
Disk /dev/nbd0: 200 GiB, 214748364800 bytes, 419430400 sectors
Units: sectors of 1 * 512 = 512 bytes
Sector size (logical/physical): 512 bytes / 512 bytes
I/O size (minimum/optimal): 512 bytes / 512 bytes
Disklabel type: gpt
Disk identifier: 21A2186E

Device      Start       End   Sectors  Size Type
/dev/nbd0p1    34     32767     32734   16M Microsoft reserved
/dev/nbd0p2 32768 419426303 419393536  200G Microsoft basic data
```

挂载分区

```bash
sudo mount /dev/nbd0p2 qcow2_mount_point
```

卸载分区和断开连接

```bash
sudo umount qcow2_mount_point
sudo qemu-nbd --disconnect /dev/nbd0
sudo rmmod nbd
```
