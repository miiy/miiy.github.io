---
layout: post
title: "postgresql"
date: 2025-10-23
tags: linux
---

## 连接

```bash
psql -h 127.0.0.1 -p 5432 -U admin -d postgres
```

## SQL

重命名表

```sql
ALTER TABLE users RENAME TO users_old;
```

复制表

```sql
CREATE TABLE backup_table AS
SELECT * FROM source_table;
```

当前时间

```sql
select now();
2025-10-23 14:05:55.051471

SELECT current_timestamp;
2025-10-23 14:05:55.051471
```

当前日期

```sql
SELECT current_date;
2025-10-23
```

当前时间

```sql
SELECT current_time;
14:07:52.165596
```

时间戳

```sql
SELECT EXTRACT(EPOCH FROM NOW())::BIGINT;
1761199876
-- 更安全的使用 FLOOR 去掉小数部分
SELECT FLOOR(EXTRACT(EPOCH FROM NOW()))::BIGINT;
1761199876
```