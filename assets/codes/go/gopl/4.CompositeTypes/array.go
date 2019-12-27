package main

import "fmt"

func main() {
	var a [3]int // 3个整数的数组
	fmt.Println(a[0]) // 输出数组的第一个元素
	fmt.Println(a[len(a) - 1]) // 输出数组的最后一个元素，即 [2]

	// 输出索引和元素
	for i, v := range a {
		fmt.Printf("%d %d\n", i, v)
	}
	// 仅输出元素
	for _, v := range a {
		fmt.Printf("%d\n", v)
	}

	// 初始化数组
	var q [3]int = [3]int{1, 2, 3}
	var r [3]int = [3]int{1, 2}
	fmt.Println(q) // [1 2 3]
	fmt.Println(r[2]) // 0

	// 如果省略号“...”出现在数组长度的位置，那么数组的长度由初始化数组的元素个数决定。以上q的定义可以简化为：
	q2 := [...]int{1, 2, 3}
	fmt.Printf("%T\n", q2) // [3]int
	// 数组的长度是数组类型的一部分，所以 [3]int 和 [4]int 是两种不同的数据类型。数组的长度必须是常亮表达式，也就是说，这个表达式的值在程序编译时就可以确定。

	q3 := [3]int{1, 2, 3}
	q3 = [3]int{2, 3, 4}
	// q3 = [4]int{1, 2, 3, 4} // 编译错误：不可以将 [4]int 赋值给 [3]int
	fmt.Println(q3)

	type Currency int
	const (
		USD Currency = iota
		EUR
		GBP
		RMB
	)
	symbol := [...]string{USD: "$", EUR: "€", GBP: "£", RMB: "￥"}
	fmt.Println(RMB, symbol[RMB]) // 3 ￥
}

/*

数组是具有固定长度且拥有零个或者多个相同数据类型元素的序列。

数组中的每个元素是通过索引来访问的，索引从 0 到数组长度减 1。

```bash
root@eff37d23cbbe:/go# go run array.go
0
0
0 0
1 0
2 0
0
0
0
[1 2 3]
0
```
*/