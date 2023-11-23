---
layout: post
title: "nmcli"
date: 2023-11-13
tags: linux
---

控制 NetworkManager 的命令行工具

查看设备状态

```bash
$ sudo nmcli device status
DEVICE           TYPE      STATE                   CONNECTION      
wlp2s0           wifi      connected               cmcc-101      
br-5ec2eb60d1a0  bridge    connected (externally)  br-5ec2eb60dda0 
lo               loopback  connected (externally)  lo              
docker0          bridge    connected (externally)  docker0         
```

查看所有连接

```bash
$ sudo nmcli connection show
NAME                UUID                                  TYPE      DEVICE          
cmcc-101            74902b90-1234-4669-8ee8-7fabc12334a1  wifi      wlp2s0          
br-5ec2eb60d1a0     175387fa-45c6-1234-a7ca-d0123c32b9c4  bridge    br-5ec2eb60dda0 
lo                  ea123692-1234-46b3-a106-af4a9fd9a605  loopback  lo              
docker0             73e11235-9652-45c6-1234-7b86c1231525  bridge    docker0         
Wired connection 1  3fb94afd-1234-49b7-b87c-b5304b123685  ethernet  --              
cmcc-101-5G         77b123e0-f9a9-1234-b7ef-46cc123ba97e  wifi      --          
```

wifi 列表

```bash
$ sudo nmcli dev wifi
IN-USE  BSSID              SSID                      MODE   CHAN  RATE        SIGNAL  BARS  SECURITY  
        52:33:F0:9F:32:04  --                        Infra  11    130 Mbit/s  100     ▂▄▆█  WPA2      
*       50:33:F0:BF:32:06  cmcc-101-5G               Infra  149   270 Mbit/s  100     ▂▄▆█  WPA1 WPA2 
        52:33:F0:AF:32:06  cmcc-5G-6                 Infra  149   270 Mbit/s  100     ▂▄▆█  WPA2      
```

连接 wifi

```bash
sudo nmcli dev wifi connect cmcc-101-5G
sudo nmcli dev wifi connect "testwifi-202" password "123456789"
```

https://access.redhat.com/documentation/zh-cn/red_hat_enterprise_linux/9/html/configuring_and_managing_networking/proc_connecting-to-a-wifi-network-by-using-nmcli_assembly_managing-wifi-connections

https://www.cnblogs.com/linxmouse/p/15938912.html