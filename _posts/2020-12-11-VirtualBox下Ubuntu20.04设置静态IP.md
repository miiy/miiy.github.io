---
layout: post
title: "VirtualBox下Ubuntu20.04设置静态IP"
date: 2020-12-11
category: linux
tags: linux virtualbox ubuntu
---

## 0x00

VirtualBox 网卡设置

* vboxnet0 192.168.56.1/24

0x01 配置虚拟机网络连接方式

* 网卡1 NAT

* 网卡2 Host-only

## 0x01 配置系统网络

查看本地网络设备

```bash
ubuntu@ubuntu:~$ ip addr
1: lo: <LOOPBACK,UP,LOWER_UP> mtu 65536 qdisc noqueue state UNKNOWN group default qlen 1000
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
    inet 127.0.0.1/8 scope host lo
       valid_lft forever preferred_lft forever
    inet6 ::1/128 scope host 
       valid_lft forever preferred_lft forever
2: enp0s3: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc fq_codel state UP group default qlen 1000
    link/ether 08:00:27:2a:d3:11 brd ff:ff:ff:ff:ff:ff
    inet 10.0.2.15/24 brd 10.0.2.255 scope global dynamic enp0s3
       valid_lft 84134sec preferred_lft 84134sec
    inet6 fe80::a00:27ff:fe2a:d381/64 scope link 
       valid_lft forever preferred_lft forever
3: enp0s8: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc fq_codel state UP group default qlen 1000
    link/ether 08:00:27:57:52:11 brd ff:ff:ff:ff:ff:ff
    inet 192.168.56.101/24 brd 192.168.56.255 scope global enp0s8
       valid_lft forever preferred_lft forever
    inet6 fe80::a00:27ff:fe57:5215/64 scope link 
       valid_lft forever preferred_lft forever
```

/etc/netplan/00-installer-config.yaml

```bash
# This is the network config written by 'subiquity'
network:
  ethernets:
    enp0s3:
      dhcp4: yes
      dhcp6: yes
      routes:
          - to: 0.0.0.0/0
            via: 10.0.2.2
            metric: 0
    enp0s8:
      dhcp4: no
      dhcp6: no
      addresses:
          - 192.168.56.101/24
      gateway4: 192.168.56.1
      routes:
          - to: 192.168.56.1/24
            via: 192.168.56.1
            metric: 100
  version: 2  
```

## 0x02 使配置生效

```bash
ubuntu@ubuntu:~$ sudo netplan apply
```

不生效就重启


## 0x03

https://askubuntu.com/questions/984445/netplan-configuration-on-ubuntu-17-04-virtual-machine

Netplan configuration examples: https://netplan.io/examples/