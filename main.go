package main

import "fmt"

func main() {
	a := make(map[string]A)
	a["a"] = A{
		Age: 10,
	}
	c := a["b"]
	fmt.Println(fmt.Sprintf("%", c))
}

type A struct {
	Age int
}
