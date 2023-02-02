---
layout: post
title: "MacOS编译安装PHP"
date: 2023-02-02
tags: macos
---

## INSTALL PHP7.4.33

```bash
brew install pkg-config autoconf libiconv oniguruma openssl libzip bzip2 gd libpng libjpeg webp

export PKG_CONFIG_PATH="/usr/local/opt/openssl/lib/pkgconfig:/usr/local/opt/libxml2/lib/pkgconfig"

./configure \
--prefix=/usr/local/php-7.4.33 \
--with-config-file-path=/usr/local/php-7.4.33/etc \
--with-config-file-scan-dir=/usr/local/php-7.4.33/etc/conf.d \
--with-pear \
--enable-fpm \
--enable-mbstring \
--enable-mysqlnd \
--with-mysqli=mysqlnd \
--with-pdo-mysql=mysqlnd \
--with-sqlite3 \
--with-pdo-sqlite=/usr/local/opt/sqlite3 \
--enable-pcntl \
--enable-shmop \
--enable-soap \
--with-xsl \
--enable-bcmath \
--with-openssl=/usr/local/opt/openssl \
--enable-sockets \
--enable-ftp \
--with-openssl-dir=/usr/local/opt/openssl \
--with-curl \
--with-zip \
--with-iconv=/usr/local/opt/libiconv \
--with-readline=/usr/local/opt/readline \
--with-zlib \
--with-bz2=/usr/local/opt/bzip2 \
--enable-exif \
--enable-gd \
--with-webp \
--with-jpeg \
--with-freetype

make
sudo make install
cp php.ini-production /usr/local/php/etc/php.ini
cd /usr/local
ln -s php-7.4.33 php
```

添加PATH环境变量

```bash
vi ~/.zshrc

export PATH="/usr/local/php/bin:/usr/local/php/sbin:$PATH"
```

Extension

```bash
pear config-set http_proxy http://127.0.0.1:1087
pecl install xdebug-3.1.6 redis mongodb swoole-4.8.12 xlswriter
vi /usr/local/php/etc/conf.d/custom.ini
```

```text
expose_php = Off
# xdebug
xdebug.start_with_request=yes
xdebug.mode=develop,debug,profile,trace
# xdebug.client_host=host.docker.internal
xdebug.client_host=127.0.0.1
xdebug.client_port=9003
xdebug.idekey=PHPSTORM
#xdebug.log="/tmp/xdebug.log"
```

```bash
vi /usr/local/php-7.4.33/etc/conf.d/extensions.ini
```

```text
zend_extension=opcache
zend_extension=xdebug
extension=redis
extension=mongodb
extension=swoole
extension=xlswriter
```

```bash
cd /usr/local/php-7.4.33/etc
cp php-fpm.conf.default php-fpm.conf
cp php-fpm.d/www.conf.default php-fpm.d/www.conf
```

当前用户启动php-fpm，慢日志www.log.slow无法写问题：

```log
[18-Jul-2020 23:27:23] WARNING: [pool www] child 82353, script '/path/to/public/index.php' (request: "GET /index.php") executing too slow (2.653167 sec), logging
[18-Jul-2020 23:27:23] NOTICE: child 82353 stopped for tracing
[18-Jul-2020 23:27:23] NOTICE: about to trace 82353
[18-Jul-2020 23:27:23] ERROR: task_for_pid() failed: (os/kern) failure (5) It seems that master process does not have enough privileges to trace processes.
[18-Jul-2020 23:27:23] NOTICE: finished trace of 82353
```

www.conf

```conf
user = 当前用户
group = staff
```

```bash
sudo php-fpm
```
