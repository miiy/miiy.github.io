---
layout: post
title: "containerd 使用"
---

<https://github.com/containerd/containerd/blob/main/docs/getting-started.md>

# ctr

containerd 自带的低级 CLI 工具

```bash
# 查看所有命名空间
sudo ctr namespaces list
# 查看镜像
sudo ctr -n k8s.io images list
# 查看容器
sudo ctr -n k8s.io containers list
```

# nerdctl

containerd 的用户友好 CLI，语法和 docker 基本一样

用途：直接用 containerd 来跑容器

安装：<https://github.com/containerd/nerdctl>

```bash
$ sudo tar -zxvf nerdctl-2.1.3-linux-arm64.tar.gz -C /usr/local/bin
```

```bash
$ sudo nerdctl -n k8s.io images
$ sudo nerdctl -n k8s.io ps -a
```

# crictl

Kubernetes CRI (Container Runtime Interface) 的 CLI 工具

用途：调试 K8s 的 Pod 和容器

```bash
$ sudo crictl images ls
$ sudo crictl ps
```