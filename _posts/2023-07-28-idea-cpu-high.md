---
layout: post
title: "解决 idea CPU 占用过高的问题"
date: 2023-07-28
tags: linux
---

可以通过查看idea自带的监控来观测是什么进程占用了CPU资源

Help >> Diagnostic Tools >> Activity Monitor

JIT 导致CPU占用过高

修改 idea.vmoptions 文件

```bash
#-Xms128m
#-Xmx750m
-XX:ReservedCodeCacheSize=512m
-XX:+UseG1GC
-XX:SoftRefLRUPolicyMSPerMB=50
#-XX:CICompilerCount=2
-XX:+HeapDumpOnOutOfMemoryError
-XX:-OmitStackTraceInFastThrow
-XX:+IgnoreUnrecognizedVMOptions
-XX:CompileCommand=exclude,com/intellij/openapi/vfs/impl/FilePartNodeRoot,trieDescend
-ea
-Dsun.io.useCanonCaches=false
-Dsun.java2d.metal=true
-Djbr.catch.SIGABRT=true
-Djdk.http.auth.tunneling.disabledSchemes=""
-Djdk.attach.allowAttachSelf=true
-Djdk.module.illegalAccess.silent=true
-Dkotlinx.coroutines.debug=off
-XX:ErrorFile=$USER_HOME/java_error_in_idea_%p.log
-XX:HeapDumpPath=$USER_HOME/java_error_in_idea.hprof

#
-Xms1024m
-Xmx2048m
# JIT 参数
 
# 设置用于编译的编译器线程数
-XX:CICompilerCount=2
# 开启分层编译
-XX:TieredStopAtLevel=1
# 控制最大数量嵌套调用内联
-XX:MaxInlineLevel=3
-XX:Tier4MinInvocationThreshold=100000
-XX:Tier4InvocationThreshold=110000
-XX:Tier4CompileThreshold=120000
```