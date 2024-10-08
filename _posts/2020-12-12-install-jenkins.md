---
layout: post
title: "安装 Jenkins"
date: 2020-12-12
tags: jenkins
---

<https://www.jenkins.io/zh/>

## 使用 docker 启动 jenkins

https://www.jenkins.io/doc/book/installing/docker/

```bash
sudo docker run -u root --name jenkins -d \
    --add-host gitlab.k8s:10.0.2.20 \
    --add-host harbor.example.com:10.0.2.20 \
    --restart always \
    -p 8082:8080 -p 50000:50000 \
    -v /data/jenkins/jenkins_home:/var/jenkins_home \
    -v /usr/bin/docker:/usr/bin/docker \
    -v /var/run/docker.sock:/var/run/docker.sock \
    jenkins/jenkins
```

## 安装 jenkins 插件

Docker plugin, Docker Pipeline, GitLab Plugin

## 流水线示例

新建任务选择流水线

构建触发器 

选择 Build when a change is pushed to GitLab. GitLab webhook URL: http://jenkins.example.com:8082/project/laravel

点击 高级 生成 Secret token

在 gitlab 项目中配置 webhook

Pipeline script

```text
env.BUILD_ID = 'latest'

node {
    stage('Git clone.') {
        // Get some code from a GitHub repository
        git branch: 'master', credentialsId: '05ddb37a-6218-4ab9-8565-4a0cd7c7217d', url: 'http://gitlab.k8s:8081/root/laravel.git'
        echo 'git colne finished.'
    }
        
    stage('Build Docker'){
        docker.withRegistry("https://harbor.example.com:4433", 'deb9142d-8e3d-48cb-9d58-b7bdebb17d41') {
            def customImage = docker.build("harbor.example.com:4433/library/laravel:${env.BUILD_ID}")
            /* Push the container to the custom Registry */
            customImage.push()
        }
    }
}
```

## gitlab 设置

打开 Admin Area >> Settings >> Network >> Outbound requests

勾选

```text
Allow requests to the local network from web hooks and services
```

允许 web hooks and service 请求本地网络


项目 >> Settings >> Webhooks

## Blue Ocean

安装插件 Blue Ocean，在首页点击进入
