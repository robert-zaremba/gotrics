package main

import (
	"fmt"
)

type X struct {
	s string
	i int
}

func main() {
	x1 := X{"jeden", 1}
	x1a := x1

	m := make(map[X]int)
	m[x1] = 2
	m[x1a] = 4
	m[X{"jeden", 2}] = 8
	println(m[x1])
	println(m[X{"jeden", 2}])
}
