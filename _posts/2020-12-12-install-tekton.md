---
layout: post
title: "安装 Tekton"
date: 2020-12-12
tags: kubernetes
---

## 安装 Tekton

<https://github.com/tektoncd/pipeline/blob/main/docs/install.md>

```bash
kubectl apply --filename https://storage.googleapis.com/tekton-releases/pipeline/latest/release.yaml
```

gcr.io 不可访问，status 状态为 ErrImagePull，手动拉取镜像

将ImagePullSecrets添加到service account

https://kubernetes.io/docs/tasks/configure-pod-container/pull-image-private-registry/

https://kubernetes.io/docs/tasks/configure-pod-container/configure-service-account/#add-imagepullsecrets-to-a-service-account

```bash
kubectl create secret docker-registry aliyunregcred --namespace=tekton-pipelines --docker-server=registry.cn-hangzhou.aliyuncs.com --docker-username=user --docker-password=password

kubectl get secrets aliyunregcred -n tekton-pipelines
kubectl describe secrets aliyunregcred -n tekton-pipelines

kubectl patch serviceaccount tekton-pipelines-controller -p '{"imagePullSecrets": [{"name": "aliyunregcred"}]}' -n tekton-pipelines
kubectl patch serviceaccount tekton-pipelines-webhook -p '{"imagePullSecrets": [{"name": "aliyunregcred"}]}' -n tekton-pipelines

kubectl get serviceaccounts tekton-pipelines-controller -n tekton-pipelines -o yaml
kubectl get serviceaccounts tekton-pipelines-webhook -n tekton-pipelines -o yaml
```

监控 pod 运行状态，直至所有组件显示 running

```bash
kubectl get pods --namespace tekton-pipelines --watch
```

如果Pod状态为ErrImagePull，删除pod

## 安装 CLI

https://github.com/tektoncd/cli/releases

```bash
cd /usr/local/src
wget https://github.com/tektoncd/cli/releases/download/v0.17.1/tkn_0.17.1_Linux_x86_64.tar.gz
mkdir mv tkn_0.17.1_Linux_x86_64
tar -zxvf tkn_0.17.1_Linux_x86_64.tar.gz -C tkn_0.17.1_Linux_x86_64
mv tkn_0.17.1_Linux_x86_64 ../
cd ../
ln -s tkn_0.17.1_Linux_x86_64/ tkn
echo "export PATH=\$PATH:/usr/local/tkn" >> ~/.profile
```

## 安装 Tekton Dashboard

https://github.com/tektoncd/dashboard/blob/main/docs/install.md

```bash
kubectl apply -f tekton-dashboard-release.yaml

kubectl patch serviceaccount tekton-dashboard -p '{"imagePullSecrets": [{"name": "aliyunregcred"}]}' -n tekton-pipelines
```

通过 ingress 访问 dashboard

tekton-dashboard-ingress.yaml

```yaml
kubectl apply -f - <<EOF
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: tekton-dashboard
  namespace: tekton-pipelines
  annotations:
    kubernetes.io/ingress.class: "nginx"
    nginx.ingress.kubernetes.io/rewrite-target: /
spec:
  rules:
  - host: "tekton-dashboard.k8s"
    http:
      paths:
      - pathType: Prefix
        path: "/"
        backend:
          service:
            name: tekton-dashboard
            port:
              number: 9097
EOF
```

```bash
curl -D- http://10.0.2.21:32430 -H 'Host: tekton-dashboard.k8s'
```

## 安装 Tekton Triggers

https://tekton.dev/docs/triggers/install/

```bash
kubectl apply -f tekton-triggers-release.yaml

kubectl patch serviceaccount tekton-triggers-webhook -p '{"imagePullSecrets": [{"name": "aliyunregcred"}]}' -n tekton-pipelines
kubectl patch serviceaccount tekton-triggers-controller -p '{"imagePullSecrets": [{"name": "aliyunregcred"}]}' -n tekton-pipelines
kubectl patch serviceaccount tekton-triggers-core-interceptors -p '{"imagePullSecrets": [{"name": "aliyunregcred"}]}' -n tekton-pipelines
```
