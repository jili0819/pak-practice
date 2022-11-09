package main

import (
	"fmt"
	"github.com/panjf2000/ants/v2"
	"time"
)

func main() {
	cronCh := make(chan uint)
	go PushToEs(cronCh)
	cronCh <- 1
	time.Sleep(2 * time.Second)
	cronCh <- 2
	time.Sleep(2 * time.Second)
	cronCh <- 3
	time.Sleep(2 * time.Second)
	cronCh <- 4
	time.Sleep(2 * time.Second)
	close(cronCh)
	time.Sleep(1 * time.Hour)
}
func PushToEs(cronCh chan uint) {
	p, _ := ants.NewPool(10)
	p.Running()
	defer p.Release()
	//ctx := context.Background()
	for {
		select {
		case v, ok := <-cronCh:
			if !ok {
				goto ExitPush
			}
			if err := ants.Submit(func() {
				fmt.Println(v)
			}); err != nil {
				fmt.Println(err)
			}
		}
	}
ExitPush:
	fmt.Println("end")
	return
}
