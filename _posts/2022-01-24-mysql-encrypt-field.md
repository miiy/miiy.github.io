---
layout: post
title: "mysql 数据库字段加密"
date: 2024-01-24
---

```sql
CREATE DATABASE `encrypt_test` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
CREATE TABLE `users` (
    `id` bigint unsigned NOT NULL AUTO_INCREMENT,
    `mobile` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
    `mobile_sha` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
    PRIMARY KEY (`id`),
    KEY `users_mobile_no_sha_index` (`mobile_sha`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
# 添加记录
INSERT INTO `users` (`mobile`, `mobile_sha`)VALUES (HEX(AES_ENCRYPT('13000000001', 'key')), SHA1('13000000001'));

SELECT * FROM `users`;
# +----+----------------------------------+------------------------------------------+
# | id | mobile                           | mobile_sha                               |
# +----+----------------------------------+------------------------------------------+
# |  1 | 2DCBE4B20147CC8A4EFA31C855F7136A | 13f4a6ec523ae680ed4d480a7daabf221685c369 |
# +----+----------------------------------+------------------------------------------+
# 精确查找
SELECT * FROM `users` WHERE `mobile_sha` = SHA1('13000000001');
# 模糊查找
SELECT * FROM `users` WHERE AES_DECRYPT(UNHEX(`mobile`), 'key') LIKE '13000%';
# 查询记录
SELECT CAST(AES_DECRYPT(UNHEX(`mobile`), 'key') AS CHAR CHARSET utf8) AS mobile, `mobile_sha` FROM `users`;
```

## 应用

1、可以把key分发到各个服务中，由各个服务对需要的数据解密

2、做个加密服务：提供单条和批量数据加密、解密，敏感数据加星

<https://www.jianshu.com/p/a06fce9c2e5a>

## 注意

MySQL8 block_encryption_mode 默认：aes-128-ecb

<https://dev.mysql.com/doc/refman/8.0/en/server-system-variables.html>
