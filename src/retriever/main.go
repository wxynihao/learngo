package main

import (
	"fmt"
	"learngo/src/retriever/mock"
	"learngo/src/retriever/real"
)

type Retriver interface {
	Get(url string) string
}

func download(r Retriver) string {
	return r.Get("http://www.baidu.com")
}

func main() {
	var r Retriver
	r = mock.Retriver{"fake download"}
	inspect(r)
	r = &real.Retriver{}
	inspect(r)
	realRetriver := r.(*real.Retriver)
	fmt.Println(realRetriver.TimeOut)
	fmt.Printf("%T %v\n", r, r)
	//fmt.Println(download(r))
}

func inspect(r Retriver) {
	fmt.Println("Inspecting", r)
	fmt.Printf(" > Type:%T Value:%v\n", r, r)
	fmt.Print(" > Type switch: ")
	switch v := r.(type) {
	case *mock.Retriver:
		fmt.Println("Contents:", v.Contents)
	case *real.Retriver:
		fmt.Println("UserAgent:", v.UserAgent)
	}
	fmt.Println()
}
