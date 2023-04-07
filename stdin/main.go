package main

import (
	"bufio"
	"fmt"
	"os"
)

// 获取输入
func main() {
	// 方式一：
	// Scan函数会识别空格左右的内容，哪怕换行符存在也不会影响Scan对内容的获取。
	// Scanln函数会识别空格左右的内容，但是一旦遇到换行符就会立即结束，不论后续还是否存在需要输入的内容
	// 都不能获取带空格的字符串

	// 方式二：使用buﬁo包里带缓冲的Reader实现带空格字符串的输入
	for {
		inputReader := bufio.NewReader(os.Stdin)
		fmt.Println("Please input your name: ")
		//input, err = inputReader.ReadString('\n')
		input, err := inputReader.ReadString('\n')
		if err == nil {
			fmt.Printf("Your name is: %s\n", input)
		}
	}

}
