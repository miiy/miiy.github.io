package main

// import "fmt"
// import "math"

import (
	"fmt"
	"math"
	"math/rand"
	"runtime"
)

func main()  {
	fmt.Println(runtime.GOOS)
    fmt.Println("My favorite number is", rand.Intn(10))
	// fmt.Println(math.pi)
	fmt.Println(math.Pi) // 已导出的
}