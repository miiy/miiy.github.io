---
layout: post
title: "Docker CMD ENTRYPOINT 区别"
---

## CMD

https://docs.docker.com/engine/reference/builder/#cmd

The `CMD` instruction has three forms:

CMD 有三种指令形式：

* `CMD ["executable","param1","param2"]` (exec form, this is the preferred form)
* `CMD ["param1","param2"]` (as default parameters to ENTRYPOINT)
* `CMD command param1 param2` (shell form)

There can only be one `CMD` instruction in a `Dockerfile`. If you list more than one CMD then only the last CMD will take effect.

一个 Dockerfile 中只能有一条 CMD 指令。如果您列出多个 CMD，则只有最后一个 CMD 会生效。

The main purpose of a `CMD` is to provide defaults for an executing container. These defaults can include an executable, or they can omit the executable, in which case you must specify an `ENTRYPOINT` instruction as well.

CMD 的主要目的是为正在执行的容器提供默认值。这些默认值可以包含可执行文件，也可以省略可执行文件，在这种情况下，您还必须指定 ENTRYPOINT 指令。

If CMD is used to provide default arguments for the ENTRYPOINT instruction, both the CMD and ENTRYPOINT instructions should be specified with the JSON array format.

如果使用 CMD 为 ENTRYPOINT 指令提供默认参数，则 CMD 和 ENTRYPOINT 指令都应使用 JSON 数组格式指定。

> Note: The exec form is parsed as a JSON array, which means that you must use double-quotes (“) around words not single-quotes (‘).

> 注意：exec 形式被解析为 JSON 数组，这意味着您必须在单词周围使用双引号 (“) 而不是单引号 (‘)。

Unlike the shell form, the exec form does not invoke a command shell. This means that normal shell processing does not happen. For example, `CMD [ "echo", "$HOME" ]` will not do variable substitution on `$HOME`. If you want shell processing then either use the shell form or execute a shell directly, for example: `CMD [ "sh", "-c", "echo $HOME" ]`. When using the exec form and executing a shell directly, as in the case for the shell form, it is the shell that is doing the environment variable expansion, not docker.

与 shell 形式不同，exec 形式不调用命令 shell。这意味着正常的 shell 处理不会发生。例如，`CMD [ "echo", "$HOME" ]` 不会对 `$HOME` 进行变量替换。如果您想要 shell 处理，那么要么使用 shell 形式，要么直接执行 shell，例如：`CMD [ "sh", "-c", "echo $HOME" ]`。当使用exec形式直接执行shell时，如shell形式，是shell在做环境变量扩展，而不是docker。

When used in the shell or exec formats, the `CMD` instruction sets the command to be executed when running the image.

当以 shell 或 exec 格式使用时，`CMD` 指令设置运行镜像时要执行的命令。

If you use the shell form of the `CMD`, then the `<command>` will execute in `/bin/sh -c`:

如果您使用 `CMD` 的 shell 形式，则 `<command>` 将在 `/bin/sh -c` 中执行：

```text
FROM ubuntu
CMD echo "This is a test." | wc -
```

If you want to run your `<command>` without a shell then you must express the command as a JSON array and give the full path to the executable. This array form is the preferred format of `CMD`. Any additional parameters must be individually expressed as strings in the array:

如果你想在没有 shell 的情况下运行你的 `<command>` 那么你必须将命令表达为一个 JSON 数组并给出可执行文件的完整路径。这种数组形式是 `CMD` 的首选格式。任何附加参数都必须在数组中单独表示为字符串：

```text
FROM ubuntu
CMD ["/usr/bin/wc","--help"]
```

If you would like your container to run the same executable every time, then you should consider using `ENTRYPOINT` in combination with `CMD`. See ENTRYPOINT.

如果您希望您的容器每次都运行相同的可执行文件，那么您应该考虑结合使用 `ENTRYPOINT` 和 `CMD`。请参见 ENTRYPOINT。

If the user specifies arguments to `docker run` then they will override the default specified in `CMD`.

如果用户为 `docker run` 指定参数，那么它们将覆盖 `CMD` 中指定的默认值。

> Note: Do not confuse `RUN` with `CMD`. `RUN` actually runs a command and commits the result; `CMD` does not execute anything at build time, but specifies the intended command for the image.

> 注意：不要混淆 `RUN` 和 `CMD`。`RUN` 实际上运行一个命令并提交结果； `CMD` 在构建时不执行任何操作，而是指定镜像的预期命令。

## ENTRYPOINT

https://docs.docker.com/engine/reference/builder/#entrypoint

ENTRYPOINT has two forms:

ENTRYPOINT 有两种形式：

The exec form, which is the preferred form:

exec 形式，这是首选形式：

```
ENTRYPOINT ["executable", "param1", "param2"]
```

The shell form:

shell 形式：

```
ENTRYPOINT command param1 param2
```

An `ENTRYPOINT` allows you to configure a container that will run as an executable.

`ENTRYPOINT` 允许您配置将作为可执行文件运行的容器。

For example, the following starts nginx with its default content, listening on port 80:

例如，下面以默认内容启动 nginx，监听端口 80：

```bash
docker run -i -t --rm -p 80:80 nginx
```

Command line arguments to `docker run <image>` will be appended after all elements in an exec form `ENTRYPOINT`, and will override all elements specified using `CMD`. This allows arguments to be passed to the entry point, i.e., `docker run <image> -d` will pass the `-d` argument to the entry point. You can override the `ENTRYPOINT` instruction using the `docker run --entrypoint flag`.

`docker run <image>` 的命令行参数将附加在 exec 形式 `ENTRYPOINT` 中的所有元素之后，并将覆盖使用 `CMD` 指定的所有元素。 这允许将参数传递到 entry point，即 `docker run <image> -d` 会将 `-d` 参数传递到 entry point。 您可以使用 `docker run --entrypoint flag` 覆盖 `ENTRYPOINT` 指令。

The shell form prevents any `CMD` or `run` command line arguments from being used, but has the disadvantage that your `ENTRYPOINT` will be started as a subcommand of `/bin/sh -c`, which does not pass signals. This means that the executable will not be the container’s `PID 1` - and will not receive Unix signals - so your executable will not receive a `SIGTERM` from `docker stop <container>`.

shell 形式阻止使用任何 `CMD` 或 `run` 命令行参数，但缺点是您的 `ENTRYPOINT` 将作为 `/bin/sh -c` 的子命令启动，它不传递信号。这意味着可执行文件不会是容器的 `PID 1` - 也不会接收 Unix 信号 - 所以你的可执行文件不会从 `docker stop <container>` 接收到 `SIGTERM` 。

Only the last `ENTRYPOINT` instruction in the `Dockerfile` will have an effect.

只有 `Dockerfile` 中的最后一条 `ENTRYPOINT` 指令才会生效。

### Exec form ENTRYPOINT example

You can use the exec form of `ENTRYPOINT` to set fairly stable default commands and arguments and then use either form of `CMD` to set additional defaults that are more likely to be changed.

```dockerfile
FROM ubuntu
ENTRYPOINT ["top", "-b"]
CMD ["-c"]
```

When you run the container, you can see that `top` is the only process:

```bash
$ docker run -it --rm --name test  top -H

top - 08:25:00 up  7:27,  0 users,  load average: 0.00, 0.01, 0.05
Threads:   1 total,   1 running,   0 sleeping,   0 stopped,   0 zombie
%Cpu(s):  0.1 us,  0.1 sy,  0.0 ni, 99.7 id,  0.0 wa,  0.0 hi,  0.0 si,  0.0 st
KiB Mem:   2056668 total,  1616832 used,   439836 free,    99352 buffers
KiB Swap:  1441840 total,        0 used,  1441840 free.  1324440 cached Mem

  PID USER      PR  NI    VIRT    RES    SHR S %CPU %MEM     TIME+ COMMAND
    1 root      20   0   19744   2336   2080 R  0.0  0.1   0:00.04 top
```

To examine the result further, you can use `docker exec`:

```bash
$ docker exec -it test ps aux

USER       PID %CPU %MEM    VSZ   RSS TTY      STAT START   TIME COMMAND
root         1  2.6  0.1  19752  2352 ?        Ss+  08:24   0:00 top -b -H
root         7  0.0  0.1  15572  2164 ?        R+   08:25   0:00 ps aux
```

And you can gracefully request `top` to shut down using `docker stop test`.

The following `Dockerfile` shows using the `ENTRYPOINT` to run Apache in the foreground (i.e., as `PID 1`):

```dockerfile
FROM debian:stable
RUN apt-get update && apt-get install -y --force-yes apache2
EXPOSE 80 443
VOLUME ["/var/www", "/var/log/apache2", "/etc/apache2"]
ENTRYPOINT ["/usr/sbin/apache2ctl", "-D", "FOREGROUND"]
```

If you need to write a starter script for a single executable, you can ensure that the final executable receives the Unix signals by using `exec` and `gosu` commands:

```bash
#!/usr/bin/env bash
set -e

if [ "$1" = 'postgres' ]; then
    chown -R postgres "$PGDATA"

    if [ -z "$(ls -A "$PGDATA")" ]; then
        gosu postgres initdb
    fi

    exec gosu postgres "$@"
fi

exec "$@"
```

Lastly, if you need to do some extra cleanup (or communicate with other containers) on shutdown, or are co-ordinating more than one executable, you may need to ensure that the `ENTRYPOINT` script receives the Unix signals, passes them on, and then does some more work:

```bash
#!/bin/sh
# Note: I've written this using sh so it works in the busybox container too

# USE the trap if you need to also do manual cleanup after the service is stopped,
#     or need to start multiple services in the one container
trap "echo TRAPed signal" HUP INT QUIT TERM

# start service in background here
/usr/sbin/apachectl start

echo "[hit enter key to exit] or run 'docker stop <container>'"
read

# stop service and clean up here
echo "stopping apache"
/usr/sbin/apachectl stop

echo "exited $0"
```

If you run this image with `docker run -it --rm -p 80:80 --name test apache`, you can then examine the container’s processes with `docker exec`, or `docker top`, and then ask the script to stop Apache:

```bash
$ docker exec -it test ps aux

USER       PID %CPU %MEM    VSZ   RSS TTY      STAT START   TIME COMMAND
root         1  0.1  0.0   4448   692 ?        Ss+  00:42   0:00 /bin/sh /run.sh 123 cmd cmd2
root        19  0.0  0.2  71304  4440 ?        Ss   00:42   0:00 /usr/sbin/apache2 -k start
www-data    20  0.2  0.2 360468  6004 ?        Sl   00:42   0:00 /usr/sbin/apache2 -k start
www-data    21  0.2  0.2 360468  6000 ?        Sl   00:42   0:00 /usr/sbin/apache2 -k start
root        81  0.0  0.1  15572  2140 ?        R+   00:44   0:00 ps aux

$ docker top test

PID                 USER                COMMAND
10035               root                {run.sh} /bin/sh /run.sh 123 cmd cmd2
10054               root                /usr/sbin/apache2 -k start
10055               33                  /usr/sbin/apache2 -k start
10056               33                  /usr/sbin/apache2 -k start

$ /usr/bin/time docker stop test

test
real	0m 0.27s
user	0m 0.03s
sys	0m 0.03s
```

> Note: You can override the `ENTRYPOINT` setting using `--entrypoint`, but this can only set the binary to exec (no `sh -c` will be used).

> Note: The exec form is parsed as a JSON array, which means that you must use double-quotes (“) around words not single-quotes (‘).

Unlike the shell form, the exec form does not invoke a command shell. This means that normal shell processing does not happen. For example, `ENTRYPOINT [ "echo", "$HOME" ]` will not do variable substitution on `$HOME`. If you want shell processing then either use the shell form or execute a shell directly, for example: `ENTRYPOINT [ "sh", "-c", "echo $HOME" ]`. When using the exec form and executing a shell directly, as in the case for the shell form, it is the shell that is doing the environment variable expansion, not docker.




### Shell form ENTRYPOINT example

You can specify a plain string for the ENTRYPOINT and it will execute in /bin/sh -c. This form will use shell processing to substitute shell environment variables, and will ignore any CMD or docker run command line arguments. To ensure that docker stop will signal any long running ENTRYPOINT executable correctly, you need to remember to start it with exec:

```bash
FROM ubuntu
ENTRYPOINT exec top -b
```

When you run this image, you’ll see the single PID 1 process:

```bash
$ docker run -it --rm --name test top
Mem: 1704520K used, 352148K free, 0K shrd, 0K buff, 140368121167873K cached
CPU:   5% usr   0% sys   0% nic  94% idle   0% io   0% irq   0% sirq
Load average: 0.08 0.03 0.05 2/98 6
  PID  PPID USER     STAT   VSZ %VSZ %CPU COMMAND
    1     0 root     R     3164   0%   0% top -b
```

Which will exit cleanly on docker stop:

```bash
$ /usr/bin/time docker stop test
test
real	0m 0.20s
user	0m 0.02s
sys	0m 0.04s
```

If you forget to add exec to the beginning of your ENTRYPOINT:

```bash
FROM ubuntu
ENTRYPOINT top -b
CMD --ignored-param1
```

You can then run it (giving it a name for the next step):

```bash

$ docker run -it --name test top --ignored-param2
Mem: 1704184K used, 352484K free, 0K shrd, 0K buff, 140621524238337K cached
CPU:   9% usr   2% sys   0% nic  88% idle   0% io   0% irq   0% sirq
Load average: 0.01 0.02 0.05 2/101 7
  PID  PPID USER     STAT   VSZ %VSZ %CPU COMMAND
    1     0 root     S     3168   0%   0% /bin/sh -c top -b cmd cmd2
    7     1 root     R     3164   0%   0% top -b
```

You can see from the output of top that the specified ENTRYPOINT is not PID 1.

If you then run docker stop test, the container will not exit cleanly - the stop command will be forced to send a SIGKILL after the timeout:

```bash
$ docker exec -it test ps aux
PID   USER     COMMAND
    1 root     /bin/sh -c top -b cmd cmd2
    7 root     top -b
    8 root     ps aux
$ /usr/bin/time docker stop test
test
real	0m 10.19s
user	0m 0.04s
sys	0m 0.03s
```

### Understand how CMD and ENTRYPOINT interact

### 了解CMD和ENTRYPOINT是如何交互的

Both CMD and ENTRYPOINT instructions define what command gets executed when running a container. There are few rules that describe their co-operation.

1. Dockerfile should specify at least one of CMD or ENTRYPOINT commands.

1. Dockerfile 应该指定至少一个 CMD 或 ENTRYPOINT 命令。

2. ENTRYPOINT should be defined when using the container as an executable.

2. 在将容器用作可执行文件时，应该定义 ENTRYPOINT

3. CMD should be used as a way of defining default arguments for an ENTRYPOINT command or for executing an ad-hoc command in a container.

4. CMD will be overridden when running the container with alternative arguments.

4. 当运行带有可选参数的容器时，CMD将被覆盖。

The table below shows what command is executed for different ENTRYPOINT / CMD combinations:


.|No ENTRYPOINT|ENTRYPOINT exec_entry p1_entry|ENTRYPOINT [“exec_entry”, “p1_entry”]
--|--|--
No CMD|error, not allowed|/bin/sh -c exec_entry p1_entry|exec_entry p1_entry
CMD [“exec_cmd”, “p1_cmd”]|exec_cmd p1_cmd|/bin/sh -c exec_entry p1_entry|exec_entry p1_entry exec_cmd p1_cmd
CMD [“p1_cmd”, “p2_cmd”]|p1_cmd p2_cmd|/bin/sh -c exec_entry p1_entry|exec_entry p1_entry p1_cmd p2_cmd
CMD exec_cmd p1_cmd|/bin/sh -c exec_cmd p1_cmd|/bin/sh -c exec_entry p1_entry|exec_entry p1_entry /bin/sh -c exec_cmd p1_cmd

> Note: If CMD is defined from the base image, setting ENTRYPOINT will reset CMD to an empty value. In this scenario, CMD must be defined in the current image to have a value.