package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println(time.Now().Truncate(time.Minute).Unix())

}

func GG() (data []int) {
	return nil
}
