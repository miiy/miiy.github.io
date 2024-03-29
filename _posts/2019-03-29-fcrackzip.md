---
layout: post
title: "zip 密码破解 fcrackzip"
date: 2019-03-29
tags: linux
---

# fcrackzip

fcrackzip - a Free/Fast Zip Password Cracker

## help

```bash
root@kali:~# fcrackzip -h

fcrackzip version 1.0, a fast/free zip password cracker
written by Marc Lehmann <pcg@goof.com> You can find more info on
http://www.goof.com/pcg/marc/

USAGE: fcrackzip
          [-b|--brute-force]            use brute force algorithm
          [-D|--dictionary]             use a dictionary
          [-B|--benchmark]              execute a small benchmark
          [-c|--charset characterset]   use characters from charset
          [-h|--help]                   show this message
          [--version]                   show the version of this program
          [-V|--validate]               sanity-check the algortihm
          [-v|--verbose]                be more verbose
          [-p|--init-password string]   use string as initial password/file
          [-l|--length min-max]         check password with length min to max
          [-u|--use-unzip]              use unzip to weed out wrong passwords
          [-m|--method num]             use method number "num" (see below)
          [-2|--modulo r/m]             only calculcate 1/m of the password
          file...                    the zipfiles to crack

methods compiled in (* = default):

 0: cpmask
 1: zip1
*2: zip2, USE_MULT_TAB
```

## Manual

```bash
root@kali:~# man frackzip
FCRACKZIP(1)                       General Commands Manual                       FCRACKZIP(1)

NAME
       fcrackzip - a Free/Fast Zip Password Cracker

SYNOPSIS
       fcrackzip  [-bDBchVvplum2]  [--brute-force]  [--dictionary]  [--benchmark]  [--charset
       characterset]  [--help]   [--validate]   [--verbose]   [--init-password   string/path]
       [--length min-max] [--use-unzip] [--method name] [--modulo r/m] file...

DESCRIPTION
       fcrackzip searches each zipfile given for encrypted files and tries to guess the pass-
       word. All files must be encrypted with the same password, the more files you  provide,
       the better.

   OPTIONS
       -h, --help
              Prints the version number and (hopefully) some helpful insights.

       -v, --verbose
              Each -v makes the program more verbose.

       -b, --brute-force
              Select  brute  force  mode. This tries all possible combinations of the letters
              you specify.

       -D, --dictionary
              Select dictionary mode. In this mode, fcrackzip  will  read  passwords  from  a
              file,  which  must  contain  one password per line and should be alphabetically
              sorted (e.g. using sort(1)).

       -c, --charset characterset-specification
              Select the characters to use in brute-force cracking. Must be one of

                a   include all lowercase characters [a-z]
                A   include all uppercase characters [A-Z]
                1   include the digits [0-9]
                !   include [!:$%&/()=?{[]}+*~#]
                :   the following characters up to the end of the spe-
                    cification string are included in the character set.
                    This way you can include any character except binary
                    null (at least under unix).

              For example, a1:$% selects lowercase characters, digits and the dollar and per-
              cent signs.

       -p, --init-password string
              Set initial (starting) password for brute-force searching to string, or use the
              file with the name string to supply passwords for dictionary searching.

       -l, --length min[-max]
              Use an initial password of length min, and check all passwords up to  passwords
              of length max (including). You can omit the max parameter.

       -u, --use-unzip
              Try  to  decompress  the first file by calling unzip with the guessed password.
              This weeds out false positives when not enough files have been given.

       -m, --method name
              Use method number "name" instead of the default  cracking  method.  The  switch
              --help  will  print  a  list of available methods. Use --benchmark to see which
              method does perform best on your machine. The name can also be  the  number  of
              the method to use.

       -2, --modulo r/m
              Calculate only r/m of the password. Not yet supported.

       -B, --benchmark
              Make a small benchmark, the output is nearly meaningless.

       -V, --validate
              Make some basic checks whether the cracker works.

ZIP PASSWORD BASICS
       Have  you  ever mis-typed a password for unzip? Unzip reacted pretty fast with 'incor-
       rect password', without decrypting the whole file. While the encryption algorithm used
       by  zip  is  relatively secure, PK made cracking easy by providing hooks for very fast
       password-checking, directly in the zip file. Understanding these  is  crucial  to  zip
       password cracking:

       For each password that is tried, the first twelve bytes of the file are decrypted. De-
       pending on the version of zip used to encrypt the file (more on that later), the first
       ten  or  eleven bytes are random, followed by one or two bytes whose values are stored
       elsewhere in the zip file, i.e. are known beforehand. If these last bytes  don't  have
       the correct (known) value, the password is definitely wrong. If the bytes are correct,
       the password might be correct, but the only method to find out is to  unzip  the  file
       and compare the uncompressed length and crc's.

       Earlier versions of pkzip (1.xx) (and, incidentally, many zip clones for other operat-
       ing systems!) stored two known bytes. Thus the error rate was roughly 1/2^16 =  0.01%.
       PKWARE  'improved'  (interesting  what  industry calls improved) the security of their
       format by only including one byte, so the possibility of false passwords is now raised
       to  0.4%.  Unfortunately,  there  is no real way to distinguish one byte from two byte
       formats, so we have to be conservative.

BRUTE FORCE MODE
       By default, brute force starts at the given starting password, and successively  tries
       all combinations until they are exhausted, printing all passwords that it detects, to-
       gether with a rough correctness indicator.

       The starting password given by the -p switch determines the  length.   fcrackzip  will
       not  currently  increase  the  password  length automatically, unless the -l switch is
       used.

DICTIONARY MODE
       This mode is similar to brute force mode, but instead of generating passwords using  a
       given  set of characters and a length, the passwords will be read from a file that you
       have to specify using the -p switch.

CP MASK
       A CP mask is a method to obscure images or parts of images using  a  password.   These
       obscured  images can be restored even when saved as JPEG files. In most of these files
       the password is actually hidden and can be decoded  easily  (using  one  of  the  many
       available  viewer  and  masking programs, e.g. xv). If you convert the image the pass-
       word, however, is lost. The cpmask crack method can be used to brute-force  these  im-
       ages. Instead of a zip file you supply the obscured part (and nothing else) of the im-
       age in the PPM-Image Format (xv and other viewers can easily do this).

       The cpmask method can only cope with password composed of  uppercase  letters,  so  be
       sure to supply the --charset A or equivalent option, together with a suitable initial-
       ization password.

EXAMPLES
       fcrackzip -c a -p aaaaaa sample.zip
              checks the encrypted files in sample.zip for all lowercase  6  character  pass-
              words (aaaaaa ... abaaba ... ghfgrg ... zzzzzz).

       fcrackzip --method cpmask --charset A --init AAAA test.ppm
              checks the obscured image test.ppm for all four character passwords.

       fcrackzip -D -p passwords.txt sample.zip
              check for every password listed in the file passwords.txt.

PERFORMANCE
       fzc,  which  seems to be widely used as a fast password cracker, claims to make 204570
       checks per second on my machine (measured under plain dos w/o memory manager).

       fcrackzip, being written in C and not in assembler, naturally is slower. Measured on a
       slightly  loaded  unix  (same  machine), it's 12 percent slower (the compiler used was
       pgcc, from http://www.gcc.ml.org/).

       To remedy this a bit, I converted small parts of the encryption core to x86  assembler
       (it  will  still  compile on non x86 machines), and now it's about 4-12 percent faster
       than fzc (again, the fcrackzip performance was measured under a  multitasking  os,  so
       there  are inevitably some meaurement errors), so there shouldn't be a tempting reason
       to switch to other programs.

       Further improvements are definitely possible: fzc took 4  years  to  get  into  shape,
       while  fcrackzip was hacked together in under 10 hours. And not to forget you have the
       source, while other programs (like fzc), even come as an encrypted  .exe  file  (maybe
       because their programmers are afraid of other people could having a look at their lack
       of programming skills?  nobody knows...)

RATIONALE
       The reason I wrote fcrackzip was NOT to have the fastest zip cracker available, but to
       provide a portable, free (thus extensible), but still fast zip password cracker. I was
       really pissed of with that dumb, nonextendable zipcrackers that were either slow, were
       too  limited,  or wouldn't run in the background (say, under unix). (And you can't run
       them on your superfast 600Mhz Alpha).

BUGS
       No automatic unzip checking.

       Stop/resume facility is missing.

       Should be able to distinguish between files with 16 bit stored CRC's and 8 bit  stored
       CRC's.

       The benchmark does not work on all systems.

       It's still early alpha.

       Method "cpmask" only accepts ppms.

       Could be faster.

AUTHOR
       fcrackzip  was  written  by Marc Lehmann <pcg@goof.com>. The main fcrackzip page is at
       http://www.goof.com/pcg/marc/fcrackzip.html)

                                Free/Fast Zip Password Cracker                   FCRACKZIP(1)
```


## 部分重要参数说明

* -b 使用暴力算法
* -D 使用字典
* -c 使用字符集中的字符
* -v 更详细的信息
* -p 使用字符串作为初始密码/文件
* -l 长度
* -u 使用解压缩清除错误的密码

## Example

```bash
root@kali:~# echo 123 > sample.txt
root@kali:~# zip -rP 123456 sample.zip sample.txt
  adding: sample.txt (stored 0%)
root@kali:~# fcrackzip -b -c 'aA1!:' -l 1-10 -u sample.zip

root@kali:~# fcrackzip -b -c '1' -l 1-6 -u sample.zip


PASSWORD FOUND!!!!: pw == 123456
```

```bash
root@kali:~# zip -rP abcd1234 sample2.zip sample.txt
  adding: sample.txt (stored 0%)
root@kali:~# fcrackzip -D -p dict.txt -u sample2.zip 


PASSWORD FOUND!!!!: pw == abcd1234
```
