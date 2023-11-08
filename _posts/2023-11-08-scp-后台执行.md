---
layout: post
title: "scp 命令后台执行"
date: 2023-11-08
tags: linux
---

当拷贝一个很大的文件的时候可以使 scp 命令在后台执行来拷贝文件。

运行 scp 命令拷贝文件

```bash
scp root@aliyun:/path/to/file .
```

按 Ctrl + z 暂停任务

```bash
file     0%  576KB  63.3KB/s 13:51:38 ET^Z[1]  + 32226 suspended (signal)  scp root@aliyun:/path/to/file .
[1]+  Stopped                 scp root@aliyun:/path/to/file .
```

查看任务信息

``bash
jobs
[1]  + suspended (signal)  scp root@ali:/path/to/file .
``

将任务放到后台运行

```bash
$ bg %1
[1]  + 32226 continued  scp root@aliyun:/path/to/file .
$ jobs
[1]  + running    scp root@aliyun:/path/to/file .
```

忽略此作业的 HUP 信号

```bash
$ disown -h %1
$ jobs
[1]+  Running                 scp root@aliyun:/path/to/file . &
```

退出重新登录查看任务是否在后台运行

```bash
$ ps -ef|grep scp
user      2821       1  0 22:05 ?        00:00:01 scp root@aliyun:/path/to/file .
```

参考：

https://blog.csdn.net/wjn2000414/article/details/80797632
