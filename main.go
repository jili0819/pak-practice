package main

import (
	"fmt"
	"github.com/shopspring/decimal"
)

const a = 1000

func main() {

	//var a time.Time

	b := int(decimal.NewFromFloat32(9.9).Mul(decimal.NewFromFloat32(100)).IntPart())
	fmt.Println(b)
}
