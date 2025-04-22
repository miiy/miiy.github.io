---
layout: post
title: "使用 logrotate 管理日志"
date: 2025-04-09
tags: linux
---

添加配置文件

```bash
vi /etc/logrotate.d/cron
```

```txt
/var/log/cron.log {
  rotate 4
  monthly
  compress
  missingok
  notifempty
}
```

验证配置文件是否生效

```bash
sudo logrotate /etc/logrotate.d/cronlog --debug
```
