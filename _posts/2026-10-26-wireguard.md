---
layout: post
title: "WireGuard"
date: 2026-10-26
tags: AI
---

一台公网服务器作为 WireGuard 服务端，Dell 服务器与三台个人设备作为客户端接入同一网段 `10.66.66.0/24`。全文分三部分：服务端、Linux 客户端、macOS 客户端。

## 网络规划

| 角色 | 密钥前缀 | 隧道地址 | 说明 |
| --- | --- | --- | --- |
| 服务端 | `server` | `10.66.66.1/24` | 唯一监听 UDP `51820` 的节点 |
| Dell 服务器 | `dell-server` | `10.66.66.2/32` | 公司内网，作为普通客户端接入 |
| Mac mini | `macmini` | `10.66.66.3/32` | |
| 手机 | `phone` | `10.66.66.11/32` | |
| MacBook Pro | `macbook` | `10.66.12/32` | |

## 一、服务端

### 1. 安装 WireGuard

在服务端 Debian 上执行：

```bash
sudo apt update
sudo apt install wireguard -y
```

### 2. 生成密钥

在服务端一次性生成本机与四台客户端的密钥对：

```bash
sudo install -d -m 700 /etc/wireguard

sudo bash -c '
umask 077
cd /etc/wireguard

for name in server dell-server macmini macbook phone; do
    wg genkey > "${name}.key"
    wg pubkey < "${name}.key" > "${name}.pub"
done

chmod 600 *.key
chmod 644 *.pub
'
```

核对生成结果：

```bash
sudo ls -l /etc/wireguard/
```

```text
server.key       server.pub
dell-server.key  dell-server.pub
macmini.key      macmini.pub
macbook.key      macbook.pub
phone.key        phone.pub
```

> **不要重复执行本步。** 命令会覆盖同名密钥文件，已下发的客户端配置会立即失效。

> 正式使用时，建议在每台设备上自行生成私钥；服务端只保留 `server.key` 和四个 `.pub` 公钥文件。

### 3. 编写 wg0.conf

取出服务端私钥和四个客户端公钥：

```bash
sudo cat /etc/wireguard/server.key

sudo cat /etc/wireguard/dell-server.pub
sudo cat /etc/wireguard/macmini.pub
sudo cat /etc/wireguard/macbook.pub
sudo cat /etc/wireguard/phone.pub
```

`server.key` 只写入下面的配置，不对外发送。

创建配置文件：

```bash
sudo vi /etc/wireguard/wg0.conf
```

写入以下内容，并把五处占位符替换为实际密钥：

```ini
[Interface]
Address = 10.66.66.1/24
ListenPort = 51820
PrivateKey = SERVER_PRIVATE_KEY

# Dell 服务器
[Peer]
PublicKey = DELL_SERVER_PUBLIC_KEY
AllowedIPs = 10.66.66.2/32

# Mac mini
[Peer]
PublicKey = MACMINI_PUBLIC_KEY
AllowedIPs = 10.66.66.3/32

# 手机
[Peer]
PublicKey = PHONE_PUBLIC_KEY
AllowedIPs = 10.66.66.11/32

# MacBook Pro
[Peer]
PublicKey = MACBOOK_PUBLIC_KEY
AllowedIPs = 10.66.66.12/32
```

保存退出后收紧权限：

```bash
sudo chmod 600 /etc/wireguard/wg0.conf
```

### 4. 开放 UDP 端口

在云平台控制台的安全组中添加入方向规则：协议选 UDP、端口 `51820`、授权对象按实际安全需求设置，不要误设成 TCP。

### 5. 启动并验证

```bash
sudo systemctl enable --now wg-quick@wg0
```

```bash
sudo systemctl status wg-quick@wg0 --no-pager
sudo wg show
ip addr show wg0
```

正常情况下应看到接口 `wg0`，地址为 `10.66.66.1/24`。启动失败时查看日志：

```bash
sudo journalctl -u wg-quick@wg0 -n 50 --no-pager
```

### WireGuard 两种访问场景

1. 访问 VPN 服务器本身 例如：手机 `10.66.66.6` 访问 ECS `10.66.66.1`。

- 流量经过防火墙 `INPUT` 链。
- 通常不需要开启 IP 转发，也不需要配置 `FORWARD` 规则。

2. 通过 VPN 服务器访问另一台客户端 例如：手机 `10.66.6` 访问 Mac `10.66.66.3`，由 ECS 转发流量。

开启 IP 转发：

```bash
sudo sysctl -w net.ipv4.ip_forward=1
```

允许 VPN 客户端之间转发：

```bash
sudo iptables -I FORWARD 1 \
  -i wg0 -o wg0 \
  -s 10.66.66.0/24 -d 10.66.66.0/24 \
  -j ACCEPT
```

注意：以上 IP 转发设置和 iptables 规则默认不会在重启后自动保留，需要另行配置持久化。


## 二、客户端（Linux，以 Dell 服务器为例）

四台客户端配置方式相同，下面是 Dell 服务器的完整示例：安装 → 配置 → 启动。密钥沿用服务端第 2 步生成的 `dell-server.key`。

### 1. 安装 WireGuard

```bash
sudo apt update
sudo apt install wireguard -y
```

Debian / Ubuntu 用上面两条；其他发行版换对应包管理器安装 `wireguard-tools`。

### 2. 编写 wg0.conf

配置里的私钥填服务端第 2 步生成的 `dell-server.key`：

```bash
sudo cat /etc/wireguard/dell-server.key
```

```bash
sudo vi /etc/wireguard/wg0.conf
```

```ini
[Interface]
Address = 10.66.66.2/24
PrivateKey = DELL_SERVER_PRIVATE_KEY

[Peer]
PublicKey = SERVER_PUBLIC_KEY
Endpoint = SERVER_PUBLIC_IP:51820
AllowedIPs = 10.66.66.0/24
PersistentKeepalive = 25
```

字段说明：

| 字段 | 取值 | 说明 |
| --- | --- | --- |
| `Address` | `10.66.66.2/24` | 网络规划表里的 `/32` 是服务端 `AllowedIPs` 的写法，客户端自己的地址用 `/24`，`wg-quick` 据此生成隧道网段路由 |
| `PrivateKey` | `dell-server.key` 的内容，即服务端第 2 步生成的私钥 | 不外传 |
| `PublicKey` | `server.pub` 的内容，即服务端第 2 步生成的服务端公钥 | 对端填服务端公钥 |
| `Endpoint` | 服务端的公网 IP + `:51820` | 客户端连往的入口地址 |
| `AllowedIPs` | `10.66.66.0/24` | 只有隧道网段走 VPN；改成 `0.0.0.0/0` 则整机流量走隧道，需同时加 `DNS =` 一行 |
| `PersistentKeepalive` | `25` | 处于 NAT 后时每 25 秒保活，防止隧道空闲失效 |

其余三台设备照抄本节配置，只替换两处：

- `Address`：分别改为 `10.66.66.3/24`、`10.66.66.11/24`、`10.66.66.12/24`；
- `PrivateKey`：各自设备的私钥。

Mac 与手机用官方 WireGuard App 导入同样的字段即可，无需命令行。

### 3. 启动并验证

```bash
sudo chmod 600 /etc/wireguard/wg0.conf
sudo systemctl enable --now wg-quick@wg0
sudo wg show
ping -c 3 10.66.66.1
```

`wg show` 的 `latest handshake` 有值、能 ping 通 `10.66.66.1`，即表示客户端已接入。

## 三、客户端（macOS）

能直接安装官方 WireGuard App 时优先用 App。无法下载 App 时只能手动搭建：macOS 没有内核态的 WireGuard 模块，隧道由 `wireguard-go` 创建 `utun` 接口，因此需要自己安装 `wireguard-go` 与 `wireguard-tools`。

镜像目录：

- `wireguard-go`：<https://mirrors.nju.edu.cn/macports/packages/wireguard-go/>
- `wireguard-tools`：<https://mirrors.nju.edu.cn/macports/packages/wireguard-tools/>

### 1. 安装 Bash 5.3

`wg-quick` 要求 Bash 4.0 以上，macOS 自带的是 3.2，需要单独编一份装在 `/usr/local/bash-5.3`，不改动系统自带的 Bash：

```bash
cd ~/Downloads
curl -LO https://mirror.nju.edu.cn/gnu/bash/bash-5.3.tar.gz
tar -xzf bash-5.3.tar.gz
cd bash-5.3
./configure --prefix="/usr/local/bash-5.3"
make -j4
sudo make install
```

源码也可以从 <https://ftp.gnu.org/gnu/bash/> 取，与上面的镜像地址等价。编译依赖 Xcode 命令行工具，缺了先执行 `xcode-select --install`。

### 2. 解压扩展包并固定路径

从上面的镜像目录取与系统匹配的包（示例为 Apple Silicon：`darwin_27.arm64`），解压到 `/usr/local`，再把各自的 `opt/local` 链接成固定名字，最终目录结构：

```text
/usr/local/bash-5.3
/usr/local/wireguard-tools -> wireguard-tools-1.0.20260223_0.darwin_27.arm64/opt/local
/usr/local/wireguard-go -> wireguard-go-0.0.20250522_0.darwin_27.arm64/opt/local
```

固定成这两个路径后，下面所有命令与 `~/.zshrc` 都不用再跟着版本号改。

### 3. 配置 PATH

在 `~/.zshrc` 中添加：

```bash
export PATH="/usr/local/wireguard-tools/bin:/usr/local/wireguard-go/bin:$PATH"
```

生效：

```bash
source ~/.zshrc
```

### 4. 配置文件

客户端配置路径：`/etc/wireguard/wg0.conf`

```bash
sudo vi /etc/wireguard/wg0.conf
```

示例（Mac mini `10.66.66.3`，MacBook Pro 改成 `.4`，其余照抄）：

```ini
[Interface]
PrivateKey = <Mac 客户端私钥>
Address = 10.66.66.3/24

[Peer]
PublicKey = <服务端公钥>
Endpoint = <服务端公网 IP>:51820
AllowedIPs = 10.66.66.0/24
PersistentKeepalive = 25
```

保护私钥：

```bash
sudo chmod 600 /etc/wireguard/wg0.conf
```

`Address` 用 `/24`，与第二部分一致，`wg-quick` 据此生成 `10.66.66.0/24` 的隧道路由。

### 5. 常用命令

启动：

```bash
sudo /usr/local/bash-5.3/bin/bash /usr/local/wireguard-tools/bin/wg-quick up wg0
```

查看状态：

```bash
sudo /usr/local/wireguard-tools/bin/wg show
```

断开：

```bash
sudo /usr/local/bash-5.3/bin/bash /usr/local/wireguard-tools/bin/wg-quick down wg0
```
