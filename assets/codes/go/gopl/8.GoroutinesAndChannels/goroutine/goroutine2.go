package main

import (
	"fmt"
	"time"
)

func main() {
	go sleep()
	fmt.Println("main print");
}

func sleep() {
	time.Sleep(100 * time.Millisecond)
}