package main

import "fmt"

func main() {
	var a = [3]int{1, 2, 3} // a是拥有3个整型元素的数组，被初始化1到3
	fmt.Println(a)          // [1 2 3]
	fmt.Println(a[1:2])     // [2]

	fmt.Println(a[len(a) - 1])
	fmt.Println(a[:len(a)])
}