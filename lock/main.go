package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var (
	cmx sync.Mutex
)

func main() {
	lock01()
	//lock02()

}

// err-group
func lock01() error {
	ch := make(chan bool)

	lock := new(context.Background(), nil, "keys", "values")
	if err := lock.Lock(); err != nil {
		return err
	}
	defer lock.UnLock()
	fmt.Println("执行业务开始")
	time.Sleep(5 * time.Second)
	fmt.Println("执行业务结束")
	ch <- true
	close(ch)
}
