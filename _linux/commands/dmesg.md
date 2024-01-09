---
layout: post
title: "dmesg"
---

dmesg：display message 显示开机信息

- -T, --ctime                 show human-readable timestamp (may be inaccurate!)

```bash
$ sudo dmesg -T
[Sun Jan  7 11:15:12 2024] Linux version 6.1.0-17-amd64 (debian-kernel@lists.debian.org) (gcc-12 (Debian 12.2.0-14) 12.2.0, GNU ld (GNU Binutils for Debian) 2.40) #1 SMP PREEMPT_DYNAMIC Debian 6.1.69-1 (2023-12-30)
[Sun Jan  7 11:15:12 2024] Command line: BOOT_IMAGE=/boot/vmlinuz-6.1.0-17-amd64 root=UUID=19dc2dc1-a81c-4104-15ae-426c1724d101 ro quiet
[Sun Jan  7 11:15:12 2024] BIOS-provided physical RAM map:
[Sun Jan  7 11:15:12 2024] BIOS-e820: [mem 0x0000000000000000-0x000000000005efff] usable
[Sun Jan  7 11:15:12 2024] BIOS-e820: [mem 0x000000000005f000-0x000000000005ffff] reserved
[Sun Jan  7 11:15:12 2024] BIOS-e820: [mem 0x0000000000060000-0x000000000009ffff] usable
[Sun Jan  7 11:15:12 2024] BIOS-e820: [mem 0x00000000000a0000-0x00000000000fffff] reserved
...
```
