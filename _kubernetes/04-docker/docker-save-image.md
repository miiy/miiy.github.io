
---
layout: post
title: "docker 保存镜像"
---

使用 docker save image.tar 导出的镜像，有时候在导入时会报错：invalid diffID for layer

使用 gzip 保存镜像文件

```bash
$ docker save image:v1 | gzip > imagev1.tar.gz
$ docker load < imagev1.tar.gz
```

<https://docs.docker.com/reference/cli/docker/image/save/>
