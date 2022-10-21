package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println(rand.Intn(4) + 1)
	fmt.Println(rand.Intn(2) + 1)
	fmt.Println(rand.Intn(4) + 1)
	fmt.Println(rand.Intn(2) + 1)
}
