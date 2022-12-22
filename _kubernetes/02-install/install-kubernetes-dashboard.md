---
layout: post
title: "安装 kubernetes dashboard"
---

https://github.com/kubernetes/dashboard

## Install

```bash
#kubectl apply -f https://raw.githubusercontent.com/kubernetes/dashboard/v2.7.0/aio/deploy/recommended.yaml
kubectl apply -f recommended.yaml
```

## Access

```bash
kubectl proxy --address='0.0.0.0'
 ```

http://localhost:8001/api/v1/namespaces/kubernetes-dashboard/services/https:kubernetes-dashboard:/proxy/

## 创建用户

https://github.com/kubernetes/dashboard/blob/master/docs/user/access-control/creating-sample-user.md

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
