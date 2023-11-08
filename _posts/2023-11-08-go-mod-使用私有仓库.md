---
layout: post
title: "go mod 使用私有仓库"
date: 2023-11-08
tags: go
---

使用 go env 配置私有仓库地址

```bash
go env -w GOPRIVATE=gitlab.test.com
```

如果私有仓库是 http 的，使用以下命令

```bash
go env -w GOINSECURE=gitlab.test.com
```

使用 git config 替换拉取地址

```bash
git config --global url."http://user:pass@gitlab.test.com/".insteadOf "http://gitlab.test.com/"
```
