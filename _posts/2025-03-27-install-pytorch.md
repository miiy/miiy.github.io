---
layout: post
title: "安装 pytorch"
date: 2025-03-27
---

## Install nvidia driver

```bash
$ sudo apt install nvidia-driver-550
$ sudo reboot
$ nvidia-smi
Wed Mar 26 10:18:27 2025
+-----------------------------------------------------------------------------------------+
| NVIDIA-SMI 550.144.03             Driver Version: 550.144.03     CUDA Version: 12.4     |
|-----------------------------------------+------------------------+----------------------+
| GPU  Name                 Persistence-M | Bus-Id          Disp.A | Volatile Uncorr. ECC |
| Fan  Temp   Perf          Pwr:Usage/Cap |           Memory-Usage | GPU-Util  Compute M. |
|                                         |                        |               MIG M. |
|=========================================+========================+======================|
|   0  NVIDIA H800                    Off |   00000000:18:00.0 Off |                    0 |
| N/A   27C    P0             81W /  700W |      14MiB /  81559MiB |      0%      Default |
|                                         |                        |             Disabled |
+-----------------------------------------+------------------------+----------------------+
|   1  NVIDIA H800                    Off |   00000000:38:00.0 Off |                    0 |
| N/A   27C    P0             79W /  700W |      14MiB /  81559MiB |      0%      Default |
|                                         |                        |             Disabled |
+-----------------------------------------+------------------------+----------------------+
|   2  NVIDIA H800                    Off |   00000000:49:00.0 Off |                    0 |
| N/A   27C    P0             83W /  700W |      14MiB /  81559MiB |      0%      Default |
|                                         |                        |             Disabled |
+-----------------------------------------+------------------------+----------------------+
|   3  NVIDIA H800                    Off |   00000000:59:00.0 Off |                    0 |
| N/A   28C    P0             81W /  700W |      14MiB /  81559MiB |      0%      Default |
|                                         |                        |             Disabled |
+-----------------------------------------+------------------------+----------------------+
|   4  NVIDIA H800                    Off |   00000000:9B:00.0 Off |                    0 |
| N/A   26C    P0             81W /  700W |      14MiB /  81559MiB |      0%      Default |
|                                         |                        |             Disabled |
+-----------------------------------------+------------------------+----------------------+
|   5  NVIDIA H800                    Off |   00000000:BB:00.0 Off |                    0 |
| N/A   27C    P0             83W /  700W |      14MiB /  81559MiB |      0%      Default |
|                                         |                        |             Disabled |
+-----------------------------------------+------------------------+----------------------+
|   6  NVIDIA H800                    Off |   00000000:CA:00.0 Off |                    0 |
| N/A   26C    P0             81W /  700W |      14MiB /  81559MiB |      0%      Default |
|                                         |                        |             Disabled |
+-----------------------------------------+------------------------+----------------------+
|   7  NVIDIA H800                    Off |   00000000:DA:00.0 Off |                    0 |
| N/A   28C    P0             81W /  700W |      14MiB /  81559MiB |      0%      Default |
|                                         |                        |             Disabled |
+-----------------------------------------+------------------------+----------------------+

+-----------------------------------------------------------------------------------------+
| Processes:                                                                              |
|  GPU   GI   CI        PID   Type   Process name                              GPU Memory |
|        ID   ID                                                               Usage      |
|=========================================================================================|
|    0   N/A  N/A      5907      G   /usr/lib/xorg/Xorg                              4MiB |
|    1   N/A  N/A      5907      G   /usr/lib/xorg/Xorg                              4MiB |
|    2   N/A  N/A      5907      G   /usr/lib/xorg/Xorg                              4MiB |
|    3   N/A  N/A      5907      G   /usr/lib/xorg/Xorg                              4MiB |
|    4   N/A  N/A      5907      G   /usr/lib/xorg/Xorg                              4MiB |
|    5   N/A  N/A      5907      G   /usr/lib/xorg/Xorg                              4MiB |
|    6   N/A  N/A      5907      G   /usr/lib/xorg/Xorg                              4MiB |
|    7   N/A  N/A      5907      G   /usr/lib/xorg/Xorg                              4MiB |
+-----------------------------------------------------------------------------------------+
```

## Install cuda toolkit

```bash
$ sudo apt install cuda-toolkit-12-4
$ nvcc --version
nvcc: NVIDIA (R) Cuda compiler driver
Copyright (c) 2005-2024 NVIDIA Corporation
Built on Thu_Mar_28_02:18:24_PDT_2024
Cuda compilation tools, release 12.4, V12.4.131
Build cuda_12.4.r12.4/compiler.34097967_0
```

## Install pytorch

<https://pytorch.org/>

```bash
pip3 install torch torchvision torchaudio
```

```bash
$ python -c "import torch; print(torch.version.cuda)"

$ python -c "import torch; print(torch.cuda.is_available())"
True
$ python -c "import torch; print(torch.cuda.device_count())"

```