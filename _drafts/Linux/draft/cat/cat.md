## cat

```bash
[root@izm5ei2y693gtrhw42czgez head]# cat --help
Usage: cat [OPTION]... [FILE]...
Concatenate FILE(s), or standard input, to standard output.

  -A, --show-all           equivalent to -vET
  -b, --number-nonblank    number nonempty output lines, overrides -n    
  -e                       equivalent to -vE
  -E, --show-ends          display $ at end of each line
  -n, --number             number all output lines
  -s, --squeeze-blank      suppress repeated empty output lines
  -t                       equivalent to -vT
  -T, --show-tabs          display TAB characters as ^I
  -u                       (ignored)
  -v, --show-nonprinting   use ^ and M- notation, except for LFD and TAB 
      --help     display this help and exit
      --version  output version information and exit

With no FILE, or when FILE is -, read standard input.

Examples:
  cat f - g  Output f's contents, then standard input, then g's contents.
  cat        Copy standard input to standard output.

GNU coreutils online help: <http://www.gnu.org/software/coreutils/>
For complete documentation, run: info coreutils 'cat invocation'
```

```bash
[root@izm5ei2y693gtrhw42czgez head]# cat -A file3
1$
$
2$
$
3$
$
4$
$
5$
[root@izm5ei2y693gtrhw42czgez head]# cat -n file3
     1  1
     2
     3  2
     4
     5  3
     6
     7  4
     8
     9  5
[root@izm5ei2y693gtrhw42czgez head]# cat -b file3
     1  1

     2  2

     3  3

     4  4

     5  5
[root@izm5ei2y693gtrhw42czgez head]# cat -n file1
     1  1
     2  2
     3  3
     4  4
     5  5
     6  6
     7  7
     8  8
     9  9
    10  10
    11  11
    12  12
    13  13
    14  14
    15  15
    16  16
    17  17
    18  18
    19  19
    20  20
[root@izm5ei2y693gtrhw42czgez head]# cat -s file4
1

2

3
[root@izm5ei2y693gtrhw42czgez head]# cat file1 file2 > file1file2
[root@izm5ei2y693gtrhw42czgez head]# cat file1file2
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