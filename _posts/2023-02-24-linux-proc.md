---
layout: post
title: "proc"
date: 2023-02-24
tags: linux
---

```bash
$ ls /proc
1     12850  174  24   3574  569  604  8833	  cpuinfo      interrupts   kpagecount	pagetypeinfo  sysrq-trigger
10    12972  175  25   4     570  606  9	  crypto       iomem	    kpageflags	partitions    sysvipc
107   12978  18   255  445   572  642  942	  devices      ioports	    loadavg	sched_debug   thread-self
1075  13003  19   26   472   574  681  945	  diskstats    irq	    locks	schedstat     timer_list
11    14     2	  27   48    59   7    acpi	  dma	       kallsyms     meminfo	self	      tty
112   15     20   28   49    590  717  buddyinfo  driver       kcore	    misc	slabinfo      uptime
113   155    21   29   50    594  759  bus	  execdomains  keys	    modules	softirqs      version
115   16     22   3    562   595  771  cgroups	  fb	       key-users    mounts	stat	      vmallocinfo
116   17     221  30   563   596  796  cmdline	  filesystems  kmsg	    mtrr	swaps	      vmstat
12    172    23   355  566   6	  8    consoles   fs	       kpagecgroup  net		sys	      zoneinfo
```

```bash
$ cat /proc/cpuinfo 
processor	: 0
vendor_id	: GenuineIntel
cpu family	: 6
model		: 79
model name	: Intel(R) Xeon(R) CPU E5-2682 v4 @ 2.50GHz
stepping	: 1
microcode	: 0x1
cpu MHz		: 2500.026
cache size	: 40960 KB
physical id	: 0
siblings	: 1
core id		: 0
cpu cores	: 1
apicid		: 0
initial apicid	: 0
fpu		: yes
fpu_exception	: yes
cpuid level	: 20
wp		: yes
flags		: fpu vme de pse tsc msr pae mce cx8 apic sep mtrr pge mca cmov pat pse36 clflush mmx fxsr sse sse2 ss ht syscall nx pdpe1gb rdtscp lm constant_tsc rep_good nopl nonstop_tsc cpuid tsc_known_freq pni pclmulqdq ssse3 fma cx16 pcid sse4_1 sse4_2 x2apic movbe popcnt tsc_deadline_timer aes xsave avx f16c rdrand hypervisor lahf_lm abm 3dnowprefetch invpcid_single pti ibrs ibpb stibp fsgsbase tsc_adjust bmi1 hle avx2 smep bmi2 erms invpcid rtm rdseed adx smap xsaveopt arat
bugs		: cpu_meltdown spectre_v1 spectre_v2 spec_store_bypass l1tf mds swapgs taa itlb_multihit mmio_stale_data
bogomips	: 5000.05
clflush size	: 64
cache_alignment	: 64
address sizes	: 46 bits physical, 48 bits virtual
power management:

```

```bash
$ cat /proc/meminfo
MemTotal:        2042900 kB
MemFree:          529904 kB
MemAvailable:    1321508 kB
Buffers:          150536 kB
Cached:           747436 kB
SwapCached:            0 kB
Active:           971416 kB
Inactive:         430612 kB
Active(anon):     504264 kB
Inactive(anon):      252 kB
Active(file):     467152 kB
Inactive(file):   430360 kB
Unevictable:           0 kB
Mlocked:               0 kB
SwapTotal:             0 kB
SwapFree:              0 kB
Dirty:                12 kB
Writeback:             0 kB
AnonPages:        492324 kB
Mapped:           204768 kB
Shmem:               464 kB
Slab:              75372 kB
SReclaimable:      51284 kB
SUnreclaim:        24088 kB
KernelStack:        3164 kB
PageTables:         3316 kB
NFS_Unstable:          0 kB
Bounce:                0 kB
WritebackTmp:          0 kB
CommitLimit:     1021448 kB
Committed_AS:    1520316 kB
VmallocTotal:   34359738367 kB
VmallocUsed:           0 kB
VmallocChunk:          0 kB
Percpu:              808 kB
HardwareCorrupted:     0 kB
AnonHugePages:    399360 kB
ShmemHugePages:        0 kB
ShmemPmdMapped:        0 kB
HugePages_Total:       0
HugePages_Free:        0
HugePages_Rsvd:        0
HugePages_Surp:        0
Hugepagesize:       2048 kB
Hugetlb:               0 kB
DirectMap4k:       83840 kB
DirectMap2M:     2013184 kB
DirectMap1G:           0 kB
```

```bash
$ cat /proc/version
Linux version 4.19.0-23-amd64 (debian-kernel@lists.debian.org) (gcc version 8.3.0 (Debian 8.3.0-6)) #1 SMP Debian 4.19.269-1 (2022-12-20)
```

```bash
$ cat /proc/mounts
sysfs /sys sysfs rw,nosuid,nodev,noexec,relatime 0 0
proc /proc proc rw,nosuid,nodev,noexec,relatime 0 0
udev /dev devtmpfs rw,nosuid,relatime,size=1006304k,nr_inodes=251576,mode=755 0 0
devpts /dev/pts devpts rw,nosuid,noexec,relatime,gid=5,mode=620,ptmxmode=000 0 0
tmpfs /run tmpfs rw,nosuid,noexec,relatime,size=204292k,mode=755 0 0
/dev/vda1 / ext4 rw,relatime,errors=remount-ro 0 0
securityfs /sys/kernel/security securityfs rw,nosuid,nodev,noexec,relatime 0 0
tmpfs /dev/shm tmpfs rw,nosuid,nodev 0 0
tmpfs /run/lock tmpfs rw,nosuid,nodev,noexec,relatime,size=5120k 0 0
tmpfs /sys/fs/cgroup tmpfs ro,nosuid,nodev,noexec,mode=755 0 0
cgroup2 /sys/fs/cgroup/unified cgroup2 rw,nosuid,nodev,noexec,relatime,nsdelegate 0 0
cgroup /sys/fs/cgroup/systemd cgroup rw,nosuid,nodev,noexec,relatime,xattr,name=systemd 0 0
pstore /sys/fs/pstore pstore rw,nosuid,nodev,noexec,relatime 0 0
bpf /sys/fs/bpf bpf rw,nosuid,nodev,noexec,relatime,mode=700 0 0
cgroup /sys/fs/cgroup/net_cls,net_prio cgroup rw,nosuid,nodev,noexec,relatime,net_cls,net_prio 0 0
cgroup /sys/fs/cgroup/devices cgroup rw,nosuid,nodev,noexec,relatime,devices 0 0
cgroup /sys/fs/cgroup/pids cgroup rw,nosuid,nodev,noexec,relatime,pids 0 0
cgroup /sys/fs/cgroup/cpu,cpuacct cgroup rw,nosuid,nodev,noexec,relatime,cpu,cpuacct 0 0
cgroup /sys/fs/cgroup/freezer cgroup rw,nosuid,nodev,noexec,relatime,freezer 0 0
cgroup /sys/fs/cgroup/memory cgroup rw,nosuid,nodev,noexec,relatime,memory 0 0
cgroup /sys/fs/cgroup/perf_event cgroup rw,nosuid,nodev,noexec,relatime,perf_event 0 0
cgroup /sys/fs/cgroup/rdma cgroup rw,nosuid,nodev,noexec,relatime,rdma 0 0
cgroup /sys/fs/cgroup/blkio cgroup rw,nosuid,nodev,noexec,relatime,blkio 0 0
cgroup /sys/fs/cgroup/cpuset cgroup rw,nosuid,nodev,noexec,relatime,cpuset 0 0
systemd-1 /proc/sys/fs/binfmt_misc autofs rw,relatime,fd=38,pgrp=1,timeout=0,minproto=5,maxproto=5,direct,pipe_ino=10209 0 0
mqueue /dev/mqueue mqueue rw,relatime 0 0
debugfs /sys/kernel/debug debugfs rw,relatime 0 0
hugetlbfs /dev/hugepages hugetlbfs rw,relatime,pagesize=2M 0 0
tracefs /sys/kernel/debug/tracing tracefs rw,relatime 0 0
tmpfs /run/user/0 tmpfs rw,nosuid,nodev,relatime,size=204288k,mode=700 0 0
```

```bash
$ cat /proc/net/dev
Inter-|   Receive                                                |  Transmit
 face |bytes    packets errs drop fifo frame compressed multicast|bytes    packets errs drop fifo colls carrier compressed
    lo:  180471     819    0    0    0     0          0         0   180471     819    0    0    0     0       0          0
docker0:       0       0    0    0    0     0          0         0        0       0    0    0    0     0       0          0
  eth0: 154678848  986010    0    0    0     0          0         0 446199126  803036    0    0    0     0       0          0
```

```bash
cat /proc/swaps
Filename        Type    Size  Used  Priority
```

```bash
cat /proc/interrupts
           CPU0       
  0:          3   IO-APIC   2-edge      timer
  1:          9   IO-APIC   1-edge      i8042
  4:        495   IO-APIC   4-edge      ttyS0
  6:          3   IO-APIC   6-edge      floppy
  8:          0   IO-APIC   8-edge      rtc0
  9:          0   IO-APIC   9-fasteoi   acpi
 11:         32   IO-APIC  11-fasteoi   uhci_hcd:usb1, virtio3
 12:         15   IO-APIC  12-edge      i8042
 14:          0   IO-APIC  14-edge      ata_piix
 15:          0   IO-APIC  15-edge      ata_piix
 24:          0   PCI-MSI 49152-edge      virtio0-config
 25:     705063   PCI-MSI 49153-edge      virtio0-input.0
 26:         13   PCI-MSI 49154-edge      virtio0-output.0
 27:          0   PCI-MSI 81920-edge      virtio2-config
 28:     262355   PCI-MSI 81921-edge      virtio2-req.0
 29:          0   PCI-MSI 65536-edge      virtio1-config
 30:         22   PCI-MSI 65537-edge      virtio1-virtqueues
NMI:          0   Non-maskable interrupts
LOC:  861860291   Local timer interrupts
SPU:          0   Spurious interrupts
PMI:          0   Performance monitoring interrupts
IWI:      16317   IRQ work interrupts
RTR:          0   APIC ICR read retries
RES:          0   Rescheduling interrupts
CAL:          0   Function call interrupts
TLB:          0   TLB shootdowns
TRM:          0   Thermal event interrupts
THR:          0   Threshold APIC interrupts
DFR:          0   Deferred Error APIC interrupts
MCE:          0   Machine check exceptions
MCP:       1844   Machine check polls
ERR:          0
MIS:          0
PIN:          0   Posted-interrupt notification event
NPI:          0   Nested posted-interrupt event
PIW:          0   Posted-interrupt wakeup event
```
