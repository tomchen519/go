package main

import (
	"fmt"
)

func main() {
	var x int
	p := &x
	*p = 10
	fmt.Println(x)
}
