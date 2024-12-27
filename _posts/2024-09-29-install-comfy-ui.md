---
layout: post
title: "安装 ComfyUI"
date: 2024-09-29
---

## Install

```bash
git clone https://github.com/comfyanonymous/ComfyUI.git
cd ComfyUI
conda create -n comfyui python==3.11
conda activate comfyui
pip install torch torchvision torchaudio --extra-index-url https://download.pytorch.org/whl/cu124
pip install -r requirements.txt
```

## Download StableDiffusion checkpoint

```bash
cp v1-5-pruned-emaonly.safetensors ComfyUI/models/checkpoints/
```

## Run

```bash
python main.py
```

<http://127.0.0.1:8188/>


## 安装 ControlNet

<https://github.com/Fannovel16/comfyui_controlnet_aux>

```bash
cd ComfyUI/custom_nodes/
git clone https://github.com/Fannovel16/comfyui_controlnet_aux/
cd comfyui_controlnet_aux
pip install -r requirements.txt
cd ../../
python main.py
```