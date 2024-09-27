---
layout: post
title: "install golang"
date: 2024-07-10
---

<https://golang.google.cn/doc/install>

## script

```bash
cd /usr/local/src
wget https://golang.google.cn/dl/go1.16.linux-amd64.tar.gz
tar -zxvf go1.16.linux-amd64.tar.gz -C /usr/local/
cd /usr/local
mv go go1.16.linux-amd64
ln -s go1.16.linux-amd64 go

cat << EOF >> ~/.profile

export PATH=\$PATH:/usr/local/go/bin
EOF

source ~/.profile
go version
```
