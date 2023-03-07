package main

import "fmt"

func main() {

	fmt.Println(func5(1))
	fmt.Println(func6(1))
}

func func5(i int) int {
	k := i
	defer func() {
		k++
	}()
	return k
}

func func6(i int) int {
	k := i
	defer func(k int) {
		k++
	}(k)
	return k
}
