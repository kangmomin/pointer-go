package main

import "fmt"

type test struct {
	A int
	B string
}

func main() {
	var a *test
	var b test
	fmt.Println(a) // <nil>
	fmt.Println(b) // {0 }

	a = &b
	fmt.Println(a)  // &{0 }
	fmt.Println(*a) // {0 } 역참조(해당 메모리 주소에 직접 접근)

	a.B = "string"
	fmt.Println(b) // {0 string}
}
