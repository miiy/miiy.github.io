---
layout: post
title: "ssh"
---

## ssh秘钥登录

输入命令后提示都按回车

```bash
su -

root@k8s-master:~# ssh-keygen -t rsa
Generating public/private rsa key pair.
Enter file in which to save the key (/root/.ssh/id_rsa):
Created directory '/root/.ssh'.
Enter passphrase (empty for no passphrase):
Enter same passphrase again:
Your identification has been saved in /root/.ssh/id_rsa.
Your public key has been saved in /root/.ssh/id_rsa.pub.
The key fingerprint is:
SHA256:EVbXZXIDFoozyv3vCBvpUHHDr0RvXBklcybPr+CnO5o root@k8s-master
The key's randomart image is:
+---[RSA 2048]----+
|        o.. .=B+O|
|       . .o.o .%+|
|        .= *   oo|
|      . o.* = . .|
|       oSo ..=  .|
|        . +.o. . |
|       . + o. o  |
|        o +.+o   |
|         oEo+=   |
+----[SHA256]-----+
```

查看生成的秘钥

```
root@k8s-master:~# ls -l .ssh/
total 8
-rw------- 1 root root 1823 Feb 26 00:58 id_rsa
-rw-r--r-- 1 root root  397 Feb 26 00:58 id_rsa.pub
```

```bash
ssh-copy-id -i ~/.ssh/id_rsa.pub root@192.168.56.201
```

### 允许root远程登录

```bash
debian@debian:~# sudo vi /etc/ssh/sshd_config
```

找到

```text
#PermitRootLogin prohibit-password
```

在下面一行添加

```text
PermitRootLogin yes
```

```bash
root@debian:~# systemctl restart sshd
```

## Examples

### port forward

```bash
ssh debian@127.0.0.1 -p6000 -L 127.0.0.1:3399:127.0.0.1:3389
ssh -p2200 -R 9200:127.0.0.1:9200 root@192.168.0.1
ssh root@192.168.0.1 "cd /data/deploy/; ./start.sh"
```

### ssh tunnel remote mode

公网服务器

```text
root@x.x.x.x
```

内网机器

```bash
ssh -N -R 127.0.0.1:6000:127.0.0.1:22 root@x.x.x.x -p 22
```

本地机器

```bash
ssh -A -J root@x.x.x.x mac@127.0.0.1 -p 6000

scp -o ProxyJump=root@x.x.x.x -P 6000 local.txt mac@127.0.0.1:~/Downloads

rsync -avz --progress -e "ssh -J root@x.x.x.x -p 6000" local.txt mac@127.0.0.1:~/test/
```
