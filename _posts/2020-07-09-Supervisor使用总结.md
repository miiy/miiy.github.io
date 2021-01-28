---
layout: post
title: "Supervisor 使用总结"
date: 2020-07-09
tags: other supervisor
---

Supervisor: A Process Control System

## Bookmarks

Supervisor: http://supervisord.org

Supervisor使用详解：https://www.jianshu.com/p/0b9054b33db3

## Installing

```bash
[root@xx /]# yum install -y supervisor
```

## 配置文件

```bash
[root@xx /]# cat /etc/supervisord.conf
```

```conf
; Sample supervisor config file.

[unix_http_server]
file=/var/run/supervisor/supervisor.sock   ; (the path to the socket file)
;chmod=0700                 ; sockef file mode (default 0700)
;chown=nobody:nogroup       ; socket file uid:gid owner
;username=user              ; (default is no username (open server))
;password=123               ; (default is no password (open server))

;[inet_http_server]         ; inet (TCP) server disabled by default
;port=127.0.0.1:9001        ; (ip_address:port specifier, *:port for all iface)
;username=user              ; (default is no username (open server))
;password=123               ; (default is no password (open server))

[supervisord]
logfile=/var/log/supervisor/supervisord.log  ; (main log file;default $CWD/supervisord.log)
logfile_maxbytes=50MB       ; (max main logfile bytes b4 rotation;default 50MB)
logfile_backups=10          ; (num of main logfile rotation backups;default 10)
loglevel=info               ; (log level;default info; others: debug,warn,trace)
pidfile=/var/run/supervisord.pid ; (supervisord pidfile;default supervisord.pid)
nodaemon=false              ; (start in foreground if true;default false)
minfds=1024                 ; (min. avail startup file descriptors;default 1024)
minprocs=200                ; (min. avail process descriptors;default 200)
;umask=022                  ; (process file creation umask;default 022)
;user=chrism                 ; (default is current user, required if root)
;identifier=supervisor       ; (supervisord identifier, default is 'supervisor')
;directory=/tmp              ; (default is not to cd during start)
;nocleanup=true              ; (don't clean up tempfiles at start;default false)
;childlogdir=/tmp            ; ('AUTO' child log dir, default $TEMP)
;environment=KEY=value       ; (key value pairs to add to environment)
;strip_ansi=false            ; (strip ansi escape codes in logs; def. false)

; the below section must remain in the config file for RPC
; (supervisorctl/web interface) to work, additional interfaces may be
; added by defining them in separate rpcinterface: sections
[rpcinterface:supervisor]
supervisor.rpcinterface_factory = supervisor.rpcinterface:make_main_rpcinterface

[supervisorctl]
serverurl=unix:///var/run/supervisor/supervisor.sock ; use a unix:// URL  for a unix socket
;serverurl=http://127.0.0.1:9001 ; use an http:// url to specify an inet socket
;username=chris              ; should be same as http_username if set
;password=123                ; should be same as http_password if set
;prompt=mysupervisor         ; cmd line prompt (default "supervisor")
;history_file=~/.sc_history  ; use readline history if available

; The below sample program section shows all possible program subsection values,
; create one or more 'real' program: sections to be able to control them under
; supervisor.

;[program:theprogramname]
;command=/bin/cat              ; the program (relative uses PATH, can take args)
;process_name=%(program_name)s ; process_name expr (default %(program_name)s)
;numprocs=1                    ; number of processes copies to start (def 1)
;directory=/tmp                ; directory to cwd to before exec (def no cwd)
;umask=022                     ; umask for process (default None)
;priority=999                  ; the relative start priority (default 999)
;autostart=true                ; start at supervisord start (default: true)
;autorestart=true              ; retstart at unexpected quit (default: true)
;startsecs=10                  ; number of secs prog must stay running (def. 1)
;startretries=3                ; max # of serial start failures (default 3)
;exitcodes=0,2                 ; 'expected' exit codes for process (default 0,2)
;stopsignal=QUIT               ; signal used to kill process (default TERM)
;stopwaitsecs=10               ; max num secs to wait b4 SIGKILL (default 10)
;user=chrism                   ; setuid to this UNIX account to run the program
;redirect_stderr=true          ; redirect proc stderr to stdout (default false)
;stdout_logfile=/a/path        ; stdout log path, NONE for none; default AUTO
;stdout_logfile_maxbytes=1MB   ; max # logfile bytes b4 rotation (default 50MB)
;stdout_logfile_backups=10     ; # of stdout logfile backups (default 10)
;stdout_capture_maxbytes=1MB   ; number of bytes in 'capturemode' (default 0)
;stdout_events_enabled=false   ; emit events on stdout writes (default false)
;stderr_logfile=/a/path        ; stderr log path, NONE for none; default AUTO
;stderr_logfile_maxbytes=1MB   ; max # logfile bytes b4 rotation (default 50MB)
;stderr_logfile_backups=10     ; # of stderr logfile backups (default 10)
;stderr_capture_maxbytes=1MB   ; number of bytes in 'capturemode' (default 0)
;stderr_events_enabled=false   ; emit events on stderr writes (default false)
;environment=A=1,B=2           ; process environment additions (def no adds)
;serverurl=AUTO                ; override serverurl computation (childutils)

; The below sample eventlistener section shows all possible
; eventlistener subsection values, create one or more 'real'
; eventlistener: sections to be able to handle event notifications
; sent by supervisor.

;[eventlistener:theeventlistenername]
;command=/bin/eventlistener    ; the program (relative uses PATH, can take args)
;process_name=%(program_name)s ; process_name expr (default %(program_name)s)
;numprocs=1                    ; number of processes copies to start (def 1)
;events=EVENT                  ; event notif. types to subscribe to (req'd)
;buffer_size=10                ; event buffer queue size (default 10)
;directory=/tmp                ; directory to cwd to before exec (def no cwd)
;umask=022                     ; umask for process (default None)
;priority=-1                   ; the relative start priority (default -1)
;autostart=true                ; start at supervisord start (default: true)
;autorestart=unexpected        ; restart at unexpected quit (default: unexpected)
;startsecs=10                  ; number of secs prog must stay running (def. 1)
;startretries=3                ; max # of serial start failures (default 3)
;exitcodes=0,2                 ; 'expected' exit codes for process (default 0,2)
;stopsignal=QUIT               ; signal used to kill process (default TERM)
;stopwaitsecs=10               ; max num secs to wait b4 SIGKILL (default 10)
;user=chrism                   ; setuid to this UNIX account to run the program
;redirect_stderr=true          ; redirect proc stderr to stdout (default false)
;stdout_logfile=/a/path        ; stdout log path, NONE for none; default AUTO
;stdout_logfile_maxbytes=1MB   ; max # logfile bytes b4 rotation (default 50MB)
;stdout_logfile_backups=10     ; # of stdout logfile backups (default 10)
;stdout_events_enabled=false   ; emit events on stdout writes (default false)
;stderr_logfile=/a/path        ; stderr log path, NONE for none; default AUTO
;stderr_logfile_maxbytes=1MB   ; max # logfile bytes b4 rotation (default 50MB)
;stderr_logfile_backups        ; # of stderr logfile backups (default 10)
;stderr_events_enabled=false   ; emit events on stderr writes (default false)
;environment=A=1,B=2           ; process environment additions
;serverurl=AUTO                ; override serverurl computation (childutils)

; The below sample group section shows all possible group values,
; create one or more 'real' group: sections to create "heterogeneous"
; process groups.

;[group:thegroupname]
;programs=progname1,progname2  ; each refers to 'x' in [program:x] definitions
;priority=999                  ; the relative start priority (default 999)

; The [include] section can just contain the "files" setting.  This
; setting can list multiple files (separated by whitespace or
; newlines).  It can also contain wildcards.  The filenames are
; interpreted as relative to this file.  Included files *cannot*
; include files themselves.

[include]
files = supervisord.d/*.ini

```

## 子进程配置示例

```bash
[root@xx /]# cat /etc/supervisord.d/test-project.ini 
```

```ini
[program:test-project-worker]
process_name=%(program_name)s_%(process_num)02d
command=/usr/local/php/bin/php -r "sleep(5);echo 1;"
autostart=true
autorestart=true
numprocs=2
redirect_stderr=true
stdout_logfile=/var/log/supervisor/test-project-worker.log
```

## 命令说明

```bash
[root@xx /]# systemctl stop supervisord
[root@xx /]# systemctl start supervisord
[root@xx /]# systemctl restart supervisord
[root@xx /]# systemctl status supervisord
# 开机启动
[root@xx /]# systemctl enable supervisord
Created symlink from /etc/systemd/system/multi-user.target.wants/supervisord.service to /usr/lib/systemd/system/supervisord.service.
[root@izm5e7a5np43amsvuiddlrz ~]# systemctl is-enabled supervisord
enabled
```

## 启用 Web Server

去掉配置文件中的注释

```conf
[inet_http_server]         ; inet (TCP) server disabled by default
port=127.0.0.1:9001        ; (ip_address:port specifier, *:port for all iface)
username=user              ; (default is no username (open server))
password=123               ; (default is no password (open server))
```

```bash
[root@xx /]# curl --basic -u user:123 http://localhost:9001
```

## 更新配置文件

```bash
[root@xx /]# vim /etc/supervisord.d/test-project.ini

[root@xx /]# supervisorctl update test-project-worker
test-project-worker: stopped
test-project-worker: updated process group
# 更新所有
[root@xx /]# supervisorctl update
test-project-worker: stopped
test-project-worker: updated process group

[root@xx /]# supervisorctl status
test-project-worker:test-project-worker_00   STARTING  
test-project-worker:test-project-worker_01   STARTING 
```


## Laravel 配置示例

```bash
[root@xx /]# cat /etc/supervisord.d/test-project.ini
```

```ini
[program:test-project-worker]
process_name=%(program_name)s_%(process_num)02d
command=/usr/bin/php /data/www/test-project/api/artisan queue:work --sleep=3 --tries=3
autostart=true
autorestart=true
numprocs=2
redirect_stderr=true
stdout_logfile=/var/log/supervisor/test-project-worker.log

[program:test-project-log-worker]
process_name=%(program_name)s_%(process_num)02d
command=/usr/bin/php /data/www/test-project/api/artisan queue:work --queue=log --sleep=3 --tries=3
autostart=true
autorestart=true
numprocs=2
redirect_stderr=true
stdout_logfile=/var/log/supervisor/test-project-log-worker.log

[program:test-project-notification-worker]
process_name=%(program_name)s_%(process_num)02d
command=/usr/bin/php /data/www/test-project/api/artisan queue:work --queue=notification --sleep=3 --tries=3
autostart=true
autorestart=true
numprocs=2
redirect_stderr=true
stdout_logfile=/var/log/supervisor/test-project-notification.log
```

## supervisorctl 命令参数

```bash
[root@xx /]# supervisorctl --help
supervisorctl -- control applications run by supervisord from the cmd line.

Usage: /usr/bin/supervisorctl [options] [action [arguments]]

Options:
-c/--configuration FILENAME -- configuration file path (searches if not given)
-h/--help -- print usage message and exit
-i/--interactive -- start an interactive shell after executing commands
-s/--serverurl URL -- URL on which supervisord server is listening
     (default "http://localhost:9001").
-u/--username USERNAME -- username to use for authentication with server
-p/--password PASSWORD -- password to use for authentication with server
-r/--history-file -- keep a readline history (if readline is available)

action [arguments] -- see below

Actions are commands like "tail" or "stop".  If -i is specified or no action is
specified on the command line, a "shell" interpreting actions typed
interactively is started.  Use the action "help" to find out about available
actions.

[root@xx /]# supervisorctl help 

default commands (type help <topic>):
=====================================
add    exit      open  reload  restart   start   tail   
avail  fg        pid   remove  shutdown  status  update 
clear  maintail  quit  reread  signal    stop    version

[root@xx /]# supervisorctl help status
status <name>		Get status for a single process
status <gname>:*	Get status for all processes in a group
status <name> <name>	Get status for multiple named processes
status			Get all process status info
[root@xx /]# supervisorctl help update
update			Reload config and add/remove as necessary, and will restart affected programs
update all		Reload config and add/remove as necessary, and will restart affected programs
update <gname> [...]	Update specific groups
[root@xx /]# supervisorctl help reload
reload 		Restart the remote supervisord.
[root@xx /]# supervisorctl help restart
restart <name>		Restart a process
restart <gname>:*	Restart all processes in a group
restart <name> <name>	Restart multiple processes or groups
restart all		Restart all processes
Note: restart does not reread config files. For that, see reread and update.
[root@xx /]# supervisorctl help stop
stop <name>		Stop a process
stop <gname>:*		Stop all processes in a group
stop <name> <name>	Stop multiple processes or groups
stop all		Stop all processes
```
