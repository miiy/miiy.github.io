---
layout: post
title: "install php"
date: 2016-06-13
tags: php
---

# Install

```bash
groupadd www
useradd -s /sbin/nologin -g www -M www
yum install -y gcc gcc-c++ autoconf \
zlib zlib-devel openssl openssl-devel libxml2 libxml2-devel bzip2-devel libcurl-devel openldap-devel readline-devel epel-release libmcrypt-devel \
libjpeg-devel libpng-devel freetype-devel gd-devel
wget http://am1.php.net/distributions/php-7.2.0.tar.gz\
&& tar -zxvf php-7.2.0.tar.gz \
&& cd php-7.2.0 \
&& ./configure \
--prefix=/usr/local/php \
--with-config-file-path=/usr/local/php/etc \
--with-config-file-scan-dir=/usr/local/php/etc/conf.d \
--enable-fpm \
--with-fpm-user=www \
--with-fpm-group=www \
--enable-mbstring \
--enable-mysqlnd \
--with-mysqli=mysqlnd \
--with-pdo-mysql=mysqlnd \
--with-mysql-sock=/var/lib/mysql/mysql.sock \
--enable-pcntl \
--enable-shmop \
--enable-opcache \
--enable-soap \
--enable-bcmath \
--enable-sockets \
--enable-ftp \
--with-curl \
--with-zip \
--with-iconv \
--with-openssl \
--with-readline \
--with-zlib \
--with-bz2 \
--enable-exif \
--with-gd \
--with-jpeg-dir \
--with-png-dir \
--with-freetype-dir \
--enable-gd-native-ttf \
 
&& make -j2 \
&& make install \
&& cp php.ini-development /usr/local/php/etc/php.ini \
&& cp /usr/local/php/etc/php-fpm.conf.default /usr/local/php/etc/php-fpm.conf \
&& cp /usr/local/php/etc/php-fpm.d/www.conf.default /usr/local/php/etc/php-fpm.d/www.conf \
&& cp sapi/fpm/init.d.php-fpm /etc/init.d/php-fpm \
&& chmod +x /etc/init.d/php-fpm \
&& chkconfig php-fpm on \
&& echo "export PATH=/usr/local/php/bin:/usr/local/php/sbin:\$PATH" >> /etc/profile \
&& source /etc/profile \
&& pecl install xdebug redis memcached mongodb swoole mcrypt-1.0.1 \
&& mkdir -p /usr/local/php/etc/conf.d \
&& touch /usr/local/php/etc/conf.d/extension.ini \
&& echo "zend_extension=xdebug" >> /usr/local/php/etc/conf.d/extension.ini \
&& echo "xdebug.remote_enable=on" >> /usr/local/php/etc/conf.d/extension.ini \
&& echo "xdebug.remote_port=9000" >> /usr/local/php/etc/conf.d/extension.ini \
&& echo "xdebug.remote_connect_back=on" >> /usr/local/php/etc/conf.d/extension.ini \
&& echo "xdebug.remote_handler=dbgp" >> /usr/local/php/etc/conf.d/extension.ini \
&& echo "extension=redis" >> /usr/local/php/etc/conf.d/extension.ini \
&& echo "extension=memcached" >> /usr/local/php/etc/conf.d/extension.ini \
&& echo "extension=mongodb" >> /usr/local/php/etc/conf.d/extension.ini \
&& echo "extension=swoole" >> /usr/local/php/etc/conf.d/extension.ini \
&& echo "extension=mcrypt" >> /usr/local/php/etc/conf.d/extension.ini
```

## CentOS7下编译安装PHP

```bash
yum install -y gcc gcc-c++ autoconf automake
# PHP扩展依赖
yum install -y libxml2-devel openssl-devel libcurl-devel  libicu-devel gd-devel libjpeg-devel libpng-devel  freetype-devel
# PHP
cd /usr/local/src
wget http://cn2.php.net/distributions/php-5.6.16.tar.gz
tar -zxvf php-5.6.16.tar.gz
cd php-5.6.16
./configure \
--prefix=/usr/local/php \
--with-config-file-path=/usr/local/php/etc \
--enable-fpm \
--with-mysql=mysqlnd \
--with-mysqli=mysqlnd \
--with-pdo-mysql=mysqlnd \
--enable-mbstring \
--enable-bcmath \
--enable-opcache \
--enable-calendar \
--enable-soap \
--enable-sockets \
--enable-zip \
--enable-ftp \
--enable-intl \
--with-icu-dir=/usr \
--with-zlib \
--with-curl \
--with-openssl \
--with-mhash \
--with-gettext \
--with-gd \
--with-png-dir \
--with-jpeg-dir \
--with-freetype-dir \
--enable-gd-native-ttf
make
make install
cp php.ini-production /usr/local/php/etc/php.ini
cp /usr/local/php/etc/php-fpm.conf.default /usr/local/php/etc/php-fpm.conf
```

### 设置php-fpm开机启动

```
cp /usr/local/src/php-5.6.16/sapi/fpm/init.d.php-fpm /etc/rc.d/init.d/php-fpm
chmod +x /etc/rc.d/init.d/php-fpm
chkconfig php-fpm on
```

