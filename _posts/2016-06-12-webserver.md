---
layout: post
title: "webserver"
date: 2016-06-12
category: linux
tags: linux
---

# Nginx

##CentOS7下安装Nginx

```bash
vi /etc/yum.repos.d/nginx.repo #添加源，内容如下
```

```
[nginx]
name=nginx repo
baseurl=http://nginx.org/packages/centos/7/$basearch/
gpgcheck=0
enabled=1
```

```bash
yum install nginx #安装
systemctl start nginx #启动
systemctl enable nginx #开机启动
```

参考：http://nginx.org/en/linux_packages.html

#Apache

###CentOS7下安装Apache

```
yum install httpd #安装Apache
systemctl start httpd.service #启动
systemctl enable httpd.service #开机启动
```

# Tomcat

###CentOS7下安装Tomcat

####安装JDK

下载对应的JDK版本如jdk-8u60-linux-x64.tar.gz，放到/usr/local/src目录中。

```
tar -zxvf jdk-8u60-linux-x64.tar.gz #解压
mv jdk1.8.0_60 ../ #移动到/usr/local目录中
```

配置环境变量，在/etc/profile中添加

```
JAVA_HOME=/usr/local/jdk1.8.0_60
PATH=$JAVA_HOME/bin:$PATH
CLASSPATH=$JAVA_HOME/lib/dt.jar:$JAVA_HOME/lib/tools.jar
export JAVA_HOME PATH CLASSPATH
```

```
source /etc/profile #使profile文件生效
java -version #查看java环境变量是否生效
```

java环境变量生效后将会看到如下信息

```
java version “1.8.0_60”
Java(TM) SE Runtime Environment (build 1.8.0_60-b27)
Java HotSpot(TM) 64-Bit Server VM (build 25.60-b23, mixed mode)
```


下载tomcat放到/usr/local/src目录中。

```bash
tar -zxvf apache-tomcat-8.0.26.tar.gz #解压
mv apache-tomcat-8.0.26 ../ #移动到/usr/local目录中
vi apache-tomcat-8.0.26/bin/catalina.sh #修改tomcat启动文件
```

建议在#!/bin/sh下面一行添加

CATALINA_HOME=/usr/local/apache-tomcat-8.0.26

修改默认端口将8080改为80

```
vi /usr/local/apache-tomcat-8.0.26/conf/server.xml #修改tomcat配置文件
```

```
<Connector port="80" protocol="HTTP/1.1"
connectionTimeout="20000"
redirectPort="8443" />
```

启动tomcat

```
cd /usr/local/apache-tomcat-8.0.26/bin #进入到程序目录
./startup.sh #启动tomcat
```

Nginx负载均衡

upstream myServer{
    ip_hash;
    server 42.156.140.7;
    server 222.187.254.227 weight=2;
    server 42.120.158.67;
    server 110.75.115.70;
}

server {
    listen       80;
    server_name  localhost;

    #charset koi8-r;
    #access_log  /var/log/nginx/log/host.access.log  main;

    location / {
        proxy_pass http://myServer;
        #root   /usr/share/nginx/html;
        #index  index.html index.htm;
    }

    #error_page  404              /404.html;

    # redirect server error pages to the static page /50x.html
    #
    error_page   500 502 503 504  /50x.html;
    location = /50x.html {
        root   /usr/share/nginx/html;
    }

    # proxy the PHP scripts to Apache listening on 127.0.0.1:80
    #
    #location ~ \.php$ {
    #    proxy_pass   http://127.0.0.1;
    #}

    # pass the PHP scripts to FastCGI server listening on 127.0.0.1:9000
    #
    #location ~ \.php$ {
    #    root           html;
    #    fastcgi_pass   127.0.0.1:9000;
    #    fastcgi_index  index.php;
    #    fastcgi_param  SCRIPT_FILENAME  /scripts$fastcgi_script_name;
