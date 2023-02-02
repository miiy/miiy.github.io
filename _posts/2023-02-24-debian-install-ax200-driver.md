---
layout: post
title: "debian 安装 Ax200 驱动"
date: 2023-02-24
tags: linux
---

安装 Ax200 驱动

```bash
sudo apt install firmware-iwlwifi bluez
sudo reboot
```

如果蓝牙有问题参考下面的错误排查方法：

```bash
bluetoothctl show
# 如果显示 No default controller available 则表示没有驱动起来
# 查看蓝牙情况
dmesg | grep blue
# 我这里显示
# bluetooth hci0: Direct firmware load for intel/ibt-20-1-3.sfi failed with error -2
# 表示缺少 ibt-20-1-3.sfi 驱动文件

# 进入驱动文件夹
cd/lib/firmware/intel
# 下载缺失的驱动，如果你缺少的是别的驱动，则替换掉文件名即可
wget https://git.kernel.org/pub/scm/linux/kernel/git/firmware/linux-firmware.git/plain/intel/ibt-20-1-3.sfi
# 下载完毕后重启设备
reboot now
# 再次查看蓝牙设备
bluetoothctl show
```

