---
layout: post
title: "journalctl"
---

journalctl：检索 systemd 日志

常用参数：

- -b --boot[=ID]             Show current boot or the specified boot
- -e --pager-end             Immediately jump to the end in the pager
- -f --follow                Follow the journal
- -u --unit=UNIT             Show logs from the specified unit

```bash
# 显示所有的消息
$ journalctl
# 显示本次启动后的所有消息
$ journalctl -b
# 显示制定单元的所有消息
$ journalctl -u netcfg
```