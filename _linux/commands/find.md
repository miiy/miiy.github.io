---
layout: post
title: "find"
---

```bash
find . -name "*.txt" -type f

find / -name "*.txt" 2>/dev/null
```

清理 .DS_Store

```bash
find /path -type f -name ".DS_Store" -exec rm {} +
```