## wc

```bash
[root@izm5ei2y693gtrhw42czgez ~]# wc --help
Usage: wc [OPTION]... [FILE]...
  or:  wc [OPTION]... --files0-from=F
Print newline, word, and byte counts for each FILE, and a total line if    
more than one FILE is specified.  With no FILE, or when FILE is -,
read standard input.  A word is a non-zero-length sequence of characters   
delimited by white space.
The options below may be used to select which counts are printed, always in
the following order: newline, word, character, byte, maximum line length.  
  -c, --bytes            print the byte counts
  -m, --chars            print the character counts
  -l, --lines            print the newline counts
      --files0-from=F    read input from the files specified by
                           NUL-terminated names in file F;
                           If F is - then read names from standard input   
  -L, --max-line-length  print the length of the longest line
  -w, --words            print the word counts
      --help     display this help and exit
      --version  output version information and exit

GNU coreutils online help: <http://www.gnu.org/software/coreutils/>        
For complete documentation, run: info coreutils 'wc invocation'
```

```bash
[root@izm5ei2y693gtrhw42czgez head]# wc file1 file2
 20  20  51 file1
 20  20  51 file2
 40  40 102 total
[root@izm5ei2y693gtrhw42czgez head]# wc -l file1
20 file1
[root@izm5ei2y693gtrhw42czgez head]# wc -l file1 file2
 20 file1
 20 file2
 40 total
[root@izm5ei2y693gtrhw42czgez head]# cat file1|wc -l
20
```
