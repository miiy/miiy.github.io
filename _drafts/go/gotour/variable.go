package main

import "fmt"

var a bool
var b, c int
var d, e int = 1, 2

func main()  {
	var f, g = true, "g"
	h := "h" // 短变量声明
	i , j := "i", 1
	fmt.Println(a, b, c, d, e, f, g, h, i, j) // false 0 0 1 2 true g h i 1

	typeConversions()

	typeInference()
}

func typeConversions() {
	var i int = 1
	j := float64(i) + 1.01
	fmt.Println(j) // 2.01
}

// 类型推导
func typeInference()  {
	var x int
	y := x
	fmt.Printf("j is of type %T\n", y) // j is of type int
}