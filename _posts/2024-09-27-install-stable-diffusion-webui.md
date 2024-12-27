---
layout: post
title: "安装 Stable Diffusion web UI"
date: 2024-09-27
---

## Install

```bash
git clone https://github.com/AUTOMATIC1111/stable-diffusion-webui
cd stable-diffusion-webui
python -m venv venv
source venv/bin/activate​
./webui.sh --xformers
```

<http://localhost:7860/>


## 安装中文扩展

打开 Extensions >> Available

取消勾选 localization 复选框

点击 Load from 按钮加载可用的插件

在页面中搜索找到 zh_Hans Localization 点击 Install

安装好后在 Settings >> User Interface >> Localization 选择 zh-Hans(Stable) 点击 Apply Settings ，Reload UI


## 安装 ControlNet

<https://github.com/Mikubill/sd-webui-controlnet.git>


打开 webui：http://127.0.0.1:7860/

找到 扩展 >> 从网址安装 TAB >> 扩展的 git 仓库网址

输入：https://github.com/Mikubill/sd-webui-controlnet.git

安装好后页面会出现提示，选择 已安装 TAB 重启 UI

下载模型

<https://huggingface.co/lllyasviel/ControlNet-v1-1/tree/main>

<https://hf-mirror.com/>

```bash
export HF_ENDPOINT=https://hf-mirror.com
huggingface-cli download --local-dir-use-symlinks False --resume-download lllyasviel/ControlNet-v1-1 --local-dir ControlNet-v1-1
```

将下载好的模型放入

stable-diffusion-webui/extensions/sd-webui-controlnet/models
