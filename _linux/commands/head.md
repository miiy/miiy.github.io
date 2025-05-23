---
layout: post
title: "head"
---

## head

```text
file1
1
2
3
4
5
6
7
8
9
10
11
12
13
14
15
16
17
18
19
20

file2
1
2
3
4
5
6
7
8
9
10
11
12
13
14
15
16
17
18
19
20
```

```bash
$ head --help
Usage: head [OPTION]... [FILE]...
Print the first 10 lines of each FILE to standard output.
With more than one FILE, precede each with a header giving the file name.
With no FILE, or when FILE is -, read standard input.
Mandatory arguments to long options are mandatory for short options too.
  -c, --bytes=[-]K         print the first K bytes of each file;
                             with the leading '-', print all but the last
                             K bytes of each file
  -n, --lines=[-]K         print the first K lines instead of the first 10;
                             with the leading '-', print all but the last
                             K lines of each file
  -q, --quiet, --silent    never print headers giving file names
  -v, --verbose            always print headers giving file names
      --help     display this help and exit
      --version  output version information and exit
K may have a multiplier suffix:
b 512, kB 1000, K 1024, MB 1000*1000, M 1024*1024,
GB 1000*1000*1000, G 1024*1024*1024, and so on for T, P, E, Z, Y.
GNU coreutils online help: <http://www.gnu.org/software/coreutils/>
For complete documentation, run: info coreutils 'head invocation'
```

```bash
$ head file1
1
2
3
4
5
6
7
8
9
10
$ head file1 file2
==> file1 <==
1
2
3
4
5
6
7
8
9
10
==> file2 <==
1
2
3
4
5
6
7
8
9
10
$ head -n 2 file1
1
2
$ head -n2 file1
1
2
$ head -c 4 file1
1
2
$ head -q file1 file2
1 
2 
3 
4 
5 
6 
7 
8 
9 
10
1 
2 
3 
4 
5 
6 
7
8
9
10
$ head -v file1
==> file1 <==
1
2
3
4
5
6
7
8
9
10
```

使用 head 和 tail 把一个文件的前900行移动到另一个文件中

```bash
head -n 900 source.txt > destination.txt
tail -n +901 source.txt > temp.txt && mv temp.txt source.txt
```
