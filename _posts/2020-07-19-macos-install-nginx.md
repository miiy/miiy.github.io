---
layout: post
title: "MacOS编译安装Nginx"
date: 2020-07-09
tags: macos
---

## Install

Building nginx from Sources: https://nginx.org/en/docs/configure.html

https://nginx.org/en/download.html

http://www.pcre.org/

https://www.openssl.org/source/

https://www.zlib.net/

```bash
mkdir -p /usr/local/src/serv
chown -R {user} /usr/local/src/serv
chgrp -R staff /usr/local/src/serv

curl -o nginx-1.18.0.tar.gz https://nginx.org/download/nginx-1.18.0.tar.gz
curl -o pcre2-10.35.tar.gz ftp://ftp.pcre.org/pub/pcre/pcre2-10.35.tar.gz
curl -o openssl-1.1.1g.tar.gz https://www.openssl.org/source/openssl-1.1.1g.tar.gz
curl -o zlib-1.2.11.tar.gz https://www.zlib.net/zlib-1.2.11.tar.gz

tar -zxvf nginx-1.18.0.tar.gz
tar -zxvf pcre2-10.35.tar.gz
tar -zxvf openssl-1.1.1g.tar.gz
tar -zxvf zlib-1.2.11.tar.gz

cd nginx-1.18.0

./configure \
--prefix=/usr/local/nginx \
--with-pcre=../pcre-8.44 \
--with-zlib=../zlib-1.2.11 \
--with-openssl=../openssl-1.1.1g \
--http-client-body-temp-path=/usr/local/nginx/client_body_temp \
--http-proxy-temp-path=/usr/local/nginx/proxy_temp \
--http-fastcgi-temp-path=/usr/local/nginx/fastcgi_temp \
--http-uwsgi-temp-path=/usr/local/nginx/uwsgi_temp \
--http-scgi-temp-path=/usr/local/nginx/scgi_temp \
--with-http_ssl_module

make
sudo make install

choown -R [user] /usr/local/nginx
```
