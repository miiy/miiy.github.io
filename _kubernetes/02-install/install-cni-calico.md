---
layout: post
title: "安装 CNI 网络插件 Calico"
---
## 什么是 Calico？

<https://docs.tigera.io/calico/latest/about/>

Calico 是一个用于网络、网络安全和可观察性的单一平台，适用于云中、本地或边缘的任何 Kubernetes 分发。无论您是刚开始使用 Kubernetes 还是大规模运营，Calico 的开源、企业和云版本都能提供您所需的网络、安全性和可观察性。

## Install Calico

Calico quickstart guide: <https://projectcalico.docs.tigera.io/getting-started/kubernetes/quickstart>

下载 tigera-operator.yaml, custom-resources.yaml

```bash
# kubectl create -f https://raw.githubusercontent.com/projectcalico/calico/v3.30.3/manifests/tigera-operator.yaml
# kubectl create -f https://raw.githubusercontent.com/projectcalico/calico/v3.30.3/manifests/custom-resources.yaml

kubectl create -f tigera-operator.yaml
kubectl create -f custom-resources.yaml
```

查看状态

```bash
$ kubectl get daemonset -n calico-system
NAME              DESIRED   CURRENT   READY   UP-TO-DATE   AVAILABLE   NODE SELECTOR            AGE
calico-node       1         1         1       1            1           kubernetes.io/os=linux   67m
csi-node-driver   1         1         1       1            1           kubernetes.io/os=linux   67m
$ kubectl get pods -n calico-system 
NAME                                       READY   STATUS    RESTARTS   AGE
calico-kube-controllers-7ddbf49bbf-7wpnv   1/1     Running   0          68m
calico-node-5kb6g                          1/1     Running   0          68m
calico-typha-58b8b7f565-6v4mj              1/1     Running   0          68m
csi-node-driver-49ggw                      2/2     Running   0          68m
goldmane-85c8f6d476-lzmcw                  1/1     Running   0          68m
whisker-7d45fc7f48-b9cfk                   2/2     Running   0          11m
```

如有错误，通过以下命令排查问题

```bash
$ kubectl get events -n calico-system
```


监控 Calico Whisker 中的网络流量

```bash
kubectl port-forward -n calico-system service/whisker 8081:8081
```
http://127.0.0.1:8081/flow-logs