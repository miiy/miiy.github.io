---
layout: post
title: "MacOS 系统下制作U盘镜像"
date: 2025-03-12
---

```bash
hdiutil convert -format UDRW -o ubuntu-22.04 ubuntu-22.04.4-desktop-amd64.iso

diskutil list

diskutil unmountDisk /dev/disk2

sudo dd if=ubuntu-22.04.dmg  of=/dev/rdisk2 bs=1m

diskutil eject /dev/disk2
```
