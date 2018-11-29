package main

import "fmt"

func basic() {
	var a int = 2
	var pa *int = &a
	*pa = 3
	fmt.Println(a)
}

func swap(a, b *int) {
	*a, *b = *b, *a
}

func swap2(a, b int) (int, int) {
	return b, a
}

func main() {
	a, b := 1, 2
	swap(&a, &b)
	fmt.Println(a, b)

	c, d := 3, 4
	c, d = swap2(c, d)
	fmt.Println(c, d)
}
