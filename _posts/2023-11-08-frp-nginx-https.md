---
layout: post
title: "使用 frp 将请求转发到内网的 https"
date: 2023-11-08
tags: linux
---

## 问题

使用 frps 转发 https，在 nginx 中这样配置是不起作用的。

```
proxy_pass https://127.0.0.1:7043
```

因为对于 SSL， frps 无法获取 hostname，即使设置了 proxy_set_header Host $host;

可以设置 proxy_ssl_name，使 frps 可以识别正确的 host

## 解决方案

frps.toml

```toml
bindPort = 7000
auth.token = "123456"
vhostHTTPPort = 7080
vhostHTTPSPort = 7043
log.level = "debug
```

nginx 配置

如果 nginx 是运行在 docker 中的，在容器启动的时候添加 hosts 以便可以访问到宿主机。

```conf
upstream test_server {
    server host.docker.internal:7643;
}

location / {
    proxy_pass       https://test_server/;
    proxy_ssl_server_name on;
    proxy_ssl_name                    $host;
}
```

内网 frpc 配置

```toml
serverAddr = "xx.xx.xx.xx"
serverPort = 7000
auth.token = "123456"

[[proxies]]
name = "testhttp"
type = "http"
localIP = "127.0.0.1"
localPort = 80
customDomains = ["a.test.com"]
requestHeaders.set.x-from-where = "frp"

[[proxies]]
name = "testhttps"
type = "https"
localIP = "127.0.0.1"
localPort = 443
customDomains = ["a.test.com"]
```

## frp windows 设置开机自启

<https://blog.itpub.net/69955379/viewspace-2731754/>

## frp Linux 设置开机自启

/etc/systemd/system/frpc.service

```txt
[Unit]
Description = frp client
After = network.target syslog.target
Wants = network.target

[Service]
Type = simple
ExecStart = /usr/local/frp/frpc -c /usr/local/frp/frpc-ali.toml
Restart=always
RestartSec=15s

[Install]
WantedBy = multi-user.target
```


## Refrences

https://github.com/fatedier/frp/issues/610

https://github.com/fatedier/frp/issues/888

https://blog.csdn.net/liuxiao723846/article/details/127749786