---
layout: post
title: "安装 JDK"
date: 2016-06-13
tags: linux
---

# Install

<https://docs.oracle.com/en/java/javase/15/install/installation-jdk-linux-platforms.html>

```bash
cd /usr/local/src
wget https://download.oracle.com/otn-pub/java/jdk/15.0.2+7/0d1cfde4252546c6931946de8db48ee2/jdk-15.0.2_linux-x64_bin.tar.gz
tar -zxvf jdk-15.0.2_linux-x64_bin.tar.gz 
mv jdk-15.0.2 ../

cat << EOF >> ~/.profile

export JAVA_HOME=/usr/local/jdk-15.0.2
export PATH=\$PATH:\$JAVA_HOME/bin
EOF

source ~/.profile
java -version
```
