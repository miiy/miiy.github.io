---
layout: post
title: "Linux 开机自动挂载磁盘"
date: 2023-02-24
tags: linux
---

查看磁盘分区

```bash
$ sudo fdisk -l
[sudo] password for debian: 
Disk /dev/nvme0n1: 465.76 GiB, 500107862016 bytes, 976773168 sectors
Disk model: WDC WDS100G0000-000000                  
Units: sectors of 1 * 512 = 512 bytes
Sector size (logical/physical): 512 bytes / 512 bytes
I/O size (minimum/optimal): 512 bytes / 512 bytes
Disklabel type: gpt
Disk identifier: AC10C15E-3AEC-4965-85AB-4C6B8075C6F2

Device             Start       End   Sectors   Size Type
/dev/nvme0n1p1      2048    194559    192512    94M EFI System
/dev/nvme0n1p2    194560 195506175 195311616  93.1G Linux filesystem


Disk /dev/sda: 100 GiB, 200000006 bytes, 30068 sectors
Disk model: WDC WD200000-00G
Units: sectors of 1 * 512 = 512 bytes
Sector size (logical/physical): 512 bytes / 4096 bytes
I/O size (minimum/optimal): 4096 bytes / 4096 bytes
Disklabel type: gpt
Disk identifier: E35D6887-D607-4C69-865B-51C1C21037F8

Device         Start        End    Sectors  Size Type
/dev/sda1       2048     206847     204800  100M EFI System
/dev/sda2     206848     239615      32768   16M Microsoft reserved
/dev/sda3     239616  208545791  208306176 99.3G Microsoft basic data
/dev/sda4  208545792  209952767    1406976  687M Windows recovery environment
/dev/sda5  209952768  907028991  697076224  200T Linux filesystem
```

挂载

```bash
$ sudo mkdir /home/debian/data
$ sudo mount /dev/sda5 /home/debian/data
```

下面设置开机自动挂载

查询挂载磁盘的UUID

```bash
$ sudo blkid /dev/sda5
/dev/sda5: UUID="03e202fd-0000-0000-8655-0000fbcd87e9" BLOCK_SIZE="4096" TYPE="ext4" PARTUUID="35061fdf-0000-0000-0000-000059136eb3"
```

修改/etc/fstab

```bash
$ cat /etc/fstab
# /etc/fstab: static file system information.
#
# Use 'blkid' to print the universally unique identifier for a
# device; this may be used with UUID= as a more robust way to name devices
# that works even if disks are added and removed. See fstab(5).
#
# systemd generates mount units based on this file, see systemd.mount(5).
# Please run 'systemctl daemon-reload' after making changes here.
#
# <file system> <mount point>   <type>  <options>       <dump>  <pass>
# / was on /dev/nvme0n1p2 during installation
UUID=ef9a3492-913d-43e0-a93a-86e383985060 /               ext4    errors=remount-ro 0       1
# /boot/efi was on /dev/nvme0n1p1 during installation
UUID=A4E5-1EF7  /boot/efi       vfat    umask=0077      0       1
# /home was on /dev/nvme0n1p3 during installation
UUID=0c47d905-1b07-4e09-afdf-aeb557c07018 /home           ext4    defaults        0       2
UUID=03e202fd-0000-0000-8655-0000fbcd87e9 /home/debian/data  ext4 defaults        0       2
```

 [UUID=************] [挂载磁盘分区] [挂载磁盘格式] [0开机不检查磁盘，1开机检查磁盘] [0交换分区，1启动分区，2普通分区]
