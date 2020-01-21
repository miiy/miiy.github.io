package main

import "fmt"

func main()  {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	// 初始化语句和后置语句是可选的。可以去掉分号相当于C语言的while
	for ;sum < 100; {
		sum += sum
	}
	fmt.Println(sum)

	// while
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)

	// 无限循环
	for {

	}
}