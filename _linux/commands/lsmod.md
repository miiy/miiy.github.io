---
layout: post
title: "lsmod"
---

lsmod：list modules 显示已载入系统的模块

```bash
$ lsmod
Module                  Size  Used by
xt_multiport           20480  13
xt_set                 16384  4
ipt_rpfilter           16384  1
ip_set_hash_ip         45056  1
ip_set_hash_net        49152  3
ip_set                 57344  3 ip_set_hash_ip,xt_set,ip_set_hash_net
vxlan                  81920  0
ipt_REJECT             16384  0
nf_reject_ipv4         16384  1 ipt_REJECT
veth                   32768  0
xfrm_user              45056  1
xfrm_algo              16384  1 xfrm_user
wireguard              94208  0
curve25519_x86_64      36864  1 wireguard
libchacha20poly1305    16384  1 wireguard
chacha_x86_64          28672  1 libchacha20poly1305
poly1305_x86_64        28672  1 libchacha20poly1305
ip6_udp_tunnel         16384  2 wireguard,vxlan
udp_tunnel             20480  2 wireguard,vxlan
libcurve25519_generic    49152  2 curve25519_x86_64,wireguard
libchacha              16384  1 chacha_x86_64
nf_conntrack_netlink    57344  0
xt_addrtype            16384  7
xt_statistic           16384  4
xt_nat                 16384  14
xt_tcpudp              20480  35
xt_MASQUERADE          20480  4
nft_chain_nat          16384  7
nf_nat                 57344  3 xt_nat,nft_chain_nat,xt_MASQUERADE
xt_mark                16384  119
xt_conntrack           16384  58
nf_conntrack          176128  5 xt_conntrack,nf_nat,xt_nat,nf_conntrack_netlink,xt_MASQUERADE
nf_defrag_ipv6         24576  1 nf_conntrack
nf_defrag_ipv4         16384  1 nf_conntrack
xt_comment             16384  503
nft_compat             20480  762
nft_counter            16384  400
nf_tables             253952  592 nft_compat,nft_counter,nft_chain_nat
libcrc32c              16384  3 nf_conntrack,nf_nat,nf_tables
nfnetlink              20480  5 nft_compat,nf_conntrack_netlink,nf_tables,ip_set
intel_rapl_msr         20480  0
intel_rapl_common      28672  1 intel_rapl_msr
intel_pmc_core_pltdrv    16384  0
intel_pmc_core         45056  0
ghash_clmulni_intel    16384  0
aesni_intel           372736  0
libaes                 16384  1 aesni_intel
crypto_simd            16384  1 aesni_intel
snd_intel8x0           49152  0
cryptd                 24576  2 crypto_simd,ghash_clmulni_intel
snd_ac97_codec        180224  1 snd_intel8x0
glue_helper            16384  1 aesni_intel
vmwgfx                380928  1
rapl                   20480  0
ac97_bus               16384  1 snd_ac97_codec
snd_pcm               143360  2 snd_intel8x0,snd_ac97_codec
ttm                   114688  1 vmwgfx
snd_timer              49152  1 snd_pcm
drm_kms_helper        278528  1 vmwgfx
joydev                 28672  0
snd                   110592  4 snd_intel8x0,snd_timer,snd_ac97_codec,snd_pcm
soundcore              16384  1 snd
sg                     36864  0
cec                    61440  1 drm_kms_helper
serio_raw              20480  0
evdev                  28672  3
vboxguest              49152  0
button                 24576  0
pcspkr                 16384  0
ac                     16384  0
br_netfilter           32768  0
bridge                258048  1 br_netfilter
stp                    16384  1 bridge
llc                    16384  2 bridge,stp
drm                   626688  4 vmwgfx,drm_kms_helper,ttm
overlay               147456  33
fuse                  167936  1
configfs               57344  1
ip_tables              36864  0
x_tables               53248  14 xt_conntrack,xt_statistic,nft_compat,xt_multiport,xt_tcpudp,xt_addrtype,xt_nat,xt_comment,xt_set,ipt_REJECT,ipt_rpfilter,ip_tables,xt_MASQUERADE,xt_mark
autofs4                53248  2
ext4                  937984  1
crc16                  16384  1 ext4
mbcache                16384  1 ext4
jbd2                  151552  1 ext4
crc32c_generic         16384  0
hid_generic            16384  0
usbhid                 65536  0
hid                   147456  2 usbhid,hid_generic
sd_mod                 61440  2
t10_pi                 16384  1 sd_mod
sr_mod                 28672  0
crc_t10dif             20480  1 t10_pi
crct10dif_generic      16384  0
cdrom                  73728  1 sr_mod
ata_generic            16384  0
ahci                   40960  1
libahci                45056  1 ahci
ata_piix               36864  0
ohci_pci               20480  0
libata                294912  4 ata_piix,libahci,ahci,ata_generic
ohci_hcd               61440  1 ohci_pci
ehci_pci               20480  0
ehci_hcd               98304  1 ehci_pci
psmouse               184320  0
usbcore               331776  5 ohci_hcd,ehci_pci,usbhid,ehci_hcd,ohci_pci
i2c_piix4              28672  0
crct10dif_pclmul       16384  1
crct10dif_common       16384  3 crct10dif_generic,crc_t10dif,crct10dif_pclmul
crc32_pclmul           16384  0
e1000                 163840  0
crc32c_intel           24576  3
usb_common             16384  3 ohci_hcd,usbcore,ehci_hcd
scsi_mod              270336  4 sd_mod,libata,sg,sr_mod
battery                24576  0
video                  57344  0
```
