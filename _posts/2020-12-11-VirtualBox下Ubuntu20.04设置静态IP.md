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

```bash
ubuntu@ubuntu:~$ cat /etc/netplan/00-installer-config.yaml 
# This is the network config written by 'subiquity'
network:
  ethernets:
    enp0s3:
      dhcp4: true
    enp0s8:
      addresses:
          - 192.168.56.101/24
      gateway4: 192.168.56.1
      nameservers:
          addresses: [192.168.56.1]
  version: 2
```
## 0x02 使配置生效

```bash
ubuntu@ubuntu:~$ sudo netplan apply
```

## 0x03

Netplan configuration examples: https://netplan.io/examples/