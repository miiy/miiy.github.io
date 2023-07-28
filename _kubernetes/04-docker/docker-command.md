---
layout: post
title: "Docker Command"
---

## Command

```bash
$ docker       

Usage:  docker [OPTIONS] COMMAND

A self-sufficient runtime for containers

Options:
      --config string      Location of client config files (default "/Users/u/.docker")
  -c, --context string     Name of the context to use to connect to the daemon (overrides DOCKER_HOST env var and default context set with "docker context use")
  -D, --debug              Enable debug mode
  -H, --host list          Daemon socket(s) to connect to
  -l, --log-level string   Set the logging level ("debug"|"info"|"warn"|"error"|"fatal") (default "info")
      --tls                Use TLS; implied by --tlsverify
      --tlscacert string   Trust certs signed only by this CA (default "/Users/u/.docker/ca.pem")
      --tlscert string     Path to TLS certificate file (default "/Users/u/.docker/cert.pem")
      --tlskey string      Path to TLS key file (default "/Users/u/.docker/key.pem")
      --tlsverify          Use TLS and verify the remote
  -v, --version            Print version information and quit

Management Commands:
  builder     Manage builds
  buildx*     Docker Buildx (Docker Inc., v0.8.2)
  compose*    Docker Compose (Docker Inc., v2.4.1)
  config      Manage Docker configs
  container   Manage containers
  context     Manage contexts
  image       Manage images
  manifest    Manage Docker image manifests and manifest lists
  network     Manage networks
  node        Manage Swarm nodes
  plugin      Manage plugins
  sbom*       View the packaged-based Software Bill Of Materials (SBOM) for an image (Anchore Inc., 0.6.0)
  scan*       Docker Scan (Docker Inc., v0.17.0)
  secret      Manage Docker secrets
  service     Manage services
  stack       Manage Docker stacks
  swarm       Manage Swarm
  system      Manage Docker
  trust       Manage trust on Docker images
  volume      Manage volumes

Commands:
  attach      Attach local standard input, output, and error streams to a running container
  build       Build an image from a Dockerfile
  commit      Create a new image from a container's changes
  cp          Copy files/folders between a container and the local filesystem
  create      Create a new container
  diff        Inspect changes to files or directories on a container's filesystem
  events      Get real time events from the server
  exec        Run a command in a running container
  export      Export a container's filesystem as a tar archive
  history     Show the history of an image
  images      List images
  import      Import the contents from a tarball to create a filesystem image
  info        Display system-wide information
  inspect     Return low-level information on Docker objects
  kill        Kill one or more running containers
  load        Load an image from a tar archive or STDIN
  login       Log in to a Docker registry
  logout      Log out from a Docker registry
  logs        Fetch the logs of a container
  pause       Pause all processes within one or more containers
  port        List port mappings or a specific mapping for the container
  ps          List containers
  pull        Pull an image or a repository from a registry
  push        Push an image or a repository to a registry
  rename      Rename a container
  restart     Restart one or more containers
  rm          Remove one or more containers
  rmi         Remove one or more images
  run         Run a command in a new container
  save        Save one or more images to a tar archive (streamed to STDOUT by default)
  search      Search the Docker Hub for images
  start       Start one or more stopped containers
  stats       Display a live stream of container(s) resource usage statistics
  stop        Stop one or more running containers
  tag         Create a tag TARGET_IMAGE that refers to SOURCE_IMAGE
  top         Display the running processes of a container
  unpause     Unpause all processes within one or more containers
  update      Update configuration of one or more containers
  version     Show the Docker version information
  wait        Block until one or more containers stop, then print their exit codes

Run 'docker COMMAND --help' for more information on a command.

To get more help with docker, check out our guides at https://docs.docker.com/go/guides/
```


```bash
# list images
docker images
# list all images （none镜像不要删除，镜像是一层一层叠加起来的，存在依赖关系）
docker images -a
# dangling image（虚悬镜像: 镜像没有仓库名或没有标签）
# list dangling image
docker images -f dangling=true

# list containers
docker ps
docker ps -a
docker ps -aq

docker rmi # 删除镜像
# delete all conatiners
docker rm $(docker ps -aq)
# delete all images
docker rmi $(docker images -aq)
# stop exited conttainers
docker stop $(docker ps -a | grep "Exited" | awk '{print $1 }') /
# delete exited conttainers
docker rm $(docker ps -a | grep "Exited" | awk '{print $1 }')
# delete <none> images
docker rmi $(docker images | grep "none" | awk '{print $3}')
# delete dangling images
docker rmi $(docker images -q -f dangling=true)

docker pull # 下载镜像
docker inspect # 查看镜像详细信息
docker search # 搜索镜像
docker build --progress plain -t test .
# 运行镜像
docker run -it containr:v1 bash
# 使用镜像 nginx:latest，-d 以后台模式启动一个容器，-p 将容器的 80 端口映射到主机的 80 端口，-v 将主机的目录 /data 映射到容器的 /data。
docker run -p 80:80 -v /data:/data -d nginx:latest
# 保存镜像
docker save ID > xxx.tar # 导出镜像
docker load < xxx.tar # 导入镜像
# 保存容器
docker export ID > xxx.tar # 导出容器
docker import xxx.tar containr:v1 # 导入容器
# 保存镜像
docker save -o php7-fpm.tar harbor.k8s/library/php7-fpm
# 导入镜像
docker import php7-fpm.tar harbor.k8s/library/php7-fpm:latest
# 推送
docker push harbor.k8s/library/php7-fpm:latest
```

## docker-compose

```bash
# Start services
docker-compose up -d nginx php-fpm mysql redis
# Stop all service
docker-compose stop
# build
docker-compose build php-fpm
# Lists containers
docker-compose ps
# mysql
docker-compose run --rm mysql bash
# php-cli
docker-compose run --rm php-cli bash
# redis
docker-compose run --rm redis redis-cli -h redis
# node
docker-compose run --rm node sh
# bind port
docker-compose run --rm -p 3001:3001 node sh
# golang
docker-compose run --rm golang bash
```
