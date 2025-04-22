---
layout: post
title: "fdisk"
---

fdisk - manipulate disk partition table

```bash
$ sudo fdisk -l
Disk /dev/loop0: 4 KiB, 4096 bytes, 8 sectors
Units: sectors of 1 * 512 = 512 bytes
Sector size (logical/physical): 512 bytes / 512 bytes
I/O size (minimum/optimal): 512 bytes / 512 bytes

Disk /dev/nvme1n1: 931.51 GiB, 1000204886016 bytes, 1953525168 sectors
Disk model: Samsung SSD 980 PRO 1TB
Units: sectors of 1 * 512 = 512 bytes
Sector size (logical/physical): 512 bytes / 512 bytes
I/O size (minimum/optimal): 512 bytes / 512 bytes
Disklabel type: gpt
Disk identifier: 1E00E000-123D-1CE2-12AD-10172101C21A

Device         Start        End    Sectors   Size Type
/dev/nvme1n1p1 65535 1953467279 1953401745 931.5G Microsoft basic data
```

```bash
sudo mount /dev/nvme1n1 /mnt/mydisk
```