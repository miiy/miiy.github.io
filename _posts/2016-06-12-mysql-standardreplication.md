---
layout: post
title: "MySQL StandardReplication"
date: 2016-06-12
category: mysql
tags: mysql
---

#Replication


####实验环境：

| Master/Slave | Platfrom | IP | APP Version |
| -- | -- | -- | -- |
| Master | Centos7_X86-64 | 192.168.31.200 | mariadb-10.0.22 |
| Slave | Centos7_X86-64 | 192.168.31.201 | mariadb-10.0.22 |


####Master配置：

在Master上创建有复制权限的用户

```sql
GRANT REPLICATION SLAVE ON *.* TO 'slave'@'192.168.31.201' IDENTIFIED BY '123456';
FLUSH PRIVILEGES;
```

> 需要注意，有一些系统配置选项可能会影响主从复制，查看下面的变量以避免发生问题:

> skip-networking,如果 "skip-networking=1"，则服务器将限制只能由localhost连接，阻止其他机器远程连到此服务器上。

> bind_address,类似地，如果服务器只监听127.0.0.1(localhost)的TCP/IP连接,则远程的 slave也不能连接到此服务器。

修改配置文件 /etc/my.cnf [mysqld]段

启用二进制日志，配置server-id

server_id值可以是整数型的数字(1 ~ 2^31-1)，在同一个复制组中的每台服务器的server_id都必须是唯一的。

```
log-bin=mysql-bin
server-id=1
```

修改完后重启Mariadb

####Slave配置：

修改配置文件 /etc/my.cnf [mysqld]段：

关闭二进制日志，启用中继日志，配置server-id

```
#log-bin=mysql-bin
relay-log=mysql-bin
server-id=2
```

修改完后重启Mariadb

确保中继日志已开启

```sql
MariaDB [(none)]> SHOW GLOBAL VARIABLES LIKE 'relay_log';
+---------------+-------+
| Variable_name | Value |
+---------------+-------+
| relay_log     |       |
+---------------+-------+
1 row in set (0.00 sec)
```

连接Master服务器

```
MariaDB [(none)]> CHANGE MASTER TO MASTER_USER='slave', MASTER_HOST='192.168.31.200', MASTER_PASSWORD='123456';
MariaDB [(none)]> SHOW SLAVE STATUS\G;
*************************** 1. row ***************************
               Slave_IO_State:
                  Master_Host: 192.168.31.200
                  Master_User: slave
                  Master_Port: 3306
                Connect_Retry: 60
              Master_Log_File:
          Read_Master_Log_Pos: 4
               Relay_Log_File: MiWiFi-R1D-srv-relay-bin.000001
                Relay_Log_Pos: 4
        Relay_Master_Log_File:
             Slave_IO_Running: No
            Slave_SQL_Running: No
              Replicate_Do_DB:
          Replicate_Ignore_DB:
           Replicate_Do_Table:
       Replicate_Ignore_Table:
      Replicate_Wild_Do_Table:
  Replicate_Wild_Ignore_Table:
                   Last_Errno: 0
                   Last_Error:
                 Skip_Counter: 0
          Exec_Master_Log_Pos: 0
              Relay_Log_Space: 0
              Until_Condition: None
               Until_Log_File:
                Until_Log_Pos: 0
           Master_SSL_Allowed: No
           Master_SSL_CA_File:
           Master_SSL_CA_Path:
              Master_SSL_Cert:
            Master_SSL_Cipher:
               Master_SSL_Key:
        Seconds_Behind_Master: NULL
Master_SSL_Verify_Server_Cert: No
                Last_IO_Errno: 0
                Last_IO_Error:
               Last_SQL_Errno: 0
               Last_SQL_Error:
  Replicate_Ignore_Server_Ids:
             Master_Server_Id: 0
1 row in set (0.00 sec)

```

在输出的结果当中:

```
Slave_IO_Running: No
Slave_SQL_Running: No
```

表明slave还没有开始复制过程。日志的位置为4而不是0，这是因为0只是日志文件的开始位置，并不是日志位置。实际上，MySQL知道的第一个事件的位置是4。

在Slave上面启动复制线程:

```
MariaDB [(none)]> START SLAVE;
MariaDB [(none)]> SHOW SLAVE STATUS\G;
*************************** 1. row ***************************
               Slave_IO_State: Waiting for master to send event
                  Master_Host: 192.168.31.200
                  Master_User: slave
                  Master_Port: 3306
                Connect_Retry: 60
              Master_Log_File: mysql-bin.000003
          Read_Master_Log_Pos: 245
               Relay_Log_File: mysql-bin.000004
                Relay_Log_Pos: 529
        Relay_Master_Log_File: mysql-bin.000003
             Slave_IO_Running: Yes
            Slave_SQL_Running: Yes
              Replicate_Do_DB:
          Replicate_Ignore_DB:
           Replicate_Do_Table:
       Replicate_Ignore_Table:
      Replicate_Wild_Do_Table:
  Replicate_Wild_Ignore_Table:
                   Last_Errno: 0
                   Last_Error:
                 Skip_Counter: 0
          Exec_Master_Log_Pos: 245
              Relay_Log_Space: 1101
              Until_Condition: None
               Until_Log_File:
                Until_Log_Pos: 0
           Master_SSL_Allowed: No
           Master_SSL_CA_File:
           Master_SSL_CA_Path:
              Master_SSL_Cert:
            Master_SSL_Cipher:
               Master_SSL_Key:
        Seconds_Behind_Master: 0
Master_SSL_Verify_Server_Cert: No
                Last_IO_Errno: 0
                Last_IO_Error:
               Last_SQL_Errno: 0
               Last_SQL_Error:
  Replicate_Ignore_Server_Ids:
             Master_Server_Id: 1
1 row in set (0.01 sec)
```

如果配置过程中出错
Could not initialize master info structure, more error messages can be found in the MySQL error log
通过`RESET SLAVE;`重新设置。


在master上，可以看到slave的I/O线程创建的连接：

```
MariaDB [(none)]> SHOW PROCESSLIST\G;
*************************** 1. row ***************************
      Id: 3
    User: slave
    Host: 192.168.31.201:40993
      db: NULL
 Command: Binlog Dump
    Time: 322
   State: Master has sent all binlog to slave; waiting for binlog to be updated
    Info: NULL
Progress: 0.000
*************************** 2. row ***************************
      Id: 4
    User: root
    Host: localhost
      db: NULL
 Command: Query
    Time: 0
   State: NULL
    Info: SHOW PROCESSLIST
Progress: 0.000
2 rows in set (0.00 sec)
```

在slave上运行

```
MariaDB [(none)]> SHOW PROCESSLIST\G;
*************************** 1. row ***************************
      Id: 2
    User: root
    Host: localhost
      db: NULL
 Command: Query
    Time: 0
   State: NULL
    Info: SHOW PROCESSLIST
Progress: 0.000
*************************** 2. row ***************************
      Id: 3
    User: system user
    Host:
      db: NULL
 Command: Connect
    Time: 465
   State: Waiting for master to send event
    Info: NULL
Progress: 0.000
*************************** 3. row ***************************
      Id: 4
    User: system user
    Host:
      db: NULL
 Command: Connect
    Time: 4825
   State: Slave has read all relay log; waiting for the slave I/O thread to upda                                                                                        te it
    Info: NULL
Progress: 0.000
3 rows in set (0.00 sec)
```

####测试

在Master上面创建一个数据库:

```
MariaDB [(none)]> CREATE DATABASE testdb1;
```

在Slave上查看：

```
MariaDB [(none)]> SHOW DATABASES;
+--------------------+
| Database           |
+--------------------+
| information_schema |
| mysql              |
| performance_schema |
| testdb1            |
+--------------------+
4 rows in set (0.00 sec)
```


####同步已有的数据

master主机打开两个shell窗口，一个shell，一个连接mariadb（锁定表期间不能关闭会话，一关闭会话就会释放锁。）

在master上中阻断mariadb的写操作：

```
MariaDB [(none)]> FLUSH TABLES WITH READ LOCK;
```

在master上导出数据库并传输到slave上。

```
[root@localhost ~]# mysqldump -u root -p testdb1 --lock-all-tables > testdb1.sql;
[root@localhost ~]# scp testdb1.sql root@192.168.31.201:/root/
[root@localhost ~]# mysql -u root -p
```

在master上获取master的二进制日志坐标

```
MariaDB [(none)]> SHOW MASTER STATUS;
+------------------+----------+--------------+------------------+
| File             | Position | Binlog_Do_DB | Binlog_Ignore_DB |
+------------------+----------+--------------+------------------+
| mysql-bin.000003 |      334 |              |                  |
+------------------+----------+--------------+------------------+
1 row in set (0.00 sec)
```

在slave上导入

```
[root@localhost ~]# mysql -u root -p testdb1 < testdb1.sql
[root@localhost ~]# mysql -u root -p
MariaDB [(none)]>
CHANGE MASTER TO
  MASTER_HOST='192.168.31.200',
  MASTER_USER='slave',
  MASTER_PASSWORD='123456',
  MASTER_PORT=3306,
  MASTER_LOG_FILE='mysql-bin.000003',
  MASTER_LOG_POS=334,
  MASTER_CONNECT_RETRY=10;

MariaDB [(none)]> START SLAVE;
MariaDB [(none)]> SHOW SLAVE STATUS\G;
*************************** 1. row ***************************
               Slave_IO_State: Waiting for master to send event
                  Master_Host: 192.168.31.200
                  Master_User: slave
                  Master_Port: 3306
                Connect_Retry: 10
              Master_Log_File: mysql-bin.000003
          Read_Master_Log_Pos: 405
               Relay_Log_File: mysql-bin.000002
                Relay_Log_Pos: 600
        Relay_Master_Log_File: mysql-bin.000003
             Slave_IO_Running: Yes
            Slave_SQL_Running: Yes
              Replicate_Do_DB:
          Replicate_Ignore_DB:
           Replicate_Do_Table:
       Replicate_Ignore_Table:
      Replicate_Wild_Do_Table:
  Replicate_Wild_Ignore_Table:
                   Last_Errno: 0
                   Last_Error:
                 Skip_Counter: 0
          Exec_Master_Log_Pos: 405
              Relay_Log_Space: 888
              Until_Condition: None
               Until_Log_File:
                Until_Log_Pos: 0
           Master_SSL_Allowed: No
           Master_SSL_CA_File:
           Master_SSL_CA_Path:
              Master_SSL_Cert:
            Master_SSL_Cipher:
               Master_SSL_Key:
        Seconds_Behind_Master: 0
Master_SSL_Verify_Server_Cert: No
                Last_IO_Errno: 0
                Last_IO_Error:
               Last_SQL_Errno: 0
               Last_SQL_Error:
  Replicate_Ignore_Server_Ids:
             Master_Server_Id: 1
1 row in set (0.00 sec)
```

在master上，解锁刚才的锁定

```
MariaDB [(none)]> UNLOCK TABLES;
```


###双主模式配置：

####实验环境：

| Master1/Master2 | Platfrom | IP | APP Version |
| -- | -- | -- | -- |
| Master | Centos7_X86-64 | 192.168.31.200 | mariadb-5.5.44 |
| Slave | Centos7_X86-64 | 192.168.31.201 | mariadb-5.5.44 |

####配置双主

####Master1配置

/etc/my.cnf [mysqld]

```
log-bin=mysql-bin
relay-log=mysql-bin
auto-increment-offset = 1 //起始值
auto-increment-increment = 2 //步长
server-id=1
```

重启Mariadb

确保中继日志选项开启，确保自动增长选项已开启

```
MariaDB [(none)]> SHOW GLOBAL VARIABLES LIKE 'relay_log';
+---------------+-----------+
| Variable_name | Value     |
+---------------+-----------+
| relay_log     | mysql-bin |
+---------------+-----------+
1 row in set (0.00 sec)

MariaDB [(none)]> SHOW GLOBAL VARIABLES LIKE 'auto_inc%';
+--------------------------+-------+
| Variable_name            | Value |
+--------------------------+-------+
| auto_increment_increment | 2     |
| auto_increment_offset    | 1     |
+--------------------------+-------+
2 rows in set (0.00 sec)

```

创建有复制权限的用户：

```
MariaDB [mysql]> GRANT REPLICATION SLAVE ON *.* TO 'replication'@'192.168.31.201' IDENTIFIED BY '123456';
MariaDB [(none)]> FLUSH PRIVILEGES;
```

###Master2配置

```
log-bin=mysql-bin
relay-log=mysql-bin
auto-increment-offset = 2 //起始值
auto-increment-increment = 2 //步长
server-id=2
```

重启Mariadb

确保中继日志选项开启确保自动增长选项已开启

```
MariaDB [(none)]> SHOW GLOBAL VARIABLES LIKE 'relay_log';
+---------------+-----------+
| Variable_name | Value     |
+---------------+-----------+
| relay_log     | mysql-bin |
+---------------+-----------+
1 row in set (0.01 sec)

MariaDB [(none)]> SHOW GLOBAL VARIABLES LIKE 'auto_inc%';
+--------------------------+-------+
| Variable_name            | Value |
+--------------------------+-------+
| auto_increment_increment | 2     |
| auto_increment_offset    | 2     |
+--------------------------+-------+
2 rows in set (0.00 sec)

```

创建有复制权限的用户：

```
MariaDB [mysql]> GRANT REPLICATION SLAVE ON *.* TO 'replication'@'192.168.31.201' IDENTIFIED BY '123456';
MariaDB [(none)]> FLUSH PRIVILEGES;
```

如果此时两台服务器均为新建立，且无其它写入操作，各服务器只需记录当前自己二进制日志文件及事件位置，以之作为另外的服务器复制起始位置即可

Master1

```
MariaDB [mysql]> SHOW MASTER STATUS;
+------------------+----------+--------------+------------------+
| File             | Position | Binlog_Do_DB | Binlog_Ignore_DB |
+------------------+----------+--------------+------------------+
| mysql-bin.000007 |      847 |              |                  |
+------------------+----------+--------------+------------------+
1 row in set (0.00 sec)
```

Master2

```
MariaDB [mysql]> SHOW MASTER STATUS;
+------------------+----------+--------------+------------------+
| File             | Position | Binlog_Do_DB | Binlog_Ignore_DB |
+------------------+----------+--------------+------------------+
| mysql-bin.000006 |      837 |              |                  |
+------------------+----------+--------------+------------------+
1 row in set (0.00 sec)

MariaDB [mysql]>

```

各服务器接下来指定对另一台服务器为自己的主服务器

Master1
```
MariaDB [(none)]>
CHANGE MASTER TO
  MASTER_HOST='192.168.31.201',
  MASTER_USER='replication',
  MASTER_PASSWORD='123456',
  MASTER_PORT=3306,
  MASTER_LOG_FILE='mysql-bin.000006',
  MASTER_LOG_POS=837,
  MASTER_CONNECT_RETRY=10;
```

Master2
```
MariaDB [(none)]>
CHANGE MASTER TO
  MASTER_HOST='192.168.31.200',
  MASTER_USER='replication',
  MASTER_PASSWORD='123456',
  MASTER_PORT=3306,
  MASTER_LOG_FILE='mysql-bin.000007',
  MASTER_LOG_POS=847,
  MASTER_CONNECT_RETRY=10;
```

```
MariaDB [mysql]> START SLAVE;
Query OK, 0 rows affected (0.01 sec)

MariaDB [mysql]> SHOW SLAVE STATUS\G;
*************************** 1. row ***************************
               Slave_IO_State: Connecting to master
                  Master_Host: 192.168.31.201
                  Master_User: replication
                  Master_Port: 3306
                Connect_Retry: 10
              Master_Log_File: mysql-bin.000006
          Read_Master_Log_Pos: 837
               Relay_Log_File: mysql-bin.000001
                Relay_Log_Pos: 4
        Relay_Master_Log_File: mysql-bin.000006
             Slave_IO_Running: Connecting
            Slave_SQL_Running: Yes
              Replicate_Do_DB:
          Replicate_Ignore_DB:
           Replicate_Do_Table:
       Replicate_Ignore_Table:
      Replicate_Wild_Do_Table:
  Replicate_Wild_Ignore_Table:
                   Last_Errno: 0
                   Last_Error:
                 Skip_Counter: 0
          Exec_Master_Log_Pos: 837
              Relay_Log_Space: 245
              Until_Condition: None
               Until_Log_File:
                Until_Log_Pos: 0
           Master_SSL_Allowed: No
           Master_SSL_CA_File:
           Master_SSL_CA_Path:
              Master_SSL_Cert:
            Master_SSL_Cipher:
               Master_SSL_Key:
        Seconds_Behind_Master: NULL
Master_SSL_Verify_Server_Cert: No
                Last_IO_Errno: 1045
                Last_IO_Error: error connecting to master 'replication@192.168.3                                                                                        1.201:3306' - retry-time: 10  retries: 86400  message: Access denied for user 'r                                                                                        eplication'@'192.168.31.200' (using password: YES)
               Last_SQL_Errno: 0
               Last_SQL_Error:
  Replicate_Ignore_Server_Ids:
             Master_Server_Id: 1
1 row in set (0.00 sec)
```

```
MariaDB [mysql]> START SLAVE;
Query OK, 0 rows affected (0.00 sec)

MariaDB [mysql]> SHOW SLAVE STATUS\G;
*************************** 1. row ***************************
               Slave_IO_State: Waiting to reconnect after a failed master event read
                  Master_Host: 192.168.31.200
                  Master_User: replication
                  Master_Port: 3306
                Connect_Retry: 10
              Master_Log_File: mysql-bin.000007
          Read_Master_Log_Pos: 847
               Relay_Log_File: mysql-bin.000001
                Relay_Log_Pos: 4
        Relay_Master_Log_File: mysql-bin.000007
             Slave_IO_Running: Connecting
            Slave_SQL_Running: Yes
              Replicate_Do_DB:
          Replicate_Ignore_DB:
           Replicate_Do_Table:
       Replicate_Ignore_Table:
      Replicate_Wild_Do_Table:
  Replicate_Wild_Ignore_Table:
                   Last_Errno: 0
                   Last_Error:
                 Skip_Counter: 0
          Exec_Master_Log_Pos: 847
              Relay_Log_Space: 245
              Until_Condition: None
               Until_Log_File:
                Until_Log_Pos: 0
           Master_SSL_Allowed: No
           Master_SSL_CA_File:
           Master_SSL_CA_Path:
              Master_SSL_Cert:
            Master_SSL_Cipher:
               Master_SSL_Key:
        Seconds_Behind_Master: NULL
Master_SSL_Verify_Server_Cert: No
                Last_IO_Errno: 0
                Last_IO_Error:
               Last_SQL_Errno: 0
               Last_SQL_Error:
  Replicate_Ignore_Server_Ids:
             Master_Server_Id: 1
1 row in set (0.00 sec)
```


参考：

https://mariadb.com/kb/en/replication-cluster-multi-master

http://www.it165.net/database/html/201404/6073.html

http://www.it165.net/database/html/201404/6074.html

http://www.it165.net/database/html/201312/5100.html
