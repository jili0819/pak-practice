package main

import (
	"fmt"
	_interface "github.com/jili/pkg-practice/interface"
)

func main() {

	// 只对外暴漏接口，需要获取单独写方法获取内部参数
	client := _interface.NewApi()
	fmt.Println(client.Add(1, 2))
}
