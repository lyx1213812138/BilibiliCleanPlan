package main

import (
	"fmt"
)

type T struct {
	A int
	B string
}

func aaa() {
	t := []T{{1, "a"}, {2, "b"}}
	fmt.Println(t)
	for i, v := range t {
		fmt.Printf("%p %p\n", &v, &t[i])
		v.A = 3
	}
	fmt.Printf("%p\n", &t[0])
}
