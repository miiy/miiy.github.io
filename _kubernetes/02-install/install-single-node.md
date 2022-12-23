---
layout: post
title: "使用 kubeadm 安装 kubernetes 单节点学习环境"
---

参考安装文档：

<https://kubernetes.io/zh-cn/docs/setup/production-environment/tools/kubeadm/install-kubeadm/>

## 基础设置

```bash
# 设置 hostname
sudo hostnamectl set-hostname "k8s-master"
# 修改 hosts
echo -e "\n192.168.122.99	k8s-master" |sudo tee -a /etc/hosts
# 禁用 swap
sudo swapoff -a
sudo sed -ri 's/.*swap.*/#&/' /etc/fstab
```

## 安装容器运行时

### 使用 apt 安装 containerd

<https://docs.docker.com/engine/install/debian/>

```bash
sudo apt-get update

sudo apt-get install \
    ca-certificates \
    curl \
    gnupg \
    lsb-release

sudo mkdir -p /etc/apt/keyrings

curl -fsSL https://download.docker.com/linux/debian/gpg | sudo gpg --dearmor -o /etc/apt/keyrings/docker.gpg

echo \
  "deb [arch=$(dpkg --print-architecture) signed-by=/etc/apt/keyrings/docker.gpg] https://download.docker.com/linux/debian \
  $(lsb_release -cs) stable" | sudo tee /etc/apt/sources.list.d/docker.list > /dev/null

sudo apt-get update

sudo apt-get install containerd.io
```

修改 conatinerd 配置文件

这里生成 containerd 默认配置文件修改，如直接编辑 config.toml 文件，在 kubeadm init 会报错` Dial to tcp:10.0.2.20:6443 failed: dial tcp 10.0.2.20:6443: connect: connection refused`，暂未找到解决方案

```bash
$ containerd config dump|grep sandbox_image
    sandbox_image = "registry.k8s.io/pause:3.6"

$ sudo mv /etc/containerd/config.toml /etc/containerd/config.toml.origin
$ containerd config default|sudo tee /etc/containerd/config.toml
$ sudo vi /etc/containerd/config.toml
```

```conf
version = 2
disabled_plugins = []

[plugins."io.containerd.grpc.v1.cri"]
    sandbox_image = "registry.aliyuncs.com/google_containers/pause:3.6"

[plugins."io.containerd.grpc.v1.cri".containerd.runtimes.runc.options]
    SystemdCgroup = true

[plugins."io.containerd.grpc.v1.cri".registry.mirrors."docker.io"]
    endpoint = ["https://hub-mirror.c.163.com"]

[plugins."io.containerd.grpc.v1.cri".registry.mirrors."quay.io"]
    endpoint = ["https://quay.mirrors.ustc.edu.cn"]
```

```bash
sudo systemctl restart containerd
```

查看配置是否生效

```bash
containerd config dump
```

> 说明: 配置的是插件`io.containerd.grpc.v1.cri`的镜像地址，使用crictl拉取镜像文件才有效：`sudo crictl --debug pull docker.io/library/golang:1.9`

## 安装 kubeadm、kubelet 和 kubectl

```bash
sudo apt-get update
sudo apt-get install -y apt-transport-https ca-certificates curl

#sudo curl -fsSLo /etc/apt/keyrings/kubernetes-archive-keyring.gpg https://packages.cloud.google.com/apt/doc/apt-key.gpg
sudo curl -fsSLo /etc/apt/keyrings/kubernetes-archive-keyring.gpg https://mirrors.aliyun.com/kubernetes/apt/doc/apt-key.gpg

#echo "deb [signed-by=/etc/apt/keyrings/kubernetes-archive-keyring.gpg] https://apt.kubernetes.io/ kubernetes-xenial main" | sudo tee /etc/apt/sources.list.d/kubernetes.list
echo "deb [signed-by=/etc/apt/keyrings/kubernetes-archive-keyring.gpg] https://mirrors.aliyun.com/kubernetes/apt kubernetes-xenial main" | sudo tee /etc/apt/sources.list.d/kubernetes.list

sudo apt-get update
sudo apt-get install -y kubelet kubeadm kubectl
sudo apt-mark hold kubelet kubeadm kubectl
```

## 转发 IPv4 并让 iptables 看到桥接流量

<https://kubernetes.io/zh-cn/docs/setup/production-environment/container-runtimes/>

```bash
cat <<EOF | sudo tee /etc/modules-load.d/k8s.conf
overlay
br_netfilter
EOF

sudo modprobe overlay
sudo modprobe br_netfilter

# 设置所需的 sysctl 参数，参数在重新启动后保持不变
cat <<EOF | sudo tee /etc/sysctl.d/k8s.conf
net.bridge.bridge-nf-call-iptables  = 1
net.bridge.bridge-nf-call-ip6tables = 1
net.ipv4.ip_forward                 = 1
EOF

# 应用 sysctl 参数而不重新启动
sudo sysctl --system
```

## 使用 kubeadm 创建集群

```bash
sudo kubeadm init \
--control-plane-endpoint=k8s-master \
--image-repository registry.aliyuncs.com/google_containers \
--kubernetes-version v1.26.0 \
--pod-network-cidr=192.168.0.0/16
```

如果出错，加`--v=9`输出详细信息

参数说明：

- --control-plane-endpoint string 为控制平面指定一个稳定的 IP 地址或 DNS 名称。
- --image-repository 默认值："registry.k8s.io"，选择用于拉取控制平面镜像的容器仓库
- --kubernetes-version 为控制平面选择一个特定的 Kubernetes 版本。
- --service-cidr 默认值："10.96.0.0/12"，为服务的虚拟 IP 地址另外指定 IP 地址段
- --pod-network-cidr 指明 pod 网络可以使用的 IP 地址段。如果设置了这个参数，控制平面将会为每一个节点自动分配 CIDRs。

<https://kubernetes.io/zh-cn/docs/reference/setup-tools/kubeadm/kubeadm-init/>

以下是 kubeadm init 的过程

```bash
$ kubeadm init --control-plane-endpoint=k8s-master --image-repository registry.aliyuncs.com/google_containers --kubernetes-version v1.26.0 --pod-network-cidr=192.168.0.0/16 --v=9
[init] Using Kubernetes version: v1.26.0
[preflight] Running pre-flight checks
[preflight] Pulling images required for setting up a Kubernetes cluster
[preflight] This might take a minute or two, depending on the speed of your internet connection
[preflight] You can also perform this action in beforehand using 'kubeadm config images pull'
[certs] Using certificateDir folder "/etc/kubernetes/pki"
[certs] Generating "ca" certificate and key
[certs] Generating "apiserver" certificate and key
[certs] apiserver serving cert is signed for DNS names [k8s-master kubernetes kubernetes.default kubernetes.default.svc kubernetes.default.svc.cluster.local] and IPs [10.96.0.1 10.0.2.20]
[certs] Generating "apiserver-kubelet-client" certificate and key
[certs] Generating "front-proxy-ca" certificate and key
[certs] Generating "front-proxy-client" certificate and key
[certs] Generating "etcd/ca" certificate and key
[certs] Generating "etcd/server" certificate and key
[certs] etcd/server serving cert is signed for DNS names [k8s-master localhost] and IPs [10.0.2.20 127.0.0.1 ::1]
[certs] Generating "etcd/peer" certificate and key
[certs] etcd/peer serving cert is signed for DNS names [k8s-master localhost] and IPs [10.0.2.20 127.0.0.1 ::1]
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
[apiclient] All control plane components are healthy after 30.935556 seconds
[upload-config] Storing the configuration used in ConfigMap "kubeadm-config" in the "kube-system" Namespace
[kubelet] Creating a ConfigMap "kubelet-config" in namespace kube-system with the configuration for the kubelets in the cluster
[upload-certs] Skipping phase. Please see --upload-certs
[mark-control-plane] Marking the node k8s-master as control-plane by adding the labels: [node-role.kubernetes.io/control-plane node.kubernetes.io/exclude-from-external-load-balancers]
[mark-control-plane] Marking the node k8s-master as control-plane by adding the taints [node-role.kubernetes.io/control-plane:NoSchedule]
[kubelet-check] Initial timeout of 40s passed.
[bootstrap-token] Using token: 4y8q06.72p867r3g2bxmgyq
[bootstrap-token] Configuring bootstrap tokens, cluster-info ConfigMap, RBAC Roles
[bootstrap-token] Configured RBAC rules to allow Node Bootstrap tokens to get nodes
[bootstrap-token] Configured RBAC rules to allow Node Bootstrap tokens to post CSRs in order for nodes to get long term certificate credentials
[bootstrap-token] Configured RBAC rules to allow the csrapprover controller automatically approve CSRs from a Node Bootstrap Token
[bootstrap-token] Configured RBAC rules to allow certificate rotation for all node client certificates in the cluster
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

kubeadm join k8s-master:6443 --token 4y8q06.72p867r3g2bxmgyq \
	--discovery-token-ca-cert-hash sha256:2f12d7005de5006cead31f2fa389c4195bfd16c6b1d50970e48a91e55da5d65a 
```

root 用户

```bash
echo export KUBECONFIG=/etc/kubernetes/admin.conf >> ~/.bashrc
source ~/.bashrc
```

普通用户

```bash
mkdir -p $HOME/.kube
sudo cp -i /etc/kubernetes/admin.conf $HOME/.kube/config
sudo chown $(id -u):$(id -g) $HOME/.kube/config
```

## 移除 master 污点

由于装的是单节点的学习环境，需要移除 master 的污点，pod 才可以调度到 master 节点

```bash
$ kubectl describe node|grep Taints
Taints:             node-role.kubernetes.io/control-plane:NoSchedule

$ kubectl taint node k8s-master node-role.kubernetes.io/control-plane-
node/k8s-master untainted
```

## 下一步

安装 CNI 网络插件