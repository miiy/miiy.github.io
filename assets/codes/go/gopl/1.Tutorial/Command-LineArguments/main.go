/*
os包以跨平台的方式，提供了一些与操作系统交互的函数和变量。
程序的命令行参数可从os包的Args变量获取；os包外部使用os.Args访问该变量。
os.Args变量是一个字符串（string）的切片（slice）

s[i] 访问单个元素
s[m:n] 获取子序列
len(s) 序列的元素数目

与大部分数编程语言类似，在 Go 中，所有的索引使用搬开区间，即包含第一个索引，不包含最后一个索引。
slice s[m:n]，其中 0 <= m <= n <= len(s)，包含n-m个元素

os.Args 的第一个元素 os.Args[0] 命令本身的名字

s[m:n] 表示从第m个到第n-1个元素的slice
如果m 或 n 缺失，默认分别是 0 或 len(s)
*/

package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}

// var 关键字生命了两个 string 类型的变量 s 和 sep。
// s 和 sep 隐式初始化为空字符串
// 当应用于字符串时，+ 操作符对字符串的值进行追加操作
