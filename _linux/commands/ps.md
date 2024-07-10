---
layout: post
title: "ps"
---

ps - report a snapshot of the current processes.

ps - 报告当前进程的快照。

查看每个进程

```bash
# standard syntax
ps -ef

# BSD syntax
ps aux
```

查看内存占用情况

```bash
ps aux --sort=-%mem | head -10
```
