---
layout: post
title: "hexdump"
---

## hexdump

hexdump - display file contents in ascii, decimal, hexadecimal, or octal

hexdump - 以 ascii，十进制，十六进制或八进制显示文件内容

```bash
$ hexdump -h
hexdump: invalid option -- 'h'

Usage:
 hexdump [options] file...

Options:
 -b              one-byte octal display
 -c              one-byte character display
 -C              canonical hex+ASCII display
 -d              two-byte decimal display
 -o              two-byte octal display
 -x              two-byte hexadecimal display
 -e format       format string to be used for displaying data
 -f format_file  file that contains format strings
 -n length       interpret only length bytes of input
 -s offset       skip offset bytes from the beginning
 -v              display without squeezing similar lines
 -V              output version information and exit
```

## Example

```bash
$ echo "ABCDEFGHIJKLMNOPQRSTUVWXYZ" > test.txt

$ hexdump -b test.txt
0000000 101 102 103 104 105 106 107 110 111 112 113 114 115 116 117 120
0000010 121 122 123 124 125 126 127 130 131 132 012
000001b

$ hexdump -c test.txt
0000000   A   B   C   D   E   F   G   H   I   J   K   L   M   N   O   P
0000010   Q   R   S   T   U   V   W   X   Y   Z  \n
000001b

$ hexdump -C test.txt
00000000  41 42 43 44 45 46 47 48  49 4a 4b 4c 4d 4e 4f 50  |ABCDEFGHIJKLMNOP|
00000010  51 52 53 54 55 56 57 58  59 5a 0a                 |QRSTUVWXYZ.|
0000001b

$ hexdump -n 10 -C test.txt
00000000  41 42 43 44 45 46 47 48  49 4a                    |ABCDEFGHIJ|
0000000a

$ hexdump -n 10 -s 2 -C test.txt
00000002  43 44 45 46 47 48 49 4a  4b 4c                    |CDEFGHIJKL|
0000000c

$ hexdump -n 100 -C a.zip
00000000  50 4b 03 04 14 00 00 00  08 00 d6 84 fb 50 9d ef  |PK...........P..|
00000010  c8 43 7a 5d 00 00 f5 c0  01 00 21 00 00 00 d0 c2  |.Cz]......!.....|
00000020  bd a8 ce c4 bc fe bc d0  2f 53 51 4c b4 f2 b7 d6  |......../SQL....|
00000030  d6 b5 2d b9 e6 d4 f2 bc  f2 ca f6 2e 64 6f 63 ec  |..-.........doc.|
00000040  5d 6b 57 db cc 11 fe dc  9e f3 fe 87 6d 7a 5a 20  |]kW.........mzZ |
00000050  c5 60 f9 86 45 42 5a 19  4c 4b 4b 20 0d a4 d7 d3  |.`..EBZ.LKK ....|
00000060  f6 08 23 c0                                       |..#.|
00000064

```