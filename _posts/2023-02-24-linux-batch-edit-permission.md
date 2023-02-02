---
layout: post
title: "批量修改权限"
date: 2023-02-24
tags: linux
---

```bash
find -type d|xargs chmod 755
find -type f|xargs chmod 644
```