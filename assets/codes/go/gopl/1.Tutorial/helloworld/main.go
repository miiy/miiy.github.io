package main

import "fmt"

func main() {
	fmt.Println("Hello, 世界")
}

/*

Go语言原生支持Unicode，它可以处理全世界任何语言的文本。


package 表示该 Go 代码所属的包。

import 用于导入该程序所依赖的包。

Go 语言的main()函数不能带参数，也不能定义返回值。


go run：编译并运行

将编译、链接和运行3个步骤合并为一步，运行完后在当前目录下也看不到任何中间文件和最终的可执行文件。

```bash
# go run 1.hello.go
Hello, world.
```
go build：编译

```bash
# go build 1.hello.go
# ./1.hello
Hello, world.
```

run:
# go run 1/helloworld/main.go

output:

Hello, 世界

 */