---
layout: post
title: "安装 GLM-4"
date: 2024-10-15
---

## Install

GLM-4: <https://github.com/THUDM/GLM-4>

```bash
git clone https://github.com/THUDM/GLM-4.git
cd GLM-4/basic_demo/
conda create -n GLM-4 python==3.11
onda activate GLM-4
pip install -r requirements.txt
```

download checkpoints

```bash
pip install modelscope
modelscope download --model ZhipuAI/glm-4-9b-chat --local_dir /data/GLM-4/model/glm-4-9b-chat
```

test

```bash
MODEL_PATH=/data/GLM-4/model/glm-4-9b-chat/ python trans_cli_demo.py
```

```bash
$ nvidia-smi
Tue Oct 15 16:19:41 2024
+---------------------------------------------------------------------------------------+
| NVIDIA-SMI 535.129.03             Driver Version: 535.129.03   CUDA Version: 12.2     |
|-----------------------------------------+----------------------+----------------------+
| GPU  Name                 Persistence-M | Bus-Id        Disp.A | Volatile Uncorr. ECC |
| Fan  Temp   Perf          Pwr:Usage/Cap |         Memory-Usage | GPU-Util  Compute M. |
|                                         |                      |               MIG M. |
|=========================================+======================+======================|
|   0  NVIDIA A100-PCIE-40GB          On  | 00000000:00:06.0 Off |                    0 |
| N/A   29C    P0              36W / 250W |  36882MiB / 40960MiB |      0%      Default |
|                                         |                      |             Disabled |
+-----------------------------------------+----------------------+----------------------+

+---------------------------------------------------------------------------------------+
| Processes:                                                                            |
|  GPU   GI   CI        PID   Type   Process name                            GPU Memory |
|        ID   ID                                                             Usage      |
|=======================================================================================|
|    0   N/A  N/A     81572      C   ...ps/path/bin/python                      2332MiB |
|    0   N/A  N/A   2057658      C   python                                    34522MiB |
+---------------------------------------------------------------------------------------+
```

## Run openai api server

```bash
export MODEL_PATH=/data/GLM-4/model/glm-4-9b-chat/
python openai_api_server.py
pip install vllm
python openai_api_server.py
```
