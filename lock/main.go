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
func lock01() {
	ch := make(chan bool)
	lock := NewRedisLock(context.Background(), nil, "keys", "values")
	lock.Lock()
	defer lock.UnLock()
	fmt.Println("执行业务开始")
	time.Sleep(5 * time.Second)
	fmt.Println("执行业务结束")
	ch <- true
	close(ch)
}
