---
layout: post
title: "安装 Helm"
date: 2020-12-12
tags: kubernetes
---

## Install

<https://helm.sh/zh/docs/intro/install/>

```bash
cd /usr/local/src
wget https://get.helm.sh/helm-v3.5.2-linux-amd64.tar.gz
tar -zxvf helm-v3.5.2-linux-amd64.tar.gz
mv linux-amd64/ helm-v3.5.2-linux-amd64
mv helm-v3.5.2-linux-amd64 ../
cd ..
ln -s helm-v3.5.2-linux-amd64/ helm
echo export PATH=/usr/local/helm:\$PATH >> ~/.profile
source ~/.profile

helm help
```
