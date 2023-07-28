---
layout: post
title: "kill"
---

```bash
$ kill --help
kill: kill [-s sigspec | -n signum | -sigspec] pid | jobspec ... or kill -l [sigspec]
Send a signal to a job.
 
Send the processes identified by PID or JOBSPEC the signal named by
SIGSPEC or SIGNUM.  If neither SIGSPEC nor SIGNUM is present, then
SIGTERM is assumed.
 
Options:
-s sig    SIG is a signal name
-n sig    SIG is a signal number
-l    list the signal names; if arguments follow `-l' they are
assumed to be signal numbers for which names should be listed
-L    synonym for -l
 
Kill is a shell builtin for two reasons: it allows job IDs to be used
instead of process IDs, and allows processes to be killed if the limit
on processes that you can create is reached.
 
Exit Status:
Returns success unless an invalid option is given or an error occurs.
```

向作业发送信号

向PID或JOBSPEC标识的进程发送SIGSPEC或SIGNUM指定的信号。

如果SIGSPEC和SIGNUM均不存在，则假定为 SIGTERM

```bash
$ kill -l
 1) SIGHUP	 2) SIGINT	 3) SIGQUIT	 4) SIGILL	 5) SIGTRAP
 6) SIGABRT	 7) SIGBUS	 8) SIGFPE	 9) SIGKILL	10) SIGUSR1
11) SIGSEGV	12) SIGUSR2	13) SIGPIPE	14) SIGALRM	15) SIGTERM
16) SIGSTKFLT	17) SIGCHLD	18) SIGCONT	19) SIGSTOP	20) SIGTSTP
21) SIGTTIN	22) SIGTTOU	23) SIGURG	24) SIGXCPU	25) SIGXFSZ
26) SIGVTALRM	27) SIGPROF	28) SIGWINCH	29) SIGIO	30) SIGPWR
31) SIGSYS	34) SIGRTMIN	35) SIGRTMIN+1	36) SIGRTMIN+2	37) SIGRTMIN+3
38) SIGRTMIN+4	39) SIGRTMIN+5	40) SIGRTMIN+6	41) SIGRTMIN+7	42) SIGRTMIN+8
43) SIGRTMIN+9	44) SIGRTMIN+10	45) SIGRTMIN+11	46) SIGRTMIN+12	47) SIGRTMIN+13
48) SIGRTMIN+14	49) SIGRTMIN+15	50) SIGRTMAX-14	51) SIGRTMAX-13	52) SIGRTMAX-12
53) SIGRTMAX-11	54) SIGRTMAX-10	55) SIGRTMAX-9	56) SIGRTMAX-8	57) SIGRTMAX-7
58) SIGRTMAX-6	59) SIGRTMAX-5	60) SIGRTMAX-4	61) SIGRTMAX-3	62) SIGRTMAX-2
63) SIGRTMAX-1	64) SIGRTMAX
```


```text
SIGKILL
kill -9  强制退出
信号不能被捕获或忽略，
 
SIGTERM
kill -15 正常停止
程序收到该信号后，可以自己决定如何处理
1、立即停止
2、释放资源后停止
3、忽略信号，继续执行程序
 
SIGINT
程序终止信号
通常由 Ctrl + C 发出
 
SIGSTOP
停止进程的执行
通常由 Ctrl + Z 发出
```
