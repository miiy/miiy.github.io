---
layout: post
title: "VirtualBox 设置静态 IP"
date: 2016-06-12
tags: virtualbox
---

## 虚拟机网络配置方式

Windows: NatNetwork 或者 共享网络 + Host-only

MacOS: NatNetwork

MacOS 使用 Host-only 模式共享网络比较复杂，推荐使用 NatNetwork，可以使用端口映射或者开一台Client虚拟机操作

文件传输示例：

```bash
ssh user@127.0.0.1 -p 2022
scp -P 2022 ~/Downloads/a.zip user@127.0.0.1:~
```

### MacOS:

VirtualBox > Preferences > Network > NatNetwork

```text
NetworkCIDR: 10.0.2.0/24
```

虚拟机 Adapter1:

```text
Attached to: NatNetwork
Name: NatNetwork
```

Debian 网络配置

```bash
root@debian:~# vi /etc/network/interfaces
```

```conf
# The primary network interface
allow-hotplug enp0s3
iface enp0s3 inet static
address 10.0.2.21
gateway 10.0.2.1
netmask 255.255.255.0
```

```bash
root@debian:~# reboot
```

### Windows

打开 Windows 网络连接设置共享，选择 VirtualBox Host-only Network

查看 VirtualBox Host-Only Ethernet Adapter

```text
IPv4 地址: 192.168.137.1
IPv4 子网掩码: 255.255.255.0
IPv4 默认网关:
IPv4 DNS 服务器:
IPv4 WINS 服务器:
```

打开 VirtualBox 主机网络管理器

网卡

```text
手动配置网卡
IPv4 地址: 192.168.137.1
IPv4 子网掩码: 255.255.255.0
```

DHCP服务器

```text
服务器地址: 192.168.137.2
服务器网络掩码: 255.255.255.0
最小地址: 192.168.137.3
最大地址: 192.168.137.254
```

Debian 网络配置

```bash
vi /etc/network/interfaces
```

```text
# The primary network interface
allow-hotplug enp0s3
address 192.168.137.10
gateway 192.168.137.1
```

```bash
reboot
```

CentOS 网络配置

```bash
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

IPADDR=192.168.137.101
NETMASK=255.255.255.0
GATEWAY=192.168.137.1

[root@localhost ~]# reboot
```
