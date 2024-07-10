---
layout: post
title: "docker 登录私库时提示 x509: certificate signed by unknown authority"
---

```bash
docker login 192.168.0.1:7888
Username: admin
Password: 
Error response from daemon: Get "https://192.168.0.1:7888/v2/": tls: failed to verify certificate: x509: certificate signed by unknown authority
```

修改 ~/.docker/daemon.json 添加

```text
"insecure-registries": ["192.168.0.1:7888"]
```

重启 docker 然后重新登录

```bash
systemctl daemon-reload
systemctl reload docker
```

<https://docs.docker.com/reference/cli/docker/login/>
