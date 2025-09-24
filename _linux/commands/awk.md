---
layout: post
title: "awk"
---

```bash
$ pip list | awk '{print $1}'
Package
-----------------
lxml
pillow
pip
setuptools
wheel
$ pip list | awk 'NR > 3 {print $1}'
lxml
pillow
pip
setuptools
wheel
```

