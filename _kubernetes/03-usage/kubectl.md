---
layout: post
title: "命令行工具 kubectl"
---

<https://kubernetes.io/zh-cn/docs/reference/kubectl/>

```bash
# 创建资源
kubectl apply -f <name.yaml>
# 应用新的 YAML 文件
kubectl apply -f <name.yaml>
# 查看节点
kubectl get nodes
# 查看版本信息
kubectl get version
kubectl get cs
kubectl get services
# 查看 deployments
kubectl get deployments
kubectl get pods
# 列出该 Deployment 创建的 Pod：
kubectl get pods -l app=nginx
kubectl get pods -o wide
# 查询所有命名空间下的pod
kubectl get po -A
# 查询所有命名空间下image运行起来的资源
kubectl get all -o wide -A
# 查看 ReplicaSet 信息
kubectl get rs
# 显示 Pod 信息
kubectl describe pod <pod-name>
# 显示 Deployment 的信息
kubectl describe deployment <deployment-name>
# 查看 pod 日志
kubectl logs kubernetes-dashboard-6c7ccbcf87-nfv8t -n kubernetes-dashboard
kubectl scale deployments/nginx-deployment --replicas=2
kubectl expose deployments/nginx-deployment --type="NodePort" --port 8080
# 删除 Deployment
kubectl delete deployment <deployment-name>
kubectl delete services/nginx-deployment
#
kubectl exec -n default nginx-deployment-66f8758855-l4s89 -- echo hello
kubectl exec -n default nginx-deployment-66f8758855-l4s89 -- bash -c 'ls /usr'
kubectl exec -n default -it nginx-deployment-66f8758855-l4s89 -- bash
```