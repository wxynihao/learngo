package main

import (
	"fmt"
	"math"
)

const filename = "abc.txt"

func conts() {
	const a, b = 3, 4
	var c int
	c = int(math.Sqrt(a*a + b*b))
	fmt.Println(filename, c)
}

func enums() {
	const (
		cpp    = 0
		java   = 1
		python = 2
		golang = 3
	)
	fmt.Println(cpp, java, python, golang)

	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)
}

func main() {
	conts()
}
