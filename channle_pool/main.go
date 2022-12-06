package main

import (
	"fmt"
	"github.com/panjf2000/ants/v2"
	"time"
)

func main() {
	cronCh := make(chan uint)
	go B(cronCh)
	var i uint
	for {
		cronCh <- i
		time.Sleep(3 * time.Millisecond)
		i++
	}
}

func B(cronCh chan uint) {
	p, _ := ants.NewPool(50)
	p.Running()
	defer p.Release()
	for {
		select {
		case v, ok := <-cronCh:
			if !ok {
				goto ExitPush
			}
			if err := p.Submit(func() {
				fmt.Println(v)
			}); err != nil {
				fmt.Println(err)
			}
		}
		//time.Sleep(1 * time.Millisecond)
	}
ExitPush:
}
