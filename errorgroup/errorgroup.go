package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"sync"
	"time"
)

func main() {
	a := []int64{1, 2, 3, 4, 5, 6, 7}
	g, _ := errgroup.WithContext(context.Background())
	g.Go(func() error {
		for index := range a {
			fmt.Println(index)
			time.Sleep(1 * time.Second)
		}
		return nil
	})
	g.Go(func() error {
		return nil
	})
	if err := g.Wait(); err != nil {
		fmt.Println(err)
		return
	}

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		time.Sleep(3 * time.Second)
		wg.Done()
		fmt.Println("主程序完成")
	}()
	go func() {
		wg.Wait()
		fmt.Println("主程序完成1")
	}()
	go func() {
		wg.Wait()
		fmt.Println("主程序完成2")
	}()
	go func() {
		wg.Wait()
		fmt.Println("主程序完成3")
	}()
	go func() {
		wg.Wait()
		fmt.Println("主程序完成4")
	}()

	time.Sleep(10 * time.Second)
}
