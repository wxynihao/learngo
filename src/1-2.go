package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

func euler() {
	var a = cmplx.Exp(1i*math.Pi) + 1
	var b = cmplx.Pow(math.E, 1i*math.Pi) + 1
	fmt.Println(a)
	fmt.Println(b)
	fmt.Printf("%.3f\n", a)
}

func triangle() {
	var a, b int = 3, 4
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}

func main() {
	euler()
	triangle()
}
