---
layout: post
title: "安装 kubernetes dashboard"
---

<https://github.com/kubernetes/dashboard>

## 安装 helm

## Install

# 下载 kubernetes-dashboard-7.13.0.tgz 手动安装

```bash
helm upgrade --install kubernetes-dashboard ./kubernetes-dashboard-7.13.0.tgz --create-namespace --namespace kubernetes-dashboard
```

```text
Release "kubernetes-dashboard" does not exist. Installing it now.
NAME: kubernetes-dashboard
LAST DEPLOYED: Thu Aug 28 21:46:09 2025
NAMESPACE: kubernetes-dashboard
STATUS: deployed
REVISION: 1
TEST SUITE: None
NOTES:
*************************************************************************************************
*** PLEASE BE PATIENT: Kubernetes Dashboard may need a few minutes to get up and become ready ***
*************************************************************************************************

Congratulations! You have just installed Kubernetes Dashboard in your cluster.

To access Dashboard run:
  kubectl -n kubernetes-dashboard port-forward svc/kubernetes-dashboard-kong-proxy 8443:443

NOTE: In case port-forward command does not work, make sure that kong service name is correct.
      Check the services in Kubernetes Dashboard namespace using:
        kubectl -n kubernetes-dashboard get svc

Dashboard will be available at:
  https://localhost:8443
```

## Access

方式1

把本地的一个端口 (这里是 8443) 映射到集群内某个 Pod 或 Service 的端口。

```bash
kubectl -n kubernetes-dashboard port-forward svc/kubernetes-dashboard-kong-proxy 8443:443
# kubectl -n kubernetes-dashboard port-forward svc/kubernetes-dashboard-kong-proxy 8443:443 --address 0.0.0.0
```

https://localhost:8443

方式2

把所有请求转发到 Kubernetes API Server

```bash
kubectl proxy --address='0.0.0.0'
 ```

<http://localhost:8001/api/v1/namespaces/kubernetes-dashboard/services/https:kubernetes-dashboard:/proxy/>

## 创建用户

<https://github.com/kubernetes/dashboard/blob/master/docs/user/access-control/creating-sample-user.md>

```bash
cat > dashboard-admin.yaml <<EOF
apiVersion: v1
kind: ServiceAccount
metadata:
  name: admin-user
  namespace: kubernetes-dashboard
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: admin-user
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: cluster-admin
subjects:
- kind: ServiceAccount
  name: admin-user
  namespace: kubernetes-dashboard
EOF

$ kubectl apply -f dashboard-admin.yaml
serviceaccount/admin-user created
clusterrolebinding.rbac.authorization.k8s.io/admin-user created
```

查看 token

```bash
$ kubectl -n kubernetes-dashboard create token admin-user
```

不需要的时候可以停用：把副本数缩到 0

```bash
kubectl get deploy -n kubernetes-dashboard -o name \
  | xargs -n1 kubectl scale -n kubernetes-dashboard --replicas=0
```