---
layout: post
title: "安装 flash-attention"
date: 2025-03-11
---

首先确保 pytorch 已经安装好

查看版本号选择 flash-attn 对应的版本

<https://github.com/kvcache-ai/ktransformers/issues/410>

```bash
$ python -c "import torch; print(torch.cuda.is_available()); print(torch.__version__); print(torch.version.cuda)"
True
2.5.1+cu121
12.1
$ python -c "import triton; print(triton.__version__)"
3.1.0
$ python -c "import torch; print(torch._C._GLIBCXX_USE_CXX11_ABI)"
False
```

下载对应版本的 whl 包，安装

<https://github.com/Dao-AILab/flash-attention/releases>

```bash
python -c "import flash_attn; print('FlashAttention is working\!')"
```
