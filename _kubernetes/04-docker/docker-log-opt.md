---
layout: post
title: "设置 docker 日志文件大小"
---

通过 --log-opt 设置日志文件大小

<https://docs.docker.com/config/containers/logging/local/>

```bash
--log-opt max-size=10m
```

查看日志文件大小

```bash
#!/bin/sh

echo "======== docker containers logs file size ========"

logs=$(find /var/lib/docker/containers/ -name *-json.log)

for log in $logs
    do
        ls -lh $log
    done
```

清空日志

```bash
cat /dev/null > containerid-json.log
```
