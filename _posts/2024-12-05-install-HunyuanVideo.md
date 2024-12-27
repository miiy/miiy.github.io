
---
layout: post
title: "安装 HunyuanVideo"
date: 2024-12-05
---

<https://github.com/Tencent/HunyuanV>

<https://github.com/Tencent/HunyuanVideo/blob/main/ckpts/README.md>

```bash
# 克隆仓库
git clone https://github.com/tencent/HunyuanVideo
cd HunyuanVideo
# 创建虚拟环境
python3 -m venv venv
source venv/bin/activate
# 安装依赖
pip install -r requirements.txt
# 安装 flash attention
# https://github.com/Dao-AILab/flash-attention/releases
# pip install flash_attn-2.5.9.post1+cu122torch2.1cxx11abiTRUE-cp311-cp311-linux_x86_64.whl
python -m pip install ninja
python -m pip install git+https://github.com/Dao-AILab/flash-attention.git@v2.5.9.post1
# 下载 HunyuanVideo 模型
pip install "huggingface_hub[cli]"
export HF_ENDPOINT=https://hf-mirror.com
huggingface-cli download --local-dir-use-symlinks False --resume-download tencent/HunyuanVideo --local-dir ./ckpts
# 下载 Text Encoder
huggingface-cli download --local-dir-use-symlinks False --resume-download xtuner/llava-llama-3-8b-v1_1-transformers --local-dir ./ckpts/llava-llama-3-8b-v1_1-transformers
# 将 llava-llama-3-8b-v1_1-transformers 的语言模型部分分离为 text_encoder
python hyvideo/utils/preprocess_text_encoder_tokenizer_utils.py --input_dir ckpts/llava-llama-3-8b-v1_1-transformers --output_dir ckpts/text_encoder
# 下载 CLIP 模型
huggingface-cli download --local-dir-use-symlinks False --resume-download openai/clip-vit-large-patch14 --local-dir ./ckpts/text_encoder_2
```

test

```bash
python3 sample_video.py \
    --video-size 720 1280 \
    --video-length 129 \
    --infer-steps 50 \
    --prompt "A cat walks on the grass, realistic style." \
    --flow-reverse \
    --use-cpu-offload \
    --save-path ./results
```
