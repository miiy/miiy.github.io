---
layout: post
title: "Nginx Ingress Controller"
date: 2020-12-12
tags: kubernetes
---

## 安装 Nginx Ingress Controller

<https://kubernetes.github.io/ingress-nginx/deploy>

<https://www.kubernetes.org.cn/3880.html>

Helm 用户指南-系列（5）-使用: <https://kubernetes.github.io/ingress-nginx/deploy/baremetal/>

```bash
root@k8s-master01:~# helm repo add ingress-nginx https://kubernetes.github.io/ingress-nginx
"ingress-nginx" has been added to your repositories
root@k8s-master01:~# helm repo list
NAME         	URL
ingress-nginx	https://kubernetes.github.io/ingress-nginx
root@k8s-master01:~# helm search repo ingress-nginx
NAME                       	CHART VERSION	APP VERSION	DESCRIPTION
ingress-nginx/ingress-nginx	3.27.0       	0.45.0     	Ingress controller for Kubernetes using NGINX a...
root@k8s-master01:~# helm pull ingress-nginx/ingress-nginx

kubectl create ns ingress-nginx
kubectl create secret docker-registry aliyunregcred --namespace=ingress-nginx --docker-server=registry.cn-hangzhou.aliyuncs.com --docker-username=user --docker-password=password

helm install ingress-nginx ./ingress-nginx-3.27.0.tgz \
--namespace=ingress-nginx \
--set controller.image.repository=registry.cn-hangzhou.aliyuncs.com/ahub/ingress-nginx \
--set controller.image.tag=v0.45.0-controller \
--set controller.image.digest= \
--set controller.dnsPolicy=ClusterFirstWithHostNet \
--set controller.hostNetwork=true \
--set controller.metrics.enabled=true \
--set controller.kind=DaemonSet \
--set defaultBackend.enabled=true \
--set defaultBackend.image.repository=registry.cn-hangzhou.aliyuncs.com/ahub/ingress-nginx \
--set defaultBackend.image.tag=defaultbackend-amd64 \
--set imagePullSecrets[0].name=aliyunregcred \
--dry-run --debug

kubectl get po -n ingress-nginx --watch
```

删除

```bash
helm delete ingress-nginx -n ingress-nginx
```
