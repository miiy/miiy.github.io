---
layout: post
title: "MacOS编译安装Nginx, Redis, PHP"
date: 2020-07-09
category: other
---

## MySQL

https://dev.mysql.com/downloads/mysql/

https://cdn.mysql.com//Downloads/MySQL-8.0/mysql-8.0.21-macos10.15-x86_64.dmg

## Nginx

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

## Redis

```bash
make PREFIX=/usr/local/redis install
```

## PHP

```bash
brew install pkg-config autoconf libiconv oniguruma openssl libzip bzip2 gd libpng libjpeg webp

export PKG_CONFIG_PATH="/usr/local/opt/openssl/lib/pkgconfig:/usr/local/opt/libxml2/lib/pkgconfig"

./configure \
--prefix=/usr/local/php \
--with-config-file-path=/usr/local/php/etc \
--with-config-file-scan-dir=/usr/local/php/etc/conf.d \
--with-pear \
--enable-fpm \
--enable-mbstring \
--enable-mysqlnd \
--with-mysqli=mysqlnd \
--with-pdo-mysql=mysqlnd \
# --with-mysql-sock=/var/lib/mysql/mysql.sock \
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
sudo maks install
```

PATH

~/.zshrc

```bash
export PATH="/usr/local/php/bin:/usr/local/php/sbin:$PATH"
```

Extension

```bash
pear config-set http_proxy http://127.0.0.1:1087
pecl install xdebug redis mongodb swoole
```

/usr/local/php/etc/conf.d/custom.ini 

```ini
expose_php = Off
# xdebug
xdebug.remote_enable=0
xdebug.remote_host=host.docker.internal
xdebug.remote_port=9000
xdebug.idekey=PHPSTORM
```

/usr/local/php/etc/conf.d/extensions.ini

```ini
zend_extension=opcache
zend_extension=xdebug
extension=redis
extension=mongodb
extension=swoole
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

sudo php-fpm

## alias

```bash
alias srvs='ps -ef|grep nginx;ps -ef|grep php;ps -ef|grep mysql;ps -ef|grep redis'

alias msqstart='sudo /usr/local/mysql/support-files/mysql.server start'
alias msqstop='sudo /usr/local/mysql/support-files/mysql.server stop'

alias rds-srv='/usr/local/redis/bin/redis-server /usr/local/redis/redis.conf'
alias rdsquit='kill -QUIT `cat /usr/local/redis/redis_6379.pid`'

alias ngx='/usr/local/nginx/sbin/nginx'
alias ngx-t='/usr/local/nginx/sbin/nginx -t'
alias ngxreload='/usr/local/nginx/sbin/nginx -s reload'
alias ngxquit='/usr/local/nginx/sbin/nginx -s quit'

alias fpmstart='sudo php-fpm -g /usr/local/php/var/run/php-fpm.pid'
alias fpmreload='sudo kill -USR2 `cat /usr/local/php/var/run/php-fpm.pid`'
alias fpmquit='sudo kill -QUIT `cat /usr/local/php/var/run/php-fpm.pid`'
```

## MacOS 自带的 php 编译参数

```text
'/Library/Caches/com.apple.xbs/Binaries/apache_mod_php/install/TempContent/Objects/php/configure'
'--prefix=/usr'
'--mandir=/usr/share/man'
'--infodir=/usr/share/info'
'--disable-dependency-tracking'
'--sysconfdir=/private/etc'
'--with-libdir=lib'
'--enable-cli'
'--with-iconv=/Applications/Xcode.app/Contents/Developer/Platforms/MacOSX.platform/Developer/SDKs/MacOSX10.15.Internal.sdk/usr'
'--with-config-file-path=/etc'
'--with-libxml-dir=/Applications/Xcode.app/Contents/Developer/Platforms/MacOSX.platform/Developer/SDKs/MacOSX10.15.Internal.sdk/usr'
'--with-openssl=/Applications/Xcode.app/Contents/Developer/Platforms/MacOSX.platform/Developer/SDKs/MacOSX10.15.Internal.sdk/usr/local/libressl'
'--with-kerberos=/Applications/Xcode.app/Contents/Developer/Platforms/MacOSX.platform/Developer/SDKs/MacOSX10.15.Internal.sdk/usr'
'--with-zlib=/Applications/Xcode.app/Contents/Developer/Platforms/MacOSX.platform/Developer/SDKs/MacOSX10.15.Internal.sdk/usr'
'--enable-bcmath'
'--with-bz2=/Applications/Xcode.app/Contents/Developer/Platforms/MacOSX.platform/Developer/SDKs/MacOSX10.15.Internal.sdk/usr'
'--enable-calendar'
'--disable-cgi'
'--with-curl=/Applications/Xcode.app/Contents/Developer/Platforms/MacOSX.platform/Developer/SDKs/MacOSX10.15.Internal.sdk/usr'
'--enable-dba'
'--with-ndbm=/Applications/Xcode.app/Contents/Developer/Platforms/MacOSX.platform/Developer/SDKs/MacOSX10.15.Internal.sdk/usr' 
'--enable-exif'
'--enable-fpm'
'--enable-ftp'
'--with-gd'
'--with-png-dir=/Library/Caches/com.apple.xbs/Binaries/apache_mod_php/install/TempContent/Root/usr/local'
'--with-jpeg-dir=/Library/Caches/com.apple.xbs/Binaries/apache_mod_php/install/TempContent/Root/usr/local'
'--enable-gd-native-ttf'
'--with-icu-dir=/Applications/Xcode.app/Contents/Developer/Platforms/MacOSX.platform/Developer/SDKs/MacOSX10.15.Internal.sdk/usr'
'--with-ldap=/Applications/Xcode.app/Contents/Developer/Platforms/MacOSX.platform/Developer/SDKs/MacOSX10.15.Internal.sdk/usr'
'--with-ldap-sasl=/Applications/Xcode.app/Contents/Developer/Platforms/MacOSX.platform/Developer/SDKs/MacOSX10.15.Internal.sdk/usr'
'--with-libedit=/Applications/Xcode.app/Contents/Developer/Platforms/MacOSX.platform/Developer/SDKs/MacOSX10.15.Internal.sdk/usr'
'--enable-mbstring'
'--enable-mbregex'
'--with-mysqli=mysqlnd'
'--without-pcre-jit'
'--with-pdo-pgsql=/Applications/Xcode.app/Contents/Developer/Toolchains/OSX10.15.xctoolchain/usr/local/bin'
'--with-pgsql=/Applications/Xcode.app/Contents/Developer/Toolchains/OSX10.15.xctoolchain/usr/local/bin'
'--without-pear'
'--with-pear=no'
'--with-pdo-mysql=mysqlnd'
'--with-mysql-sock=/var/mysql/mysql.sock'
'--disable-phpdbg'
'--with-readline=/Applications/Xcode.app/Contents/Developer/Platforms/MacOSX.platform/Developer/SDKs/MacOSX10.15.Internal.sdk/usr'
'--enable-shmop'
'--with-snmp=/Applications/Xcode.app/Contents/Developer/Platforms/MacOSX.platform/Developer/SDKs/MacOSX10.15.Internal.sdk/usr'
'--enable-soap'
'--enable-sockets'
'--enable-sysvmsg'
'--enable-sysvsem'
'--enable-sysvshm'
'--with-tidy=/Applications/Xcode.app/Contents/Developer/Platforms/MacOSX.platform/Developer/SDKs/MacOSX10.15.Internal.sdk/usr'
'--enable-wddx'
'--with-xmlrpc'
'--with-iconv-dir=/Applications/Xcode.app/Contents/Developer/Platforms/MacOSX.platform/Developer/SDKs/MacOSX10.15.Internal.sdk/usr'
'--with-xsl=/Applications/Xcode.app/Contents/Developer/Platforms/MacOSX.platform/Developer/SDKs/MacOSX10.15.Internal.sdk/usr'
'--with-apxs2=/Applications/Xcode.app/Contents/Developer/Toolchains/OSX10.15.xctoolchain/usr/local/bin/apxs'
'YACC=/Applications/Xcode.app/Contents/Developer/Toolchains/XcodeDefault.xctoolchain/usr/bin/bison'
```

/Applications/Xcode.app/Contents/Developer/Platforms/MacOSX.platform/Developer/SDKs/MacOSX.sdk/usr/