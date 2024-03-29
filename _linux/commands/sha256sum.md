---
layout: post
title: "sha256sum"
---

## sha256sum

sha256sum - compute and check SHA256 message digest

```bash
$ sha256sum --help
Usage: sha256sum [OPTION]... [FILE]...
Print or check SHA256 (256-bit) checksums.
With no FILE, or when FILE is -, read standard input.

  -b, --binary         read in binary mode
  -c, --check          read SHA256 sums from the FILEs and check them
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

The sums are computed as described in FIPS-180-2.  When checking, the input
should be a former output of this program.  The default mode is to print
a line with checksum, a character indicating input mode ('*' for binary,
space for text), and name for each FILE.

GNU coreutils online help: <http://www.gnu.org/software/coreutils/>
For complete documentation, run: info coreutils 'sha256sum invocation'
```


## Example

```bash
$ ls /bin/*sum
/bin/cksum  /bin/innochecksum  /bin/md5sum  /bin/sha1sum  /bin/sha224sum  /bin/sha256sum  /bin/sha384sum  /bin/sha512sum  /bin/sum
```

linux

```bash
$ sha1sum a.txt
2224c10bbfa664739149a8dee9c8f816480bb719  a.txt
$ sha256sum a.txt
6a07a9d561211f1b5639068dc56c0b5eb1960d913f698413ca84ceb3f9698cf1  a.txt
```

Mac

```bash
$ shasum a.txt
ad3a23c53b7ff98c10bd35af2bb0dbd63b8e3d21  a.txt
$ shasum -a 256 a.txt
76e42df215554391da9f95e7c235994a5c87d9215a311147130fa4f342c840ea  a.txt
```
