---
layout: post
title: "top"
---

## top

```bash
$ top
top - 02:06:55 up 4 days,  6:32,  4 users,  load average: 0.00, 0.01, 0.05
Tasks: 120 total,   2 running, 118 sleeping,   0 stopped,   0 zombie
%Cpu(s):  0.0 us,  0.3 sy,  0.0 ni, 99.7 id,  0.0 wa,  0.0 hi,  0.0 si,  0.0 st
KiB Mem :   999920 total,   110224 free,   420940 used,   468756 buff/cache
KiB Swap:  2097148 total,  2072824 free,    24324 used.   317452 avail Mem

   PID USER      PR  NI    VIRT    RES    SHR S %CPU %MEM     TIME+ COMMAND
 20408 root      20   0  305080   5160   3796 S  0.7  0.5   4:52.11 vmtoolsd
 73034 root      20   0  157704   2212   1552 S  0.7  0.2   0:05.82 top
 73239 root      20   0  157708   2156   1500 R  0.3  0.2   0:00.02 top
     1 root      20   0  144336   5540   3040 S  0.0  0.6   0:08.17 systemd
     2 root      20   0       0      0      0 S  0.0  0.0   0:00.10 kthreadd
     3 root      20   0       0      0      0 S  0.0  0.0   0:05.03 ksoftirqd/0
     7 root      rt   0       0      0      0 S  0.0  0.0   0:00.00 migration/0
     8 root      20   0       0      0      0 S  0.0  0.0   0:00.00 rcu_bh
     9 root      20   0       0      0      0 R  0.0  0.0   2:27.79 rcu_sched
    10 root      rt   0       0      0      0 S  0.0  0.0   0:02.52 watchdog/0
    12 root      20   0       0      0      0 S  0.0  0.0   0:00.00 kdevtmpfs
    13 root       0 -20       0      0      0 S  0.0  0.0   0:00.00 netns
    14 root      20   0       0      0      0 S  0.0  0.0   0:00.16 khungtaskd
    15 root       0 -20       0      0      0 S  0.0  0.0   0:01.26 writeback
    16 root       0 -20       0      0      0 S  0.0  0.0   0:00.00 kintegrityd
    17 root       0 -20       0      0      0 S  0.0  0.0   0:00.00 bioset
    18 root       0 -20       0      0      0 S  0.0  0.0   0:00.00 kblockd
    19 root       0 -20       0      0      0 S  0.0  0.0   0:00.00 md
    25 root      20   0       0      0      0 S  0.0  0.0   0:04.35 kswapd0
    26 root      25   5       0      0      0 S  0.0  0.0   0:00.00 ksmd
    27 root      39  19       0      0      0 S  0.0  0.0   0:02.24 khugepaged
    28 root      20   0       0      0      0 S  0.0  0.0   0:00.00 fsnotify_mark
    29 root       0 -20       0      0      0 S  0.0  0.0   0:00.00 crypto
    37 root       0 -20       0      0      0 S  0.0  0.0   0:00.00 kthrotld
    39 root       0 -20       0      0      0 S  0.0  0.0   0:00.00 kmpath_rda
```

```
top - 02:06:55 up 4 days,  6:32,  4 users,  load average: 0.00, 0.01, 0.05
服务器概况: 当前程序名称 - 当前系统时间 系统已运行时间,当前用户登录数量,cpu的平均负载:1,5,15分钟内

Tasks: 120 total,   2 running, 118 sleeping,   0 stopped,   0 zombie
进程信息:进程总数,正在运行进程数,睡眠进程数,停止进程数,僵尸进程数

%Cpu(s):  0.0 us,  0.3 sy,  0.0 ni, 99.7 id,  0.0 wa,  0.0 hi,  0.0 si,  0.0 st
CPU信息：用户空间CPU占比,内核空间CPU占比,用户进程空间改过优先级的进程CPU占比,空闲CPU占比,待输入输出CPU占比,硬中断（Hardware IRQ）CPU占比,软中断（Software Interrupts）CPU占比,(Steal time) 是当 hypervisor 服务另一个虚拟处理器的时候，虚拟 CPU 等待实际 CPU 的时间占比

KiB Mem :   999920 total,   110224 free,   420940 used,   468756 buff/cache
内存信息:物理内存总量,空闲总量,使用中总量,缓存的内存量

KiB Swap:  2097148 total,  2072824 free,    24324 used.   317452 avail Mem
swap交换分区信息:交换区总量,空闲总量,使用中总量,缓存的内存量

进程（任务）状态监控
PID:进程id
USER:进程所有者
PR:进程优先级
NI:nice值。负值表示高优先级，正值表示低优先级
VIRT:进程使用的虚拟内存总量，单位kb。VIRT=SWAP+RES
RES:进程使用的、未被换出的物理内存大小，单位kb。RES=CODE+DATA
SHR:共享内存大小，单位kb
S:进程状态。D=不可中断的睡眠状态 R=运行 S=睡眠 T=跟踪/停止 Z=僵尸进程
%CPU:上次更新到现在的CPU时间占用百分比
%MEM:进程使用的物理内存百分比
TIME+:进程使用的CPU时间总计，单位1/100秒
COMMAND:进程名称（命令名/命令行）
```



load average 在 wikipedia 中的解释是 the system load is a measure of the amount of work that a computer system is doing 也就是对当前 CPU 工作量的度量，具体来说也就是指运行队列的平均长度，也就是等待CPU的平均进程数相关的一个计算值。

我们该如何看待这个load average 数据呢？

假设我们的系统是单CPU单内核的，把它比喻成是一条单向的桥，把CPU任务比作汽车。

- load = 0 的时候意味着这个桥上并没有车，cpu 没有任何任务；
- load < 1 的时候意味着桥上的车并不多，一切都还是很流畅的，cpu 的任务并不多，资源还很充足；
- load = 1 的时候就意味着桥已经被车给沾满了，没有一点空隙，cpu 的已经在全力工作了，所有的资源都被用完了，当然还好，这还在能力范围之内，只是有点慢而已；
- load > 1 的时候就意味着不仅仅是桥上已经被车占满了，就连桥外都被占满了，cpu 已经在全力的工作了，系统资源的用完了，但是还是有大量的进程在请求，在等待。若是这个值大于２，大于３，超过 CPU 工作能力的 2，３。而若是这个值 > 5 说明系统已经在超负荷运作了。

这是单个 CPU 单核的情况，而实际生活中我们需要将得到的这个值除以我们的核数来看。我们可以通过一下的命令来查看 CPU 的个数与核心数

```bash
# 总核数 = 物理CPU个数 X 每颗物理CPU的核数
# 总逻辑CPU数 = 物理CPU个数 X 每颗物理CPU的核数 X 超线程数

# 查看物理CPU个数
cat /proc/cpuinfo| grep "physical id"| sort| uniq| wc -l

# 查看每个物理CPU中core的个数(即核数)
cat /proc/cpuinfo| grep "cpu cores"| uniq

# 查看逻辑CPU的个数
cat /proc/cpuinfo| grep "processor"| wc -l
```
通过上面的指数我们可以得知 load 的临界值为 1 ，但是在实际生活中，比较有经验的运维或者系统管理员会将临界值定为0.7。这里的指数都是除以核心数以后的值，不要混淆了

若是 load < 0.7 并不会去关注他；
若是 0.7< load < 1 的时候我们就需要稍微关注一下了，虽然还可以应付但是这个值已经离临界不远了；
若是 load = 1 的时候我们就需要警惕了，因为这个时候已经没有更多的资源的了，已经在全力以赴了；
若是 load > 5 的时候系统已经快不行了，这个时候你需要加班解决问题了
通常我们都会先看 15 分钟的值来看这个大体的趋势，然后再看 5 分钟的值对比来看是否有下降的趋势。

查看 busybox 的代码可以知道，数据是每 5 秒钟就检查一次活跃的进程数，然后计算出该值，然后 load 从/proc/loadavg 中读取的。而这个 load 的值是如何计算的呢，这是 load 的计算的源码





我们回归正题，来看 top 的第二行数据，基本上第二行是进程的一个情况统计


PU 利用率，是对一个时间段内 CPU 使用状况的统计，通过这个指标可以看出在某一个时间段内 CPU 被占用的情况，Load Average 是 CPU 的 Load，它所包含的信息不是 CPU 的使用率状况，而是在一段时间内 CPU 正在处理以及等待 CPU 处理的进程数情况统计信息，这两个指标并不一样。

来看 top 的第四行数据，这一行基本上是内存的一个使用情况的统计了

内容	解释


注意

系统的中可用的物理内存最大值并不是 free 这个单一的值，而是 free + buffers + swap 中的 cached 的和
来看 top 的第五行数据，这一行基本上是交换区的一个使用情况的统计了

注意

NICE 值叫做静态优先级，是用户空间的一个优先级值，其取值范围是-20至19。这个值越小，表示进程”优先级”越高，而值越大“优先级”越低。nice值中的 -20 到 19，中 -20 优先级最高， 0 是默认的值，而 19 优先级最低

PR 值表示 Priority 值叫动态优先级，是进程在内核中实际的优先级值，进程优先级的取值范围是通过一个宏定义的，这个宏的名称是MAX_PRIO，它的值为140。Linux实际上实现了140个优先级范围，取值范围是从0-139，这个值越小，优先级越高。而这其中的 0 - 99 是实时的值，而 100 - 139 是给用户的。

其中 PR 中的 100 to 139 值部分有这么一个对应 PR = 20 + (-20 to +19)，这里的 -20 to +19 便是nice值，所以说两个虽然都是优先级，而且有千丝万缕的关系，但是他们的值，他们的作用范围并不相同

VIRT 任务所使用的虚拟内存的总数，其中包含所有的代码，数据，共享库和被换出 swap空间的页面等所占据空间的总数
在上文我们曾经说过 top 是一个前台程序，所以是一个可以交互的

常用交互命令	解释
q	退出程序
I	切换显示平均负载和启动时间的信息
P	根据CPU使用百分比大小进行排序
M	根据驻留内存大小进行排序
i	忽略闲置和僵死的进程，这是一个开关式命令
k	终止一个进程，系统提示输入 PID 及发送的信号值。一般终止进程用15信号，不能正常结束则使用9信号。安全模式下该命令被屏蔽。
好好的利用 top 能够很有效的帮助我们观察到系统的瓶颈所在，或者是系统的问题所在

(linux经常管理与查看指令：ps、pstree、top、kill、jobs和nice)[http://blog.csdn.net/watson2016/article/details/71189142]
