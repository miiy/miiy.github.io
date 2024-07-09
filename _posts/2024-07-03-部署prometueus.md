---
layout: post
title: "部署 prometueus"
date: 2024-07-03
tags: linux
---

公网服务器

```
frontend

backend

monitoring
  grafana: :3000
  prometheus: 9090
  pushgateway: 0.0.0.0:9091
  node-exporter: :9100
  alertmanager: :9093
```

内网服务器

```
frontend

backend
  node-exporer: :9100
  node-exporter-pusher:
```

## prometueus

创建 TLS 证书

Prometheus Server And TLS: <https://o11y.eu/blog/prometheus-server-tls/>

basic auth and tls: <https://github.com/prometheus/exporter-toolkit/blob/master/docs/web-config.yml>

```bash
cd /data/prometheus/conf

openssl req -new -newkey rsa:2048 -days 365 -nodes -x509 -keyout prometheus.key -out prometheus.crt -subj "/C=CN/ST=Beijing/L=Beijing/O=exampleOrg/CN=prometheus" -addext "subjectAltName = DNS:localhost"
```

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
    - static_configs:
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

创建网络

```bash
docker network create monitoring
```

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
    prom/prometheus:v2.53.0 --config.file=/etc/prometheus/prometheus.yml --web.config.file=/etc/prometheus/web-config.yml
```

--user 1000

```
open /prometheus/queries.active: permission denied
```

<https://github.com/prometheus/prometheus/issues/5976>


访问：<https://127.0.0.1:9091/>


## grafana

```bash
docker run -d --name=grafana --restart=always \
    --user "$(id -u)" \
    -v /data/grafana/data:/var/lib/grafana \
    --net monitoring \
    grafana/grafana:11.1.0
```

```bash
docker network connect frontend grafana
```

## node-exporter

<https://prometheus.io/docs/guides/cadvisor/>

```bash
docker run -d --name node-exporter --restart=always \
    -v "/proc:/host/proc:ro" \
    -v "/sys:/host/sys:ro" \
    -v "/:/rootfs:ro" \
    --net monitoring \
    prom/node-exporter:v1.8.1
```

## push-gateway

pushgateway: <https://github.com/prometheus/pushgateway>

创建 TLS 证书

```bash
cd /data/pushgateway/conf/

openssl req -new -newkey rsa:2048 -days 365 -nodes -x509 -keyout pushgateway.key -out pushgateway.crt -subj "/C=CN/ST=Beijing/L=Beijing/O=exampleOrg/CN=pushgateway" -addext "subjectAltName = DNS:pushgateway"
cp pushgateway.crt /data/prometheus/conf
```

web-config.yaml

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
    -p 9091:9091 \
    -v /data/pushgateway/config:/pushgateway/config \
    --network monitoring \
    prom/pushgateway:v1.9.0 --web.config.file=/pushgateway/config/web-config.yaml
```

发送 metrics

```bash
echo "my_metric 2"  | gzip | curl --insecure -u username:password -H 'Content-Encoding: gzip' --data-binary @- https://pushgateway:9091/metrics/job/test-job/instance/nodename
curl --insecure -u http:password -X DELETE  https://pushgateway:9091/metrics/job/test-job/instance/nodename
```

## node-exporter-pusher

将内网的 node-exporter 发送到 pushgateway

```bash
docker run --name node-exporter-pusher -d --restart=always \
    -v /home/debian/data/node-exporter-pusher/.env:/root/.env \
    --network backend \
    node-exporter-pusher
```

## PushProx

<https://github.com/prometheus-community/PushProx>

prometheus-community-PushProx介绍：<https://blog.csdn.net/doyzfly/article/details/120752044>

## alertmanager

```bash
docker run --name alertmanager -d --restart=always \
    --network monitoring \
    -v /data/alertmanager/config:/alertmanager/config \
    prom/alertmanager:v0.27.0 --config.file=/etc/alertmanager/alertmanager.yml
```


