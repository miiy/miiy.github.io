---
layout: post
title: "mount"
---

指定挂载后文件所有者的用户ID和组ID

```bash
sudo mount -o uid=1000,gid=1000 /dev/sda1 ~/path
```

卸载

```bash
sudo umount /dev/sda1
```
