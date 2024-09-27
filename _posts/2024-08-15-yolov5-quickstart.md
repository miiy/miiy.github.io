---
layout: post
title: "yolov5 快速入门"
date: 2024-08-15
---

<https://github.com/ultralytics/yolov5>

<https://github.com/ultralytics/yolov5/blob/master/tutorial.ipynb>

<https://docs.ultralytics.com/yolov5/quickstart_tutorial/>

## Install

```
git clone https://github.com/ultralytics/yolov5.git
cd yolov5
conda create -n yolov5 python==3.11
conda activate yolov5
pip install -r requirements.txt
```

check

```python
import torch
import utils
display = utils.notebook_init()  # checks

YOLOv5 🚀 v7.0-353-g5eca7b9c Python-3.11.0 torch-2.4.0+cu121 CUDA:0 (NVIDIA GeForce RTX 4090 D, 24210MiB)
Setup complete ✅ (24 CPUs, 62.5 GB RAM, 895.1/915.3 GB disk)
```

## Detect

detect.py 在各种源上运行YOLOv5 推理，从最新的YOLOv5版本自动下载模型，并将结果保存到 runs/detect 。

```bash
python detect.py --weights yolov5s.pt --img 640 --conf 0.25 --source data/images
```