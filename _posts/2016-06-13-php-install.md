---
layout: post
title: "php install"
date: 2016-06-13
tags: php
---

# Install

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

### 安装PHP扩展

```bash
vi /usr/local/php/etc/php.ini #添加extension_dir与extension格式如下：

extension_dir = "/usr/local/php/lib/php/extensions/no-debug-non-zts-20131226"
extension=redis.so
extension=memcache.so
extension=memcached.so
```

### Redis

```
wget https://github.com/phpredis/phpredis/archive/2.2.7.tar.gz
tar -zxvf phpredis-2.2.7.tar.gz
cd phpredis-2.2.7
/usr/local/php/bin/phpize
./configure --with-php-config=/usr/local/php/bin/php-config
make && make install
```

示例代码：

```php
<?php
   $redis = new Redis();
   $redis->connect('127.0.0.1', 6379);
   $redis->set("mykey", "redis");
   echo $redis->get("mykey");
?>

```

### Memcache

```
cd /usr/local/src
wget http://pecl.php.net/get/memcache-2.2.7.tgz
tar -zxvf memcache-2.2.7.tgz
cd memcache-2.2.7
/usr/local/php/bin/phpize
./configure --with-php-config=/usr/local/php/bin/php-config
make && make install
```

示例代码

```php
<?php
$memcache = new Memcache;
$memcache->connect('localhost', 11211);
$memcache->set('mykey', 'memcache');
echo $memcache->get('mykey');
?>
```

Memcached

```bash
cd /usr/local/src
wget http://pecl.php.net/get/memcached-2.2.0.tgz
tar -zxvf memcached-2.2.0.tgz
cd memcached-2.2.0
/usr/local/php/bin/phpize
./configure --with-php-config=/usr/local/php/bin/php-config --with-libmemcached-dir=/usr/local/libmemcached --disable-memcached-sasl
make && make install
```

测试代码

```php
<?php
$memcached = new Memcached;
$memcached->addServer('localhost', 11211);
$memcached->set('mykey', 'memcached');
echo $memcached->get('mykey');
?>
```

## CentOS7下编译安装Redis

```bash
cd /usr/local/src
wget http://download.redis.io/releases/redis-3.0.5.tar.gz
tar -zxzf redis-3.0.5.tar.gz
cd redis-3.0.5
make #如果出错make MALLOC=libc
make install
```
make install执行完成后，会在/usr/local/bin目录下生成一下个可执行文件：

* redis-server：Redis服务器的daemon启动程序
* redis-cli：Redis命令行操作工具。也可以用telnet根据其纯文本协议
来操作
* redis-benchmark：Redis性能测试工具，测试Redis在当前系统下的读写性能
* redis-check-aof：数据修复
* redis-check-dump：检查导出工具

自启动

```bash
mkdir /etc/redis
cp redis.conf /etc/redis/6379.conf
vi /etc/redis/6379.conf #修改daemonize yes使进程在后台运行
cp /usr/local/src/redis-3.0.5/utils/redis_init_script /etc/rc.d/init.d/redisd
chmod +x /etc/rc.d/init.d/redisd
vi /etc/rc.d/init.d/redisd #在头部加入#chkconfig: 2345 90 10不然会报错service redis does not support chkconfig
chkconfig redisd on
```

参考：

http://redis.io/download

http://futeng.iteye.com/blog/2071867

http://www.cnblogs.com/zhuhongbao/archive/2013/06/04/3117997.html

http://www.runoob.com/redis/redis-php.html

##CentOS7下编译安装Memcached

```bash
yum -y install libevent-deve #安装Memcached依赖
cd /usr/local/src
wget http://www.memcached.org/files/memcached-1.4.25.tar.gz
tar -zxvf memcached-1.4.25.tar.gz
cd memcached-1.4.25
./configure --prefix=/usr/local/memcached
make && make install
```

启动选项：
* d是启动一个守护进程；
* m是分配给Memcache使用的内存数量，单位是MB；
* u是运行Memcache的用户；
* l是监听的服务器IP地址，可以有多个地址；
* p是设置Memcache监听的端口，，最好是1024以上的端口；
* c是最大运行的并发连接数，默认是1024；
* P是设置保存Memcache的pid文件。

```bash
/usr/local/memcached/bin/memcached -d -m 64m -p 11211 -u root
```

设置启动脚本

```bash
cp /usr/local/src/memcached-1.4.25/scripts/memcached.sysv /etc/rc.d/init.d/memcached
chmod +x /etc/rc.d/init.d/memcached
vi /etc/rc.d/init.d/memcached #修改start函数中daemon memcached位置
chkconfig memcached on
```

参考：

http://www.linuxidc.com/Linux/2015-04/116240.htm

http://www.cnblogs.com/technet/archive/2011/09/11/2173485.html

http://www.runoob.com/Memcached/php-connect-memcached.html

##CentOS7下编译安装libmemcached

```bash
cd /usr/local/src
wget https://launchpad.net/libmemcached/1.0/1.0.18/+download/libmemcached-1.0.18.tar.gz
tar -zxvf libmemcached-1.0.18.tar.gz
cd libmemcached-1.0.18
./configure --prefix=/usr/local/libmemcached --with-memcached
make && make install
```

## CentOS7下编译安装Nodejs

```bash
cd /usr/local/src
wget https://nodejs.org/dist/v4.2.1/node-v4.2.1.tar.gz
tar -zxvf node-v4.2.1.tar.gz
cd node-v4.2.1
./configure
make && make install
```

## 在Centos7中安装Docker

Docker 软件包已经包含在默认的 CentOS-Extras 软件源里，安装命令如下：

```bash
tee /etc/yum.repos.d/docker.repo <<-'EOF'
[dockerrepo]
name=Docker Repository
baseurl=https://yum.dockerproject.org/repo/main/centos/$releasever/
enabled=1
gpgcheck=1
gpgkey=https://yum.dockerproject.org/gpg
EOF
yum install docker-engine
systemctl start docker #启动Docker
#验证Docker是否正常工作
docker pull centos #从仓库下载镜像
docker images #显示本地镜像
docker run -i -t centos /bin/bash #使用镜像来启动一个容器。
```

参考：

https://docs.docker.com/engine/installation/centos/
