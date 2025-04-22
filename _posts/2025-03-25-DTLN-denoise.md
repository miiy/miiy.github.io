---
layout: post
title: "DTLN 降噪"
date: 2025-03-25
---

## Install

```bash
git clone https://github.com/breizhn/DTLN.git
conda create -n dtln python==3.10
conda activate dtln
pip install soundfile librosa tensorflow wavinfo
```

## Inference

```bash
python run_evaluation.py -i ./data/src/ -o ./data/target/ -m ./pretrained_model/model.h5
```