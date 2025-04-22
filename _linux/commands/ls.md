---
layout: post
title: "ls"
---

统计文件和文件夹数

```bash
ls dir/ | wc -l
```

按文件大小排序

```bash
ls -lSh dir/ | head -n 10
ls -lShr dir/ | head -n 10 | awk '{print $9}'
```

按时间排序

```bash
ls -lth dir/ | head -n 10
```

