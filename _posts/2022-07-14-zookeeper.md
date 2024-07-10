---
layout: post
title: "安装zookeeper"
date: 2022-07-14
tags: linux
---

## 安装 ZooKeeper

ZooKeeper Getting Started Guide <https://zookeeper.apache.org/doc/current/zookeeperStarted.html>


```bash
cd /usr/local/src
wget https://ftp.tsukuba.wide.ad.jp/software/apache/zookeeper/zookeeper-3.6.2/apache-zookeeper-3.6.2-bin.tar.gz
tar -zxvf apache-zookeeper-3.6.2-bin.tar.gz
mv apache-zookeeper-3.6.2-bin ../
cd ../
ln -s apache-zookeeper-3.6.2-bin zookeeper
cp zookeeper/conf/zoo_sample.cfg zookeeper/conf/zoo.cfg
```

修改配置文件 dataDir


## 启动单机模式

```bash
$ bin/zkServer.sh start
/usr/bin/java
ZooKeeper JMX enabled by default
Using config: /usr/local/zookeeper/bin/../conf/zoo.cfg
Starting zookeeper ... STARTED

$ bin/zkServer.sh
/usr/bin/java
ZooKeeper JMX enabled by default
Using config: /usr/local/zookeeper/bin/../conf/zoo.cfg
Usage: bin/zkServer.sh [--config <conf-dir>] {start|start-foreground|stop|version|restart|status|print-cmd}
$ bin/zkServer.sh status
/usr/bin/java
ZooKeeper JMX enabled by default
Using config: /usr/local/zookeeper/bin/../conf/zoo.cfg
Client port found: 2181. Client address: localhost. Client SSL: false.
Mode: standalone
```

连接ZooKeeper

```bash
$ bin/zkCli.sh -server 127.0.0.1:2181
```

## 基本操作

```bash
[zk: 127.0.0.1:2181(CONNECTED) 0] help
ZooKeeper -server host:port -client-configuration properties-file cmd args
	addWatch [-m mode] path # optional mode is one of [PERSISTENT, PERSISTENT_RECURSIVE] - default is PERSISTENT_RECURSIVE
	addauth scheme auth
	close
	config [-c] [-w] [-s]
	connect host:port
	create [-s] [-e] [-c] [-t ttl] path [data] [acl]
	delete [-v version] path
	deleteall path [-b batch size]
	delquota [-n|-b] path
	get [-s] [-w] path
	getAcl [-s] path
	getAllChildrenNumber path
	getEphemerals path
	history
	listquota path
	ls [-s] [-w] [-R] path
	printwatches on|off
	quit
	reconfig [-s] [-v version] [[-file path] | [-members serverID=host:port1:port2;port3[,...]*]] | [-add serverId=host:port1:port2;port3[,...]]* [-remove serverId[,...]*]
	redo cmdno
	removewatches path [-c|-d|-a] [-l]
	set [-s] [-v version] path data
	setAcl [-s] [-v version] [-R] path acl
	setquota -n|-b val path
	stat [-w] path
	sync path
	version
```

```bash
[zk: 127.0.0.1:2181(CONNECTED) 23] ls /
[zookeeper]
[zk: 127.0.0.1:2181(CONNECTED) 24] create /zktest testdata
Created /zktest
[zk: 127.0.0.1:2181(CONNECTED) 25] ls /
[zktest, zookeeper]
[zk: 127.0.0.1:2181(CONNECTED) 26] get /zktest
testdata
[zk: 127.0.0.1:2181(CONNECTED) 29] create /zktest/a
Created /zktest/a
[zk: 127.0.0.1:2181(CONNECTED) 30] get /zktest/a
null
# 创建节点
[zk: 127.0.0.1:2181(CONNECTED) 31] create /zktest/b b
Created /zktest/b
[zk: 127.0.0.1:2181(CONNECTED) 32] get /zktest/b
b
# 删除节点
[zk: 127.0.0.1:2181(CONNECTED) 33] delete /zktest/b
[zk: 127.0.0.1:2181(CONNECTED) 34] ls /zktest
[a]
[zk: 127.0.0.1:2181(CONNECTED) 36] delete /zktest
Node not empty: /zktest
# 删除所有节点
[zk: 127.0.0.1:2181(CONNECTED) 37] deleteall /zktest
```

