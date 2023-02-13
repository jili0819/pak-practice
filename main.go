package main

import (
	"fmt"
	"time"
)

const a = 1000

func main() {
	a := []int{0, 1, 2, 3, 4, 5}
	b := a[0:3]
	fmt.Println(len(a), cap(a))
	fmt.Println(len(b), cap(b))
	b[0] = 1
	fmt.Println(a, b)
	return

	chans := make(chan int, 100)
	closeChan := make(chan bool)
	go test01(chans, closeChan)
	i := 0
	for i < 10 {
		i++
		go test02(chans, i)
	}
	go func() {
		time.Sleep(3 * time.Second)
		close(closeChan)
	}()
	time.Sleep(5 * time.Second)
}

func test01(chans chan int, closeChan chan bool) {
	for {
		select {
		case _, ok := <-closeChan:
			if !ok {
				close(chans)
				goto sendEnd
			}
		default:
			chans <- time.Now().Second()
		}
		time.Sleep(1 * time.Second)
	}
sendEnd:
	fmt.Println("test01 close func")
}

func test02(chans chan int, index int) {
	for {
		result, ok := <-chans
		if !ok {
			// 已关闭
			break
		}
		fmt.Println(index, result)
	}
	fmt.Println("test02 close func")
}
