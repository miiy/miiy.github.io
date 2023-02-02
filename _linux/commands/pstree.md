---
layout: post
title: "pstree"
---

Help

```bash
$ pstree --help
pstree: unrecognized option '--help'
Usage: pstree [ -a ] [ -c ] [ -h | -H PID ] [ -l ] [ -n ] [ -p ] [ -g ] [ -u ]
              [ -A | -G | -U ] [ PID | USER ]
       pstree -V
Display a tree of processes.
  -a, --arguments     show command line arguments
  -A, --ascii         use ASCII line drawing characters
  -c, --compact       don't compact identical subtrees
  -h, --highlight-all highlight current process and its ancestors
  -H PID,
  --highlight-pid=PID highlight this process and its ancestors
  -g, --show-pgids    show process group ids; implies -c
  -G, --vt100         use VT100 line drawing characters
  -l, --long          don't truncate long lines
  -n, --numeric-sort  sort output by PID
  -N type,
  --ns-sort=type      sort by namespace type (ipc, mnt, net, pid, user, uts)
  -p, --show-pids     show PIDs; implies -c
  -s, --show-parents  show parents of the selected process
  -S, --ns-changes    show namespace transitions
  -u, --uid-changes   show uid transitions
  -U, --unicode       use UTF-8 (Unicode) line drawing characters
  -V, --version       display version information
  -Z,
  --security-context   show SELinux security contexts
  PID    start at this PID; default is 1 (init)
  USER   show only trees rooted at processes of this user
```


Manual

```bash
$ man pstree
PSTREE(1)                                                User Commands                                                PSTREE(1)
NAME
       pstree - display a tree of processes
SYNOPSIS
       pstree [-a, --arguments] [-c, --compact] [-h, --highlight-all, -Hpid, --highlight-pid pid] [-g] --show-pgids]
       [-l, --long] [-n, --numeric-sort] [-N, --ns-sortns [-p, --show-pids] [-s, --show-parents] [-S, --ns-changes]
       [-u, --uid-changes] [-Z, --security-context] [-A, --ascii, -G, --vt100, -U, --unicode] [pid, user]
       pstree -V, --version
DESCRIPTION
       pstree  shows  running processes as a tree.  The tree is rooted at either pid or init if pid is omitted.  If a user name
       is specified, all process trees rooted at processes owned by that user are shown.
       pstree visually merges identical branches by putting them in square brackets and  prefixing  them  with  the  repetition
       count, e.g.
           init-+-getty
                |-getty
                |-getty
                `-getty
       becomes
           init---4*[getty]
       Child threads of a process are found under the parent process and are shown with the process name in curly braces, e.g.
           icecast2---13*[{icecast2}]
       If  pstree  is  called  as  pstree.x11  then it will prompt the user at the end of the line to press return and will not
       return until that has happened.  This is useful for when pstree is run in a xterminal.
       Certain kernel or mount parameters, such as the hidepid option for procfs, will hide information for some processes.  In
       these  situations  pstree  will  attempt  to  build the tree without this information, showing process names as question
       marks.
OPTIONS
       -a     Show command line arguments.  If the command line of a process is swapped out, that process is shown in parenthe‐
              ses.  -a implicitly disables compaction for processes but not threads.
       -A     Use ASCII characters to draw the tree.
       -c     Disable compaction of identical subtrees.  By default, subtrees are compacted whenever possible.
       -G     Use VT100 line drawing characters.
       -h     Highlight the current process and its ancestors.  This is a no-op if the terminal doesn't support highlighting or
              if neither the current process nor any of its ancestors are in the subtree being shown.
       -H     Like -h, but highlight the specified process instead.  Unlike with -h, pstree fails when using -H if highlighting
              is not available.
       -g     Show  PGIDs.   Process Group IDs are shown as decimal numbers in parentheses after each process name.  -p implic‐
              itly disables compaction.  If both PIDs and PGIDs are displayed then PIDs are shown first.
       -l     Display long lines.  By default, lines are truncated to the display width or 132 if output is sent to  a  non-tty
              or if the display width is unknown.
       -n     Sort processes with the same ancestor by PID instead of by name.  (Numeric sort.)
       -N     Show  individual  trees  for  each namespace of the type specified.  The available types are: ipc, mnt, net, pid,
              user, uts.  Regular users don't have access to other users' processes information, so the output will be limited.
       -p     Show PIDs.  PIDs are shown as decimal numbers in parentheses after each process  name.   -p  implicitly  disables
              compaction.
       -s     Show parent processes of the specified process.
       -S     Show namespaces transitions.  Like -N, the output is limited when running as a regular user.
       -u     Show  uid transitions.  Whenever the uid of a process differs from the uid of its parent, the new uid is shown in
              parentheses after the process name.
       -U     Use UTF-8 (Unicode) line drawing characters.  Under Linux 1.1-54 and above, UTF-8 mode is entered on the  console
              with echo -e ' 33%8' and left with echo -e ' 33%@'
       -V     Display version information.
       -Z     (SELinux)  Show  security context for each process.  This flag will only work if pstree is compilied with SELinux
              support.
FILES
       /proc  location of the proc file system
BUGS
       Some character sets may be incompatible with the VT100 characters.
SEE ALSO
       ps(1), top(1).
psmisc                                                     2012-07-28                                                 PSTREE(1)
```

Example

```bash
$ pstree
systemd─┬─AliYunDun───18*[{AliYunDun}]
        ├─AliYunDunUpdate───3*[{AliYunDunUpdate}]
        ├─CmsGoAgent.linu───27*[{CmsGoAgent.linu}]
        ├─2*[agetty]
        ├─aliyun-service───5*[{aliyun-service}]
        ├─atd
        ├─auditd───{auditd}
        ├─crond
        ├─dbus-daemon
        ├─dhclient
        ├─irqbalance
        ├─java─┬─controller───2*[{controller}]
        │      └─198*[{java}]
        ├─mysqld───45*[{mysqld}]
        ├─nginx───nginx
        ├─node─┬─node───9*[{node}]
        │      └─9*[{node}]
        ├─ntpd
        ├─php-fpm───20*[php-fpm]
        ├─polkitd───5*[{polkitd}]
        ├─redis-server───3*[{redis-server}]
        ├─rsyslogd───2*[{rsyslogd}]
        ├─sshd───sshd───bash───pstree
        ├─systemd-journal
        ├─systemd-logind
        ├─systemd-udevd
        └─tuned───4*[{tuned}]
```
