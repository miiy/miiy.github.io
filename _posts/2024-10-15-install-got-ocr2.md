---
layout: post
title: "安装 GOT-OCR 2.0"
date: 2024-10-15
---

## Install

install dependencies

```bash
git clone https://github.com/Ucas-HaoranWei/GOT-OCR2.0.git
cd GOT-OCR2.0/GOT-OCR-2.0-master
conda create -n got-ocr python==3.11
conda activate got-ocr
pip install -e .
```

download checkpoints

```bash
pip install -U huggingface_hub
export HF_ENDPOINT=https://hf-mirror.com
huggingface-cli download --repo-type model --resume-download stepfun-ai/GOT-OCR2_0 --local-dir-use-symlinks False --local-dir stepfun-ai/GOT-OCR2_0
```

## Running

测试 demo 中的图 4

```bash
python3 GOT/demo/run_ocr_2.0.py  --model-name  stepfun-ai/GOT-OCR2_0/  --image-file  demo.jpg  --type ocr
<|im_start|>system
You should follow the instructions carefully and explain your answers in detail.<|im_end|><|im_start|>user
<img><imgpad>...<imgpad></img>
OCR: <|im_end|><|im_start|>assistant

20.如图，E是以菱形ABCD的边AD为直径的半圆弧上一点，∠BAD=60°，AB=BE=2DE=2，且M为AD的中 点. （1）求证：平面BEM⊥平面DEM； （2）设P为E上任意一点，求二面角B-PD-C的余弦值取值范围. 21.如图，点P在抛物线C：y=x²外，过P作抛物线C的两切线，设两切点分别为A(x1,x2)，B(x2,x2)，记线段AB 的中点为M. （1）求切线PA，PB的方程； （2）设点P为圆D：x²+y+2）=1上的点，当|AB|取最大值时，求点P的纵坐标. 22.已知函数f（x)=alnx+x（a≠0），g（x)=e²+bx²（b∈R） （1）记h（x)=f（x)+x²，试讨论函数h（x）的单调性； （2）若曲线y=f（x）与曲线y=g（x）在x=1处的切线都过点（0，1）.求证：当x>0时，f（x)+g（x）-1 x≥e-1
$ python3 GOT/demo/run_ocr_2.0.py  --model-name  stepfun-ai/GOT-OCR2_0/  --image-file  demo.jpg  --type format
<|im_start|>system
You should follow the instructions carefully and explain your answers in detail.<|im_end|><|im_start|>user
<img><imgpad>...<imgpad></img>
OCR with format: <|im_end|><|im_start|>assistant

20. 如图, \(E\) 是以菱形 \(A B C D\) 的边 \(A D\) 为直径的半圆弧上一点, \(\angle B A D=60^{\circ}, A B=B E=2 D E=2\), 且 \(M\) 为 \(A D\) 的中点.
(1) 求证: 平面 \(B E M \perp\) 平面 \(D E M\);
(2) 设 \(P\) 为 \(\sqrt{2} E\) 上任意一点, 求二面角 \(B-P D-C\) 的余弦值取值范围.

21. 如图, 点 \(P\) 在抛物线 \(C: y=x^{2}\) 外, 过 \(P\) 作抛物线 \(C\) 的两切线, 设两切点分别为 \(A\left(x_{1}, x_{1}^{2}\right), B\left(x_{2}, x_{2}^{2}\right)\), 记线段 \(A B\)的中点为 \(M\).
(1) 求切线 \(P A, P B\) 的方程;
(2) 设点 \(P\) 为圆 \(D: x^{2}+(y+2)^{2}=1\) 上的点, 当 \(\frac{|A B|}{|P M|}\) 取最大值时, 求点 \(P\) 的纵坐标.

22. 已知函数 \(f(x)=a \ln x+x(a \neq 0), g(x)=\mathrm{e}^{x}+b x^{2}(b \in \mathbf{R})\).
(1) 记 \(h(x)=f(x)+x^{2}\), 试讨论函数 \(h(x)\) 的单调性;
(2) 若曲线 \(y=f(x)\) 与曲线 \(y=g(x)\) 在 \(x=1\) 处的切线都过点 \((0,1)\). 求证: 当 \(x>0\) 时, \(f(x)+\frac{g(x)-1}{x} \geq \mathrm{e}-1\)
```

demo: <https://huggingface.co/spaces/stepfun-ai/GOT_official_online_demo>

stepfun-ai/GOT-OCR2_0: <https://huggingface.co/stepfun-ai/GOT-OCR2_0>

HF-Mirror: <https://hf-mirror.com/>