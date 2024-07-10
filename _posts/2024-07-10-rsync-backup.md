---
layout: post
title: "使用 crontab + rsync 自动备份到备用硬盘"
date: 2024-06-28
tags: linux
---

## script

```bash
#!/bin/bash

# https://rsync.samba.org/
# https://rsync.samba.org/examples.html

# cron
# sudo crontab -e
# 0 0 * * * sudo /path/to/backup.sh >> /var/log/cron.log 2>&1

BACKUP_DIR=/home/
DST_DIR=/media/debian/device/home/

if [ -d "$DST_DIR" ]; then
    rsync -av --delete --ignore-errors $BACKUP_DIR $DST_DIR
    echo -e "\n"
    echo "backup time: $(date '+%Y-%m-%d %H:%M:%S')"
else
    echo "$DST_DIR not exisit."
fi
```
