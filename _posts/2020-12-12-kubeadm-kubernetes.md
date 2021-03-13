---
layout: post
title: "使用kubeadm快速安装kubernetes集群"
date: 2020-12-12
tags: kubernetes
---

## 系统要求

https://kubernetes.io/docs/setup/production-environment/tools/kubeadm/install-kubeadm/#before-you-begin

## 规划

```
k8s-soft     10.0.2.20 1C1G nfs-kernel-server, docker, harbor
k8s-master01 10.0.2.21 2C2G docker, kubeadm、kubelet、kubectl, nfs-common, k9s, helm
k8s-node01   10.0.2.22 2C2G docker, kubeadm、kubelet、kubectl, nfs-common
k8s-node02   20.0.2.23 2C2G docker, kubeadm、kubelet、kubectl, nfs-common
```

## 系统初始化

k8s-master01, k8s-node01, k8s-node02 执行：

配置每台机器的网络

配置用户 sudo 权限

禁用 swap

设置 hostname

安装 docker

安装 kubeadm、kubelet 和 kubectl

```bash
sudo apt-get update && sudo apt-get install -y apt-transport-https curl
curl -s https://mirrors.aliyun.com/kubernetes/apt/doc/apt-key.gpg | sudo apt-key add -
cat <<EOF | sudo tee /etc/apt/sources.list.d/kubernetes.list
deb https://mirrors.aliyun.com/kubernetes/apt kubernetes-xenial main
EOF
sudo apt-get update
sudo apt-get install -y kubelet kubeadm kubectl
sudo apt-mark hold kubelet kubeadm kubectl
```

## 配置 master 节点

设置 hosts

```bash
vi /etc/hosts
```

```text
10.0.2.21       k8s-master01
10.0.2.22       k8s-node01
10.0.2.23       k8s-node02
```

### 允许 iptables 检查桥接流量

https://kubernetes.io/zh/docs/setup/production-environment/tools/kubeadm/install-kubeadm/

```bash
cat <<EOF | sudo tee /etc/modules-load.d/k8s.conf
br_netfilter
EOF

cat <<EOF | sudo tee /etc/sysctl.d/k8s.conf
net.bridge.bridge-nf-call-ip6tables = 1
net.bridge.bridge-nf-call-iptables = 1
EOF
sudo sysctl --system
```

### 初始化 master 节点

```bash
su -
kubeadm init \
--apiserver-advertise-address=10.0.2.21 \
--image-repository registry.aliyuncs.com/google_containers \
--kubernetes-version v1.20.4 \
--service-cidr=10.96.0.0/12 \
--pod-network-cidr=10.244.0.0/16

echo export KUBECONFIG=/etc/kubernetes/admin.conf >> ~/.profile
```

示例

```bash
debian@k8s-master01:~# su -
root@k8s-master01:~# kubeadm init \
> --apiserver-advertise-address=10.0.2.21 \
> --image-repository registry.aliyuncs.com/google_containers \
> --kubernetes-version v1.20.4 \
> --service-cidr=10.96.0.0/12 \
> --pod-network-cidr=10.244.0.0/16
[init] Using Kubernetes version: v1.20.4
[preflight] Running pre-flight checks
	[WARNING IsDockerSystemdCheck]: detected "cgroupfs" as the Docker cgroup driver. The recommended driver is "systemd". Please follow the guide at https://kubernetes.io/docs/setup/cri/
	[WARNING SystemVerification]: this Docker version is not on the list of validated versions: 20.10.5. Latest validated version: 19.03
	[WARNING SystemVerification]: missing optional cgroups: hugetlb
[preflight] Pulling images required for setting up a Kubernetes cluster
[preflight] This might take a minute or two, depending on the speed of your internet connection
[preflight] You can also perform this action in beforehand using 'kubeadm config images pull'
[certs] Using certificateDir folder "/etc/kubernetes/pki"
[certs] Generating "ca" certificate and key
[certs] Generating "apiserver" certificate and key
[certs] apiserver serving cert is signed for DNS names [k8s-master01 kubernetes kubernetes.default kubernetes.default.svc kubernetes.default.svc.cluster.local] and IPs [10.96.0.1 10.0.2.21]
[certs] Generating "apiserver-kubelet-client" certificate and key
[certs] Generating "front-proxy-ca" certificate and key
[certs] Generating "front-proxy-client" certificate and key
[certs] Generating "etcd/ca" certificate and key
[certs] Generating "etcd/server" certificate and key
[certs] etcd/server serving cert is signed for DNS names [k8s-master01 localhost] and IPs [10.0.2.21 127.0.0.1 ::1]
[certs] Generating "etcd/peer" certificate and key
[certs] etcd/peer serving cert is signed for DNS names [k8s-master01 localhost] and IPs [10.0.2.21 127.0.0.1 ::1]
[certs] Generating "etcd/healthcheck-client" certificate and key
[certs] Generating "apiserver-etcd-client" certificate and key
[certs] Generating "sa" key and public key
[kubeconfig] Using kubeconfig folder "/etc/kubernetes"
[kubeconfig] Writing "admin.conf" kubeconfig file
[kubeconfig] Writing "kubelet.conf" kubeconfig file
[kubeconfig] Writing "controller-manager.conf" kubeconfig file
[kubeconfig] Writing "scheduler.conf" kubeconfig file
[kubelet-start] Writing kubelet environment file with flags to file "/var/lib/kubelet/kubeadm-flags.env"
[kubelet-start] Writing kubelet configuration to file "/var/lib/kubelet/config.yaml"
[kubelet-start] Starting the kubelet
[control-plane] Using manifest folder "/etc/kubernetes/manifests"
[control-plane] Creating static Pod manifest for "kube-apiserver"
[control-plane] Creating static Pod manifest for "kube-controller-manager"
[control-plane] Creating static Pod manifest for "kube-scheduler"
[etcd] Creating static Pod manifest for local etcd in "/etc/kubernetes/manifests"
[wait-control-plane] Waiting for the kubelet to boot up the control plane as static Pods from directory "/etc/kubernetes/manifests". This can take up to 4m0s
[kubelet-check] Initial timeout of 40s passed.
[apiclient] All control plane components are healthy after 84.503062 seconds
[upload-config] Storing the configuration used in ConfigMap "kubeadm-config" in the "kube-system" Namespace
[kubelet] Creating a ConfigMap "kubelet-config-1.20" in namespace kube-system with the configuration for the kubelets in the cluster
[upload-certs] Skipping phase. Please see --upload-certs
[mark-control-plane] Marking the node k8s-master01 as control-plane by adding the labels "node-role.kubernetes.io/master=''" and "node-role.kubernetes.io/control-plane='' (deprecated)"
[mark-control-plane] Marking the node k8s-master01 as control-plane by adding the taints [node-role.kubernetes.io/master:NoSchedule]
[bootstrap-token] Using token: 3o6jxw.y18zwlytyjvah276
[bootstrap-token] Configuring bootstrap tokens, cluster-info ConfigMap, RBAC Roles
[bootstrap-token] configured RBAC rules to allow Node Bootstrap tokens to get nodes
[bootstrap-token] configured RBAC rules to allow Node Bootstrap tokens to post CSRs in order for nodes to get long term certificate credentials
[bootstrap-token] configured RBAC rules to allow the csrapprover controller automatically approve CSRs from a Node Bootstrap Token
[bootstrap-token] configured RBAC rules to allow certificate rotation for all node client certificates in the cluster
[bootstrap-token] Creating the "cluster-info" ConfigMap in the "kube-public" namespace
[kubelet-finalize] Updating "/etc/kubernetes/kubelet.conf" to point to a rotatable kubelet client certificate and key
[addons] Applied essential addon: CoreDNS
[addons] Applied essential addon: kube-proxy

Your Kubernetes control-plane has initialized successfully!

To start using your cluster, you need to run the following as a regular user:

  mkdir -p $HOME/.kube
  sudo cp -i /etc/kubernetes/admin.conf $HOME/.kube/config
  sudo chown $(id -u):$(id -g) $HOME/.kube/config

Alternatively, if you are the root user, you can run:

  export KUBECONFIG=/etc/kubernetes/admin.conf

You should now deploy a pod network to the cluster.
Run "kubectl apply -f [podnetwork].yaml" with one of the options listed at:
  https://kubernetes.io/docs/concepts/cluster-administration/addons/

Then you can join any number of worker nodes by running the following on each as root:

kubeadm join 10.0.2.21:6443 --token 3o6jxw.y18zwlytyjvah276 \
    --discovery-token-ca-cert-hash sha256:48387cbe4a3e4eb7b68c71cd24279d074d4ebc66928acf33c0ebcfe403842bcc 
root@k8s-master01:~# 
root@k8s-master01:~# echo export KUBECONFIG=/etc/kubernetes/admin.conf >> ~/.bashrc
root@k8s-master01:~# 
root@k8s-master01:~# kubectl get nodes -o wide
NAME           STATUS     ROLES                  AGE   VERSION   INTERNAL-IP   EXTERNAL-IP   OS-IMAGE                       KERNEL-VERSION    CONTAINER-RUNTIME
k8s-master01   NotReady   control-plane,master   62s   v1.20.4   10.0.2.21     <none>        Debian GNU/Linux 10 (buster)   4.19.0-14-amd64   docker://20.10.5
root@k8s-master01:~# 
root@k8s-master01:~# docker images
REPOSITORY                                                        TAG        IMAGE ID       CREATED         SIZE
registry.aliyuncs.com/google_containers/kube-proxy                v1.20.4    c29e6c583067   2 weeks ago     118MB
registry.aliyuncs.com/google_containers/kube-apiserver            v1.20.4    ae5eb22e4a9d   2 weeks ago     122MB
registry.aliyuncs.com/google_containers/kube-controller-manager   v1.20.4    0a41a1414c53   2 weeks ago     116MB
registry.aliyuncs.com/google_containers/kube-scheduler            v1.20.4    5f8cb769bd73   2 weeks ago     47.3MB
registry.aliyuncs.com/google_containers/etcd                      3.4.13-0   0369cf4303ff   6 months ago    253MB
registry.aliyuncs.com/google_containers/coredns                   1.7.0      bfe3a36ebd25   8 months ago    45.2MB
registry.aliyuncs.com/google_containers/pause                     3.2        80d28bedfe5d   12 months ago   683kB
```

## 配置 node 节点

向集群添加新节点，在 k8s-node01, k8s-node02 执行 kubeadm init 输出的 kubeadm join 命令

```bash
su -
kubeadm join 10.0.2.21:6443 --token 3o6jxw.y18zwlytyjvah276 \
    --discovery-token-ca-cert-hash sha256:48387cbe4a3e4eb7b68c71cd24279d074d4ebc66928acf33c0ebcfe403842bcc 
```

默认 token 有效期为24小时，重新创建 token 命令如下：

```bash
root@k8s-master01:~# kubeadm token create --print-join-command
kubeadm join 10.0.2.21:6443 --token r8s55a.9iezxncq5h29vrpe     --discovery-token-ca-cert-hash sha256:48387cbe4a3e4eb7b68c71cd24279d074d4ebc66928acf33c0ebcfe403842bcc 
```

## 安装 CNI 网络插件

在 k8s-master01 上执行

```bash
root@k8s-master01:~# kubectl get nodes -o wide
NAME           STATUS     ROLES                  AGE     VERSION   INTERNAL-IP   EXTERNAL-IP   OS-IMAGE                       KERNEL-VERSION    CONTAINER-RUNTIME
k8s-master01   NotReady   control-plane,master   3m40s   v1.20.4   10.0.2.21     <none>        Debian GNU/Linux 10 (buster)   4.19.0-14-amd64   docker://20.10.5
k8s-node01     NotReady   <none>                 5s      v1.20.4   10.0.2.22     <none>        Debian GNU/Linux 10 (buster)   4.19.0-14-amd64   docker://20.10.5
k8s-node02     NotReady   <none>                 2s      v1.20.4   10.0.2.23     <none>        Debian GNU/Linux 10 (buster)   4.19.0-14-amd64   docker://20.10.5
root@k8s-master01:~#
root@k8s-master01:~# kubectl apply -f https://raw.githubusercontent.com/coreos/flannel/master/Documentation/kube-flannel.yml
podsecuritypolicy.policy/psp.flannel.unprivileged created
clusterrole.rbac.authorization.k8s.io/flannel created
clusterrolebinding.rbac.authorization.k8s.io/flannel created
serviceaccount/flannel created
configmap/kube-flannel-cfg created
daemonset.apps/kube-flannel-ds created
root@k8s-master01:~# kubectl get pods -n kube-system
NAME                                   READY   STATUS              RESTARTS   AGE
coredns-7f89b7bc75-h8xdb               0/1     Pending             0          113s
coredns-7f89b7bc75-tgzvj               0/1     Pending             0          113s
etcd-k8s-master01                      1/1     Running             0          2m10s
kube-apiserver-k8s-master01            1/1     Running             0          2m10s
kube-controller-manager-k8s-master01   1/1     Running             0          2m10s
kube-flannel-ds-8sb8b                  0/1     Init:0/1            0          22s
kube-flannel-ds-drpb6                  0/1     Init:0/1            0          22s
kube-flannel-ds-w7fcc                  0/1     Init:0/1            0          22s
kube-proxy-q6bnv                       0/1     ContainerCreating   0          80s
kube-proxy-sd6zz                       1/1     Running             0          113s
kube-proxy-sm5cm                       0/1     ContainerCreating   0          69s
kube-scheduler-k8s-master01            1/1     Running             0          2m10s
```

## 安装 Kubernetes Dashboard

https://github.com/kubernetes/dashboard/blob/master/docs/user/installation.md

在 k8s-master01 上执行

```bash
root@k8s-master01:~# wget https://raw.githubusercontent.com/kubernetes/dashboard/v2.0.0/aio/deploy/recommended.yaml
root@k8s-master01:~# mv recommended.yaml dashboard-recommended.yaml
```

访问 Dashboard

https://github.com/kubernetes/dashboard/tree/master/docs/user/accessing-dashboard

使用 kube proxy 默认它只能在本地（从启动它的机器）访问

新增type: NodePort 和 nodePort：30443，以便能实现非本机访问

```bash
root@k8s-master01:~# vi dashboard-recommended.yaml
```

```yaml
kind: Service
apiVersion: v1
metadata:
  labels:
    k8s-app: kubernetes-dashboard
  name: kubernetes-dashboard
  namespace: kubernetes-dashboard
spec:
  type: NodePort # 新增
  ports:
    - port: 443
      targetPort: 8443
      nodePort: 30443  # 新增
  selector:
    k8s-app: kubernetes-dashboard
```

```bash
root@k8s-master01:~# kubectl apply -f dashboard-recommended.yaml
namespace/kubernetes-dashboard created
serviceaccount/kubernetes-dashboard created
service/kubernetes-dashboard created
secret/kubernetes-dashboard-certs created
secret/kubernetes-dashboard-csrf created
secret/kubernetes-dashboard-key-holder created
configmap/kubernetes-dashboard-settings created
role.rbac.authorization.k8s.io/kubernetes-dashboard created
clusterrole.rbac.authorization.k8s.io/kubernetes-dashboard created
rolebinding.rbac.authorization.k8s.io/kubernetes-dashboard created
clusterrolebinding.rbac.authorization.k8s.io/kubernetes-dashboard created
deployment.apps/kubernetes-dashboard created
service/dashboard-metrics-scraper created
deployment.apps/dashboard-metrics-scraper created
root@k8s-master01:~#
root@k8s-master01:~# kubectl get pods,svc -n kubernetes-dashboard
NAME                                             READY   STATUS              RESTARTS   AGE
pod/dashboard-metrics-scraper-7b59f7d4df-859xz   0/1     ContainerCreating   0          14s
pod/kubernetes-dashboard-74d688b6bc-v67wh        0/1     ContainerCreating   0          14s

NAME                                TYPE        CLUSTER-IP       EXTERNAL-IP   PORT(S)         AGE
service/dashboard-metrics-scraper   ClusterIP   10.98.204.59     <none>        8000/TCP        14s
service/kubernetes-dashboard        NodePort    10.104.183.246   <none>        443:30443/TCP   14s
```

创建 Dashboard 管理员

https://github.com/kubernetes/dashboard/blob/master/docs/user/access-control/creating-sample-user.md

```bash
root@k8s-master01:~# cat > dashboard-admin.yaml <<EOF
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
root@k8s-master01:~# kubectl apply -f dashboard-admin.yaml
serviceaccount/admin-user created
clusterrolebinding.rbac.authorization.k8s.io/admin-user created
```

获取 Token

```bash
kubectl -n kubernetes-dashboard get secret $(kubectl -n kubernetes-dashboard get sa/admin-user -o jsonpath="{.secrets[0].name}") -o go-template="{{.data.token | base64decode}}"
```

打开 https://NodeIp:30443/

Chrome 提示：Your connection is not private，用键盘输入: thisisunsafe

## 测试kubernetes集群

在 Kubernetes 集群中创建一个 deployment，验证是否正常运行

```bash
root@k8s-master01:~# kubectl create deployment nginx --image=nginx
deployment.apps/nginx created
root@k8s-master01:~# kubectl expose deployment nginx --port=80 --type=NodePort
service/nginx exposed
root@k8s-master01:~# kubectl get pod,svc
NAME                         READY   STATUS              RESTARTS   AGE
pod/nginx-6799fc88d8-mqdc4   0/1     ContainerCreating   0          12s

NAME                 TYPE        CLUSTER-IP      EXTERNAL-IP   PORT(S)        AGE
service/kubernetes   ClusterIP   10.96.0.1       <none>        443/TCP        20m
service/nginx        NodePort    10.100.81.161   <none>        80:31927/TCP   5s
```

任一节点访问：http://NodeIP:Port

curl http://10.0.2.21:31927

## K9s

https://k9scli.io/topics/install/

```bash
cd /usr/local/src
mkdir k9s_v0.24.2
tar -zxvf k9s_Linux_x86_64.tar.gz -C /usr/local/src/k9s_v0.24.2/
mv k9s_v0.24.2/ ../
cd ../
ln -s k9s_v0.24.2/ k9s
echo export PATH=/usr/local/k9s:\$PATH >> ~/.profile
source ~/.profile
```
