---
layout: post
title: "debian 禁用挂起"
date: 2023-02-24
tags: linux
---

<https://wiki.debian.org/Suspend>


/etc/systemd/sleep.conf.d/nosuspend.conf


```conf
[Sleep]
AllowSuspend=no
AllowHibernation=no
AllowSuspendThenHibernate=no
AllowHybridSleep=no
```

```bash
sudo systemctl status sleep.target suspend.target hibernate.target hybrid-sleep.target 
```
