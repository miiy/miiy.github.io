package main

import "fmt"

func main() {
	hello() // hello
	fmt.Println(add(1, 2)) // 3
	fmt.Println(add2(1, 2, 3.1)) // 6.1
	fmt.Println(swap(1, 2)) // 2 1
	fmt.Println(split(2)) // 3 1
}

func hello() {
	fmt.Println("hello")
}

func add(x int, y int) int {
	return x + y
}
// 当连续两个或多个函数的已命名形参类型相同时，除最后一个类型以外，其它都可以省略。
func add2(x, y int, z float32) float32 {
	return float32(x) + float32(y) + z
}

func swap (x, y int) (int, int) {
	return y, x
}

// Go 的返回值可被命名，它们会被视作定义在函数顶部的变量。
// 没有参数的 return 语句返回已命名的返回值。
func split(x int) (y, z int) {
	y = x + 1
	z = x - 1
	return
}