package main

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/prometheus/common/log"
	"sync"
	"time"
)

type A struct {
	Age int
}

var a = sync.Pool{
	New: func() interface{} {
		return &A{}
	},
}

func main() {
	fmt.Println(1)
	totalChan := make(chan int, 10)
	closeChan := make(chan struct{}, 1)
	totalChan <- 0
	closeChan <- struct{}{}
	totalWaitGroup := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		totalWaitGroup.Add(1)
		go func(i int) {
			defer totalWaitGroup.Done()
			for {
				select {
				case total, ok := <-totalChan:
					if total >= 10 {
						select {
						case _, ok1 := <-closeChan:
							if ok1 == false {
								close(closeChan)
							} else {
								if !ok {
									close(totalChan)
								}
							}
						}
						goto end
					}
					fmt.Println(fmt.Sprintf("groutiune %d total:%d", i, total))
					totalChan <- total + 1
				}
				time.Sleep(10 * time.Millisecond)
				continue
			end:
				fmt.Println(fmt.Sprintf("groutiune %d end", i))
				break
			}
		}(i)
	}
	totalWaitGroup.Wait()
}

func handleError(err error) {
	log.Errorf(err.Error())
	//fmt.Println(log.Errorf(err.Error()))
	return
}

func AA(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.Wrapf(errors.New("b is zero"), "b id error,%d", b)
	}
	return a / b, nil
}
