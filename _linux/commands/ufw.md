---
layout: post
title: "ufw"
---

## install

```bash
$ sudo apt install ufw
$ sudo ufw enable
Firewall is active and enabled on system startup
$ sudo ufw status
Status: active
```

## usage

```bash
sudo ufw allow 22/tcp
sudo ufw allow 80/tcp
sudo ufw allow 443/tcp
sudo ufw allow from 172.18.0.0/16 to any port 8080 proto tcp
```

```bash
sudo ufw delete allow 80/tcp
```

```bash
To                         Action      From
--                         ------      ----
80                         ALLOW       Anywhere                  
443                        ALLOW       Anywhere                  
80 (v6)                    ALLOW       Anywhere (v6)             
443 (v6)                   ALLOW       Anywhere (v6) 
```
