/*
goroutine

在 Go 里，每一个并发执行活动成为 goroutine。

主 goroutine：当一个程序启动时，只有一个 goroutine 来调用 main 函数，称它为主 goroutine。新的 goroutine 通过 go 语句进行创建。

f() // 调用 f(); 等待它返回
go f() // 新建一个调用 f() 的 goroutine，不用等待

主 goroutine 计算第 45 个裴波那契数。因为它使用了非常低效的递归算法，所以它需要大量的时间来执行。

spinner 旋转器 用来指示程序依然在运行。

当 main 函数返回时，所有的 goroutine 都强制地直接终结，然后退出程序

指示器和裴波那契数计算在同时运行
*/

package main

import (
	"fmt"
    "time"
)

func main() {
	go spinner(100 * time.Millisecond)
	const n = 45
	fibN := fib(n) // slow
	fmt.Printf("\rFibonacci(%d) = %d\n", n, fibN)
}

func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}

/*

```bash
# go run goroutine.go
Fibonacci(45) = 1134903170
```
*/