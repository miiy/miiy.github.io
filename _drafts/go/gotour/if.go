package main

import (
	"fmt"
	"math/rand"
)

func main()  {
	a := 1
	s := ifa(a)
	fmt.Println(s)
}

func ifa(a int) string  {
	if a > 1 {
		return "a bigger than 1"
	}
	return "a not bigger than 1"
}

func ifWithAShortStatement()  {
	if a := rand.Int(); a <
}