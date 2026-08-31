---
layout: post
title: "macos 通过 launched 自启动服务"
date: 2026-07-24
tags: other
---

macOS 用 launchd 做服务的开机自启动，相当于 Linux 的 systemd。

## 创建 LaunchDaemon

sudo vi /Library/LaunchDaemons/com.myservice.client.plist

```xml
<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN"
"http://www.apple.com/DTDs/PropertyList-1.0.dtd">

<plist version="1.0">
<dict>

    <key>Label</key>
    <string>com.myservice</string>

    <key>ProgramArguments</key>
    <array>
        <string>/usr/local/bin/myservice</string>
        <string>-c</string>
        <string>/Users/你的用户名/myservice/config.toml</string>
    </array>

    <key>RunAtLoad</key>
    <true/>

    <key>KeepAlive</key>
    <true/>

    <key>StandardOutPath</key>
    <string>/tmp/myservice.log</string>

    <key>StandardErrorPath</key>
    <string>/tmp/myservice.err</string>

</dict>
</plist>
```

## 设置权限

```bash
sudo chown root:wheel /Library/LaunchDaemons/com.myservice.plist
sudo chmod 644 /Library/LaunchDaemons/com.myservice.plist
```

## 加载服务

```bash
sudo launchctl bootstrap system /Library/LaunchDaemons/com.myservice.plist
```

## 查看服务

```bash
sudo launchctl list | grep myservice
sudo launchctl print system/com.myservice
```

## 测试停止/启动/重启

```bash
# 停止：
sudo launchctl bootout system /Library/LaunchDaemons/com.myservice.plist
# 启动：
sudo launchctl bootstrap system /Library/LaunchDaemons/com.myservice.plist
# 重启
sudo launchctl kickstart -k system/com.myservice
```