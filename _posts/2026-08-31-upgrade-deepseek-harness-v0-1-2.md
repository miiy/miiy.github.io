---
layout: post
title: "升级 deepseek-harness 到 dsh-v0.1.2-alpha.2"
date: 2026-08-31
tags: AI
---

## 命令

```bash
git pull
rm -rf node_modules
pnpm run clean
pnpm run build
pnpm install
# 插件不兼容，更新插件
pnpm dsh plugin --profile web list
pnpm dsh plugin --profile web add dshmarket@1.38.1
pnpm dsh plugin --profile web add dsh-better-sidebar@v0.18.0-alpha.0
pnpm dsh --profile web
```