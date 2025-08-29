---
layout: post
title: "使用 kubeadm 安装 kubernetes 单节点学习环境"
---

版本：v1.33.0

参考安装文档：

<https://kubernetes.io/zh-cn/docs/setup/production-environment/tools/kubeadm/install-kubeadm/>

## 基础设置

```bash
# 设置 hostname
sudo hostnamectl set-hostname "k8s-master"
# 修改 hosts
echo -e "\n127.0.0.1	k8s-master" |sudo tee -a /etc/hosts
# 禁用 swap
sudo swapoff -a
sudo sed -ri 's/.*swap.*/#&/' /etc/fstab
```

## 安装容器运行时

### 使用 apt 安装 containerd

<https://docs.docker.com/engine/install/debian/>

使用阿里云镜像安装

<https://developer.aliyun.com/mirror/docker-ce>

```bash
# Add Docker's official GPG key:
sudo apt-get update
sudo apt-get install ca-certificates curl
sudo curl -fsSL https://mirrors.aliyun.com/docker-ce/linux/debian/gpg -o /etc/apt/keyrings/docker.gpg
sudo chmod a+r /etc/apt/keyrings/docker.gpg

echo \
  "deb [arch=$(dpkg --print-architecture) signed-by=/etc/apt/keyrings/docker.gpg] https://mirrors.aliyun.com/docker-ce/linux/debian \
  "$(. /etc/os-release && echo "$VERSION_CODENAME")" stable" | \
  sudo tee /etc/apt/sources.list.d/docker.list > /dev/null
sudo apt-get update
sudo apt-get install containerd.io
```

### 配置 containerd

<https://kubernetes.io/zh-cn/docs/setup/production-environment/container-runtimes/#containerd-systemd>

生成 conatinerd 默认配置文件并修改

```bash
$ sudo cp /etc/containerd/config.toml /etc/containerd/config.toml.old
# $ containerd config default > /etc/containerd/config.toml
$ containerd config default|sudo tee /etc/containerd/config.toml
sudo vi /etc/containerd/config.toml
```

配置 systemd cgroup 驱动

修改 pause 镜像地址

设置镜像源

```conf
[plugins."io.containerd.grpc.v1.cri"]
    sandbox_image = "registry.aliyuncs.com/google_containers/pause:3.8"

[plugins."io.containerd.grpc.v1.cri".containerd.runtimes.runc.options]
    SystemdCgroup = true

[plugins."io.containerd.grpc.v1.cri".registry]
   config_path = "/etc/containerd/certs.d"
```

```bash
# docker.io
sudo mkdir -p /etc/containerd/certs.d/docker.io

cat <<EOF | sudo tee /etc/containerd/certs.d/docker.io/hosts.toml
server = "https://docker.io"

[host."https://docker.m.daocloud.io"]
  capabilities = ["pull", "resolve"]
EOF

# quay.io
sudo mkdir -p /etc/containerd/certs.d/quay.io

cat <<EOF | sudo tee /etc/containerd/certs.d/quay.io/hosts.toml
server = "https://docker.io"

[host."https://docker.m.daocloud.io"]
  capabilities = ["pull", "resolve"]
EOF
```

```bash
sudo systemctl restart containerd
```

查看配置是否生效

```bash
containerd config dump
```

> 说明: 配置的是插件`io.containerd.grpc.v1.cri`的镜像地址，使用crictl拉取镜像文件才有效：`sudo crictl --debug pull docker.io/library/nginx:1.27`

## 启用 IPv4 数据包转发 

<https://kubernetes.io/zh-cn/docs/setup/production-environment/container-runtimes/>

手动启用 IPv4 数据包转发：

```bash
# 设置所需的 sysctl 参数，参数在重新启动后保持不变
cat <<EOF | sudo tee /etc/sysctl.d/k8s.conf
net.ipv4.ip_forward = 1
EOF

# 应用 sysctl 参数而不重新启动
sudo sysctl --system
```

使用以下命令验证 net.ipv4.ip_forward 是否设置为 1：

```bash
sudo sysctl net.ipv4.ip_forward
```

## 安装 kubeadm、kubelet 和 kubectl

<https://developer.aliyun.com/mirror/kubernetes/>

```bash
sudo apt-get update
sudo apt-get install -y apt-transport-https ca-certificates curl gpg

# curl -fsSL https://pkgs.k8s.io/core:/stable:/v1.33/deb/Release.key | sudo gpg --dearmor -o /etc/apt/keyrings/kubernetes-apt-keyring.gpg
curl -fsSL https://mirrors.aliyun.com/kubernetes-new/core/stable/v1.33/deb/Release.key | sudo gpg --dearmor -o /etc/apt/keyrings/kubernetes-apt-keyring.gpg

# echo 'deb [signed-by=/etc/apt/keyrings/kubernetes-apt-keyring.gpg] https://pkgs.k8s.io/core:/stable:/v1.33/deb/ /' | sudo tee /etc/apt/sources.list.d/kubernetes.list
echo 'deb [signed-by=/etc/apt/keyrings/kubernetes-apt-keyring.gpg] https://mirrors.aliyun.com/kubernetes-new/core/stable/v1.33/deb/ /' | sudo tee /etc/apt/sources.list.d/kubernetes.list

sudo apt-get update
sudo apt-get install -y kubelet kubeadm kubectl
sudo apt-mark hold kubelet kubeadm kubectl
```

## 使用 kubeadm 创建集群

```bash
sudo kubeadm init \
--control-plane-endpoint=k8s-master \
--image-repository registry.aliyuncs.com/google_containers \
--kubernetes-version v1.33.0 \
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

如果执行出现错误，运行 sudo kubeadm reset -f 后，再次执行

```bash
$ kubeadm init --control-plane-endpoint=k8s-master --image-repository registry.aliyuncs.com/google_containers --kubernetes-version v1.33.0 --pod-network-cidr=192.168.0.0/16 --v=9

...

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

You can now join any number of control-plane nodes by copying certificate authorities
and service account keys on each node and then running the following as root:

  kubeadm join k8s-master:6443 --token oseqli.fe4ucqn8tfk1xcmo \
	--discovery-token-ca-cert-hash sha256:062c69c85341d715afa33c45b30becab722e10e4c70de822771672d5ab303c7f \
	--control-plane

Then you can join any number of worker nodes by running the following on each as root:

kubeadm join k8s-master:6443 --token oseqli.fe4ucqn8tfk1xcmo \
	--discovery-token-ca-cert-hash sha256:062c69c85341d715afa33c45b30becab722e10e4c70de822771672d5ab303c7f
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

## 安装 CNI 网络插件

<https://miiy.github.io/kubernetes/02-install/install-cni-calico>

## 移除 master 污点

<https://kubernetes.io/zh-cn/docs/concepts/scheduling-eviction/taint-and-toleration/>

由于装的是单节点的学习环境，需要移除 master 的污点，pod 才可以调度到 master 节点

```bash
$ kubectl describe node|grep Taints
Taints:             node-role.kubernetes.io/control-plane:NoSchedule

$ kubectl taint nodes k8s-master node-role.kubernetes.io/control-plane:NoSchedule-
node/k8s-master untainted
```
