---
layout: post
title: "iconv"
---

```bash
$ iconv --help
Usage: iconv [OPTION...] [FILE...]
Convert encoding of given files from one encoding to another.
 Input/Output format specification:
  -f, --from-code=NAME       encoding of original text
  -t, --to-code=NAME         encoding for output
 Information:
  -l, --list                 list all known coded character sets
 Output control:
  -c                         omit invalid characters from output
  -o, --output=FILE          output file
  -s, --silent               suppress warnings
      --verbose              print progress information
  -?, --help                 Give this help list
      --usage                Give a short usage message
  -V, --version              Print program version
Mandatory or optional arguments to long options are also mandatory or optional
for any corresponding short options.
For bug reporting instructions, please see:
<http://www.gnu.org/software/libc/bugs.html>.
```

## Example

```bash
iconv -l # 列出当前支持的字符编码
iconv -f gb2312 -t utf8 file2 -o file3 # 将文件file1转码转后输出到file2中
```
