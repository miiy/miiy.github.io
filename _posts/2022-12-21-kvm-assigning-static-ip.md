---
layout: post
title: "kvm 设置固定 ip"
date: 2022-12-21
tags: linux
---

```bash
$ sudo virsh net-list
 Name      State    Autostart   Persistent
--------------------------------------------
 default   active   yes         yes

$ sudo virsh net-dumpxml default
<network>
  <name>default</name>
  <uuid>0f6bd2e7-c497-4560-a748-b155a4000322</uuid>
  <forward mode='nat'>
    <nat>
      <port start='1024' end='65535'/>
    </nat>
  </forward>
  <bridge name='virbr0' stp='on' delay='0'/>
  <mac address='52:54:00:61:32:33'/>
  <ip address='192.168.122.1' netmask='255.255.255.0'>
    <dhcp>
      <range start='192.168.122.2' end='192.168.122.254'/>
      <host mac='52:54:00:3f:ae:7e' name='win10' ip='192.168.122.99'/>
    </dhcp>
  </ip>
</network>

$ sudo virsh net-edit default
Network default XML configuration edited.

$ sudo virsh net-edit default
Network default XML configuration edited.

$ sudo virsh net-destroy default
Network default destroyed

$ sudo virsh net-start default
Network default started

$ sudo systemctl restart libvirtd.service
$ sudo virsh net-dumpxml default
```
