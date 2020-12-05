---
layout: post
title: "VirtualBox下CentOS8设置静态IP"
date: 2016-06-12
category: linux
tags: linux virtualbox centos8
---


## 配置网络信息

1、虚拟机网络连接方式选择桥接网卡

2、查看本地网络网关

我的电脑是

```
192.168.124.1
```

3、配置网卡

```
[root@localhost ~]# cat /etc/sysconfig/network-scripts/ifcfg-enp0s3
TYPE=Ethernet
PROXY_METHOD=none
BROWSER_ONLY=no
#BOOTPROTO=dhcp
BOOTPROTO=static
DEFROUTE=yes
IPV4_FAILURE_FATAL=no
IPV6INIT=yes
IPV6_AUTOCONF=yes
IPV6_DEFROUTE=yes
IPV6_FAILURE_FATAL=no
IPV6_ADDR_GEN_MODE=stable-privacy
NAME=enp0s3
UUID=90f09831-a504-41eb-ad9e-4cd7382ba08d
DEVICE=enp0s3
#ONBOOT=no
ONBOOT=yes

IPADDR=192.168.124.101
NETMASK=255.255.255.0
GATEWAY=192.168.124.1

[root@localhost ~]# reboot
```
