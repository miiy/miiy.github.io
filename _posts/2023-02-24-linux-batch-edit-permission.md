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

上面这个命令如果碰到带有空格的文件名会出错


```bash
find . -type d -exec chmod -fv 755 {} \;
find . -type f -exec chmod -fv 644 {} \;
```

<https://askubuntu.com/questions/604489/what-are-the-default-permissions-of-directories-in-home-desktop-downloads-etc>