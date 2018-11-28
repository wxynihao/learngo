package main

import "fmt"

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

func main() {
	fmt.Println(div(10, 3))
	q, r := div(10, 3)
	fmt.Println(q, r)
	a, _ := div2(10, 3)
	fmt.Println(a)
}
