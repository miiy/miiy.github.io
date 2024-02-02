---
layout: post
title: "VideoReTalking 安装"
date: 2024-01-18
tags: AI
---


## Install

```bash
conda create -n video_retalking python=3.11
conda activate video_retalking
pip install torch torchvision torchaudio

pip install -i https://mirrors.aliyun.com/pypi/simple/ tb-nightly
sed -i 's/dlib==19.24.0/dlib==19.24.2/g' requirements.txt
pip install -r requirements.txt

export HF_ENDPOINT=https://hf-mirror.com
huggingface-cli download camenduru/video-retalking
ln -s /data/.cache/huggingface/hub/models--camenduru--video-retalking/snapshots/c9234f7d2ad684f340ee8e4a1ab18aaba9ccf463/ checkpoints
ls checkpoints
```

## run

```bash
$ python3 inference.py --face examples/face/1.mp4 --audio examples/audio/1.wav --outfile results/1_1.mp4
$ python3 inference.py --face examples/face/1_3s.mp4 --audio examples/audio/nvp_3s.wav --outfile results/1_3s.mp4
/path/to/miniconda3/envs/video_retalking/lib/python3.11/site-packages/torchvision/transforms/functional_tensor.py:5: UserWarning: The torchvision.transforms.functional_tensor module is deprecated in 0.15 and will be **removed in 0.17**. Please don't rely on it. You probably just need to use APIs in torchvision.transforms.functional or in torchvision.transforms.v2.functional.
  warnings.warn(
[Info] Using cuda for inference.
[Step 0] Number of frames available for inference: 66
[Step 1] Using saved landmarks.
[Step 2] 3DMM Extraction In Video:: 100%|██████████████████████████████████████████████████████████████████████████████████████████████████████████| 66/66 [00:01<00:00, 53.62it/s]
using expression center
Load checkpoint from: checkpoints/DNet.pt
Load checkpoint from: checkpoints/LNet.pth
Load checkpoint from: checkpoints/ENet.pth
[Step 3] Using saved stabilized video.
[Step 4] Load audio; Length of mel chunks: 64
[Step 5] Reference Enhancement: 100%|██████████████████████████████████████████████████████████████████████████████████████████████████████████████| 64/64 [00:08<00:00,  7.14it/s]
landmark Det:: 100%|███████████████████████████████████████████████████████████████████████████████████████████████████████████████████████████████| 64/64 [00:08<00:00,  7.37it/s]
100%|███████████████████████████████████████████████████████████████████████████████████████████████████████████████████████████████████████████| 64/64 [00:00<00:00, 24299.40it/s]
100%|█████████████████████████████████████████████████████████████████████████████████████████████████████████████████████████████████████████████| 64/64 [00:00<00:00, 522.02it/s]
FaceDet:: 100%|████████████████████████████████████████████████████████████████████████████████████████████████████████████████████████████████████| 16/16 [00:04<00:00,  3.31it/s]
start 1705488432.6658988██████████████████████████████████████████████████████████████████████████████████████████████████████████████████▊        | 15/16 [00:04<00:00,  7.62it/s]
model: 9.195826053619385
end 27.434037685394287
[Step 6] Lip Synthesis:: 100%|███████████████████████████████████████████████████████████████████████████████████████████████████████████████████████| 1/1 [00:43<00:00, 43.01s/it]
outfile: results/1_3s.mp4
```
