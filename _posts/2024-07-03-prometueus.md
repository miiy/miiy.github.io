---
layout: post
title: "Prometueus 监控"
date: 2024-07-03
tags: linux
---

## 拓补结构

公网服务器

```
frontend netrowk

backend netrowk

monitoring netrowk
  grafana: :3000
  prometheus: :9090 
  pushgateway: 0.0.0.0:9091
  pushprox-proxy: :8080 pushprox.example.com
  node-exporter: :9100
  smartctl-exporter: :9633
  alertmanager: :9093
```

内网服务器

```
host
  node-exporer: 127.0.0.1:9100
  smartctl-exporter: 127.0.0.1:9633
  pushprox-client:
```

## prometueus

### 创建网络

```bash
docker network create monitoring
```

### 创建 TLS 证书

Prometheus Server And TLS: <https://o11y.eu/blog/prometheus-server-tls/>

basic auth and tls: <https://github.com/prometheus/exporter-toolkit/blob/master/docs/web-config.yml>

```bash
cd /data/prometheus/config

openssl req -new -newkey rsa:2048 -days 365 -nodes -x509 -keyout prometheus.key -out prometheus.crt -subj "/C=CN/ST=Beijing/L=Beijing/O=exampleOrg/CN=prometheus" -addext "subjectAltName = DNS:localhost"
```

### 准备配置文件

web-config.yml

```yml
# TLS and basic authentication configuration example.
#
# Additionally, a certificate and a key file are needed.
tls_server_config:
  cert_file: prometheus.crt
  key_file: prometheus.key

# Usernames and passwords required to connect.
# Passwords are hashed with bcrypt: https://github.com/prometheus/exporter-toolkit/blob/master/docs/web-configuration.md#about-bcrypt.
basic_auth_users:
  prometheus:
```

http basic auth 密码创建: <https://o11y.tools/pwgen/>

prometheus.yml

```yml
# my global config
global:
  scrape_interval: 15s # Set the scrape interval to every 15 seconds. Default is every 1 minute.
  evaluation_interval: 15s # Evaluate rules every 15 seconds. The default is every 1 minute.
  # scrape_timeout is set to the global default (10s).

# Alertmanager configuration
alerting:
  alertmanagers:
      static_configs:
        - targets:
            - alertmanager:9093

# Load rules once and periodically evaluate them according to the global 'evaluation_interval'.
rule_files:
  # - "first_rules.yml"
  # - "second_rules.yml"
  - "alert.yml"

# A scrape configuration containing exactly one endpoint to scrape:
# Here it's Prometheus itself.
scrape_configs:
  # The job name is added as a label `job=<job_name>` to any timeseries scraped from this config.
  - job_name: "prometheus"

    # metrics_path defaults to '/metrics'
    # scheme defaults to 'http'.
    scheme: https
    tls_config:
      ca_file: prometheus.crt
    basic_auth:
      username: prometheus
      password: ""
    static_configs:
      - targets: ["localhost:9090"]
  - job_name: node
    static_configs:
      - targets: ['node-exporter:9100']
  - job_name: node-proxy
    proxy_url: http://pushprox-proxy:8080/
    static_configs:
      - targets: ['x-client:9100']  # Presuming the FQDN of the client is "client".
  - job_name: "pushgateway"
    scheme: https
    basic_auth:
      username: prometheus
      password: ""
    tls_config:
      ca_file: pushgateway.crt
    static_configs:
      - targets: ['pushgateway:9091']
```

alert.yml

{% raw %}
```yml
groups:
- name: Instances
  rules:
  - alert: InstanceDown
    expr: up == 0
    for: 1m
    labels:
      severity: page
    # Prometheus templates apply here in the annotation and label fields of the alert.
    annotations:
      description: '{{ $labels.instance }} of job {{ $labels.job }} has been down for more than 1 minutes.'
      summary: 'Instance {{ $labels.instance }} down'
```
{% endraw %}


启动 prometheus

```bash
docker run --name prometheus -d --restart=always \
    --user "$(id -u)" \
    -v /data/prometheus/config/prometheus.yml:/etc/prometheus/prometheus.yml \
    -v /data/prometheus/config/web-config.yml:/etc/prometheus/web-config.yml \
    -v /data/prometheus/config/prometheus.crt:/etc/prometheus/prometheus.crt \
    -v /data/prometheus/config/prometheus.key:/etc/prometheus/prometheus.key \
    -v /data/prometheus/config/pushgateway.crt:/etc/prometheus/pushgateway.crt \
    -v /data/prometheus/config/alert.yml:/etc/prometheus/alert.yml \
    -v /data/prometheus/data:/prometheus \
    --net monitoring \
    prom/prometheus:v3.13.1 --config.file=/etc/prometheus/prometheus.yml --web.config.file=/etc/prometheus/web-config.yml
```

如果需要 pushgateway 请先生成证书

加参数 --user，--user 1000 否则会报错：open /prometheus/queries.active: permission denied，参考 <https://github.com/prometheus/prometheus/issues/5976>

访问：<https://127.0.0.1:9091/>


## grafana

```bash
docker run -d --name=grafana --restart=always \
    --user "$(id -u)" \
    -v /data/grafana/data:/var/lib/grafana \
    --net monitoring \
    grafana/grafana:11.6
```

```bash
docker network connect frontend grafana
```

## exporter

### node-exporter

<https://prometheus.io/docs/guides/cadvisor/>

```bash
docker run -d --name node-exporter --restart=always \
    -v "/proc:/host/proc:ro" \
    -v "/sys:/host/sys:ro" \
    -v "/:/rootfs:ro" \
    -p 127.0.0.1:9100:9100 \
    --net monitoring \
    prom/node-exporter:v1.12.0
```

### smartctl exporter

```bash
docker run -d \
  --name=smartctl-exporter --restart=always \
  --privileged \
  --user=root \
  -v /dev:/dev:ro \
  -p 127.0.0.1:9633:9633 \
  --network monitoring \
  prometheuscommunity/smartctl-exporter:v0.14.0 \
  --smartctl.interval=5m \
  --smartctl.device=/dev/nvme0 \
  --smartctl.device=/dev/sda \
  --smartctl.device="/dev/sdb;sat"
```

### dcgm exporter

```bash
docker run -d --gpus all --cap-add SYS_ADMIN --name dcgm-exporter --restart=always \
    -p 127.0.0.1:9400:9400 \
    --network monitoring \
    nvidia/dcgm-exporter:4.8.3
```

### cadvisor

```bash
sudo docker run \
  --volume=/:/rootfs:ro \
  --volume=/var/run:/var/run:ro \
  --volume=/sys:/sys:ro \
  --volume=/var/lib/docker/:/var/lib/docker:ro \
  --volume=/dev/disk/:/dev/disk:ro \
  --restart=always \
  --network monitoring \
  --publish=127.0.0.1:9880:8080 \
  --detach=true \
  --name=cadvisor \
  --privileged \
  --device=/dev/kmsg \
  ghcr.io/google/cadvisor:0.60.5
```

## push-gateway

pushgateway: <https://github.com/prometheus/pushgateway>

创建 TLS 证书

```bash
cd /data/pushgateway/config/

openssl req -new -newkey rsa:2048 -days 365 -nodes -x509 -keyout pushgateway.key -out pushgateway.crt -subj "/C=CN/ST=Beijing/L=Beijing/O=exampleOrg/CN=pushgateway" -addext "subjectAltName = DNS:pushgateway"
cp pushgateway.crt /data/prometheus/config
```

web-config.yml

```yaml
# TLS and basic authentication configuration example.
#
# Additionally, a certificate and a key file are needed.
tls_server_config:
  cert_file: pushgateway.crt
  key_file: pushgateway.key

# Usernames and passwords required to connect.
# Passwords are hashed with bcrypt: https://github.com/prometheus/exporter-toolkit/blob/master/docs/web-configuration.md#about-bcrypt.

basic_auth_users:
  prometheus: 
  http: 
```

启动 pushgateway

```bash
docker run --name pushgateway -d --restart=always \
    --user "$(id -u)" \
    -v /data/pushgateway/config:/pushgateway/config \
    --network monitoring \
    prom/pushgateway:v1.11.3 --web.config.file=/pushgateway/config/web-config.yml
```
    -p 9091:9091 \

发送 metrics

```bash
echo "my_metric 2"  | gzip | curl --insecure -u username:password -H 'Content-Encoding: gzip' --data-binary @- https://pushgateway:9091/metrics/job/test-job/instance/nodename
curl --insecure -u http:password -X DELETE https://pushgateway:9091/metrics/job/test-job/instance/nodename
```

## PushProx

通过 PushProx 抓取内网的 node-exporter

<https://github.com/prometheus-community/PushProx>

prometheus-community-PushProx介绍：<https://blog.csdn.net/doyzfly/article/details/120752044>

```bash
docker pull prometheuscommunity/pushprox:master
```

server

```bash
docker run --name pushprox-proxy -d --restart=always \
    --network monitoring \
    prometheuscommunity/pushprox:v0.2.0

docker network connect frontend pushprox-proxy
```

curl {pushprox-proxy}:8080/metrics

创建自签名证书

配置 nginx

client

```bash
sudo docker run --name pushprox-client -d --restart=always \
    --entrypoint /app/pushprox-client \
    --network host \
    --add-host pushprox.example.com:39.100.100.100 \
    --add-host dell-client:127.0.0.1 \
    -v /srv/docker/pushprox-client/certs/:/app/certs/ \
    prometheuscommunity/pushprox:v0.2.0 \
    --fqdn=dell-client \
    --proxy-url=https://pushprox.example.com/ \
    --tls.cacert=/app/certs/ca.crt \
    --tls.cert=/app/certs/client.crt \
    --tls.key=/app/certs/client.key
```

## alertmanager

```bash
docker run --name alertmanager -d --restart=always \
    --network monitoring \
    -v /data/alertmanager/config/alertmanager.yml:/etc/alertmanager/alertmanager.yml \
    -v /data/alertmanager/config/web-config.yml:/etc/alertmanager/web-config.yml \
    prom/alertmanager:v0.33.0 --config.file=/etc/alertmanager/alertmanager.yml --web.config.file=/etc/alertmanager/web-config.yml
```
    -p 127.0.0.1:9093:9093 \

## alertmanager-webhook-adapter

```bash
docker run --name alertmanager-webhook-adapter -d --restart=always \
    --network monitoring \
    -e TZ=Asia/Shanghai \
    bougou/alertmanager-webhook-adapter:latest \
    --signature "Docxlib"
```
    -p 127.0.0.1:8090:8090 \
    -v /data/alertmanager-webhook-adapter/templates:/templates \

模版参数：
    --tmpl-lang zh --tmpl-dir /templates

## alertmanager 配置

web-config.yml

```yaml
# TLS and basic authentication configuration example.
#
# Additionally, a certificate and a key file are needed.

# Usernames and passwords required to connect.
# Passwords are hashed with bcrypt: https://github.com/prometheus/exporter-toolkit/blob/master/docs/web-configuration.md#about-bcrypt.
basic_auth_users:
  prometheus:
```

alertmanager.yml

```yaml
route:
  group_by: ['alertname']
  group_wait: 30s
  group_interval: 5m
  repeat_interval: 1h
  receiver: 'feishu'
receivers:
  - name: 'web.hook'
    webhook_configs:
      - url: 'http://127.0.0.1:5001/'
  - name: 'feishu'
    webhook_configs:
      - url: 'http://alertmanager-webhook-adapter:8090/webhook/send?channel_type=feishu&token=机器人ID'
        # 发送恢复通知
        send_resolved: true
inhibit_rules:
  - source_match:
      severity: 'critical'
    target_match:
      severity: 'warning'
    equal: ['alertname', 'dev', 'instance']
```


## PromQL basic example

<https://prometheus.io/docs/prometheus/latest/querying/basics/>

```text
# 查询所有 http_requests_total 时间序列，返回每条时间序列最新的样本
http_requests_total

# 在 {} 中按标签精确匹配
http_requests_total{job="prometheus",group="canary"}

# 标签过滤：!= 表示不等于
http_requests_total{method!="GET"}

# 标签正则匹配：=~ 表示匹配，!~ 表示不匹配
http_requests_total{environment=~"staging|testing|development",method!~"GET|POST"}

# 查询指定时间点的数据
http_requests_total @ 1609746000

# 查询指定时间点前 5 分钟的数据
http_requests_total @ 1609746000 offset 5m

# 查询过去 5 分钟的数据（Range Vector）
http_requests_total[5m]

# 查询过去 5 分钟，每隔 30 秒一个数据点
http_requests_total[5m:30s]

# 计算过去 5 分钟的平均每秒增长速率（Counter 最常用）
rate(http_requests_total[5m])

# 计算过去 5 分钟内总共增加了多少
increase(http_requests_total[5m])

# 当前值减去 5 分钟前的值
delta(http_requests_total[5m])

# 统计时间序列数量
count(http_requests_total)

# 求和
sum(http_requests_total)

# 按标签分组求和
sum by (job) (http_requests_total)

# 忽略其他标签，仅按 job、instance 分组
sum by (job, instance) (http_requests_total)

# 平均值
avg(http_requests_total)

# 最大值、最小值
max(http_requests_total)
min(http_requests_total)

# 排序
sort(http_requests_total)
sort_desc(http_requests_total)

# 取值最大的前 10 条
topk(10, http_requests_total)

# 取值最小的前 10 条
bottomk(10, http_requests_total)

# 当前时间（Unix 时间戳）
time()

# 返回不存在的时间序列（常用于告警）
absent(http_requests_total{job="api"})
```

## References

Prometheus+Grafana监控MySQL_ITPUB博客：<https://blog.itpub.net/69982604/viewspace-2743207/>

Introduction | prometheus-book：<https://yunlzheng.gitbook.io/prometheus-book>

Kubernetes技术栈-K8s|Docker|Istio|Python|Golang|云原生：<https://www.k8stech.net/>

监控神器：Prometheus 轻松入门，真香！：<https://mp.weixin.qq.com/s/W38FcwGmwPj1tp_87FVC1A>

<https://github.com/phper95/pkg/tree/master/prome>

<https://github.com/sretalk/prometheus-rules>