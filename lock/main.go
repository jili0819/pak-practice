package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"log"
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
	key := "key01"
	var g errgroup.Group
	ch := make(chan bool)

	// 业务处理
	g.Go(func() error {
		fmt.Println("执行业务开始")
		time.Sleep(5 * time.Second)
		fmt.Println("执行业务结束")
		ch <- true
		close(ch)
		return nil
	})
	// 锁续期
	g.Go(func() error {
		ticker := time.NewTicker(time.Second * 1)
		for {
			select {
			case <-ticker.C:
				fmt.Println("定时续期分布式锁有效期")
				if err := LockExpire(context.Background(), key); err != nil {
					log.Println(err.Error())
					return err
				}
			case <-ch:
				fmt.Println("执行业务结束,关闭续期,删除锁")
				ticker.Stop()
				if err := UnLock(context.Background(), key); err != nil {
					log.Println(err.Error())
					return err
				}
				return nil
			}
			time.Sleep(10 * time.Millisecond)
		}
	})
	if err := g.Wait(); err != nil {
		fmt.Println(err)
		return
	}
}
