---
layout: post
title: "Docker 支持 GPU"
---

## Install

```bash
sudo apt install nvidia-container-toolkit
```

<https://docs.nvidia.com/datacenter/cloud-native/container-toolkit/latest/index.html>

## Usage

```bash
sudo docker run --rm --runtime=nvidia --gpus all ubuntu nvidia-smi
```

<https://docs.nvidia.com/datacenter/cloud-native/container-toolkit/latest/sample-workload.html#running-a-sample-workload-with-docker>
