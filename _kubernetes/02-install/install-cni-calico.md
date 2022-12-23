---
layout: post
title: "安装 CNI 网络插件 Calico"
---

## Install Calico

<https://projectcalico.docs.tigera.io/getting-started/kubernetes/quickstart>

下载 tigera-operator.yaml, custom-resources.yaml

```bash
# kubectl create -f https://raw.githubusercontent.com/projectcalico/calico/v3.24.5/manifests/tigera-operator.yaml
# kubectl create -f https://raw.githubusercontent.com/projectcalico/calico/v3.24.5/manifests/custom-resources.yaml

kubectl create -f tigera-operator.yaml
kubectl create -f custom-resources.yaml
```

检查 event 发现有这样的 warning `serviceaccount "csi-node-driver" not found`，创建 serviceaccount

```bash
$ kubectl get events -n calico-system
LAST SEEN   TYPE      REASON         OBJECT                              MESSAGE
45m         Warning   FailedCreate   daemonset/csi-node-driver           Error creating: pods "csi-node-driver-" is forbidden: error looking up service account calico-system/csi-node-driver: serviceaccount "csi-node-driver" not found

$ kubectl get daemonset -n calico-system
NAME              DESIRED   CURRENT   READY   UP-TO-DATE   AVAILABLE   NODE SELECTOR            AGE
calico-node       1         1         1       1            1           kubernetes.io/os=linux   24h
csi-node-driver   0         0         0       0            0           kubernetes.io/os=linux   24h

$ kubectl create serviceaccount csi-node-driver -n calico-system
serviceaccount/csi-node-driver created
$ kubectl get serviceaccount -n calico-system  
NAME                      SECRETS   AGE
calico-kube-controllers   0         24h
calico-node               0         24h
calico-typha              0         24h
csi-node-driver           0         10s
default                   0         24h
$ kubectl get daemonset -n calico-system
NAME              DESIRED   CURRENT   READY   UP-TO-DATE   AVAILABLE   NODE SELECTOR            AGE
calico-node       1         1         1       1            1           kubernetes.io/os=linux   24h
csi-node-driver   1         1         1       1            1           kubernetes.io/os=linux   24h
$ kubectl get pods -n calico-system 
NAME                                       READY   STATUS    RESTARTS       AGE
calico-kube-controllers-67df98bdc8-j4kpt   1/1     Running   1 (3h9m ago)   24h
calico-node-vx68q                          1/1     Running   1 (3h9m ago)   24h
calico-typha-657c94c555-8tdpk              1/1     Running   1 (3h9m ago)   24h
csi-node-driver-xtmr8                      2/2     Running   0              73s
```