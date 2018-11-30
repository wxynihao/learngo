package main

import "fmt"

func arr() {
	var arr1 [5]int
	arr2 := [3]int{1, 3, 5}
	arr3 := [...]int{2, 4, 6, 8, 10}
	var grid [4][5]int

	fmt.Println(arr1, arr2, arr3)
	fmt.Println(grid)
}

func traversal() {
	var arr []int
	for i := range arr {
		fmt.Println(i)
	}
	for i, v := range arr {
		fmt.Println(i, v)
	}
	for _, v := range arr {
		fmt.Println(v)
	}

}

func main() {
	arr()
}
