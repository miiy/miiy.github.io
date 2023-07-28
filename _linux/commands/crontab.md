---
layout: post
title: "crontab"
---

Linux crontab 是用来定期执行命令的程序

分 时 日 月 周

格式如下：

```text
*    *    *    *    *    command
-    -    -    -    -
|    |    |    |    |
|    |    |    |    +----- 星期中星期几 (0 - 6) (星期天 为0)
|    |    |    +---------- 月份 (1 - 12) 
|    |    +--------------- 一个月中的第几天 (1 - 31)
|    +-------------------- 小时 (0 - 23)
+------------------------- 分钟 (0 - 59)
```

https://www.runoob.com/linux/linux-comm-crontab.html

https://www.cnblogs.com/intval/p/5763929.html
