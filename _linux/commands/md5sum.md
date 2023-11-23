---
layout: post
title: "md5sum"
---

## md5sum

md5sum - compute and check MD5 message digest

```
$ md5sum --help
Usage: md5sum [OPTION]... [FILE]...
Print or check MD5 (128-bit) checksums.
With no FILE, or when FILE is -, read standard input.

  -b, --binary         read in binary mode
  -c, --check          read MD5 sums from the FILEs and check them
      --tag            create a BSD-style checksum
  -t, --text           read in text mode (default)
  Note: There is no difference between binary and text mode option on GNU system.

The following four options are useful only when verifying checksums:
      --quiet          don't print OK for each successfully verified file
      --status         don't output anything, status code shows success
      --strict         exit non-zero for improperly formatted checksum lines
  -w, --warn           warn about improperly formatted checksum lines

      --help     display this help and exit
      --version  output version information and exit

The sums are computed as described in RFC 1321.  When checking, the input
should be a former output of this program.  The default mode is to print
a line with checksum, a character indicating input mode ('*' for binary,
space for text), and name for each FILE.

GNU coreutils online help: <http://www.gnu.org/software/coreutils/>
For complete documentation, run: info coreutils 'md5sum invocation
```

## Example

```bash
$ ls /bin/*sum
/bin/cksum  /bin/innochecksum  /bin/md5sum  /bin/sha1sum  /bin/sha224sum  /bin/sha256sum  /bin/sha384sum  /bin/sha512sum  /bin/sum
```

linux

```bash
$ md5sum a.txt
4e4c17d165577447097d7792ff76d19a  a.txt
```

Mac

```bash
$ md5 a.txt
MD5 (a.txt) = 12f2c624967fa4d0a88575c13a14753a
```
