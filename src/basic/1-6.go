package main

import (
	"fmt"
	"reflect"
	"runtime"
)

func div(a, b int) (int, int) {
	return a / b, a % b
}

func div2(a, b int) (q, r int) {
	return a / b, a % b
}

func div3(a, b int) (q, r int) {
	q = a / b
	r = a % b
	return
}

func evalWithError(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		q, _ := div(a, b)
		return q, nil
	default:
		return 0, fmt.Errorf(
			"unsupported operation: %s", op)
	}
}

func apply(op func(int, int) int, a, b int) int {
	//通过反射拿到函数的方法
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("Calling function %s with args "+
		"(%d, %d)\n", opName, a, b)

	return op(a, b)
}

func sum(numbers ...int) int {
	s := 0
	for i := range numbers {
		s += numbers[i]
	}
	return s
}

func main() {
	fmt.Println(div(10, 3))
	q, r := div(10, 3)
	fmt.Println(q, r)
	a, _ := div2(10, 3)
	fmt.Println(a)
}
