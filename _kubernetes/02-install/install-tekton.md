## 安装 tekton

修改 conatinerd 配置文件

```conf
[plugins."io.containerd.grpc.v1.cri".cri.registry.mirrors."registry.cn-hangzhou.aliyuncs.com"]
    endpoint = ["https://registry.cn-hangzhou.aliyuncs.com"]

    [plugins."io.containerd.grpc.v1.cri".registry.auths."registry.cn-hangzhou.aliyuncs.com"]
        username = "username"
        password = "password"
```

重启 containerd

```bash
sudo systemctl restart containerd
```

通过阿里云容器镜像服务构建镜像，修改yaml镜像信息为阿里云，然后安装

https://github.com/miiy/aliyun-registry

## 安装 Tekton CLI

```bash
# Get the tar.xz
curl -LO https://github.com/tektoncd/cli/releases/download/v0.29.0/tkn_0.29.0_Linux_x86_64.tar.gz
# Extract tkn to your PATH (e.g. /usr/local/bin)
sudo tar xvzf tkn_0.29.0_Linux_x86_64.tar.gz -C /usr/local/bin/ tkn
```

## Access Tekton Dashboard

```bash
kubectl port-forward --address 0.0.0.0  -n tekton-pipelines service/tekton-dashboard 9097:9097
```
