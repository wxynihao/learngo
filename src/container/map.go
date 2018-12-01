package main

import "fmt"

func baseMap() {
	m := map[string]string{
		"id":   "1001",
		"name": "zhangsan",
	}
	fmt.Println(m, len(m))

	for k, v := range m {
		fmt.Println(k, v)
	}

	fmt.Println(m["id"])

	if id, ok := m["id"]; ok {
		fmt.Println(id, ok)
	}

	delete(m, "name")
}

func enptyMap() {
	m2 := make(map[string]int)
	var m3 map[string]int
	fmt.Println(m2, m3)
}

func main() {
	baseMap()
}
