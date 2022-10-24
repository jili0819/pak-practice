package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/singleflight"
	"time"
)

func main() {
	var sg singleflight.Group
	var b *A
	result, _, _ := sg.Do("111", func() (interface{}, error) {
		return &A{
			Name: "w",
		}, nil
	})
	b = result.(*A)
	fmt.Println(b)

	/*ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	fmt.Println(getData(ctx, "keys"))*/

}

type A struct {
	Name string `json:"name"`
}

// 超时控制
func getData(ctx context.Context, key string) (interface{}, error) {
	var sg singleflight.Group
	result := sg.DoChan(key, func() (interface{}, error) {
		time.Sleep(4 * time.Second)
		return "11111", nil
	})
	// 调用的时候传入一个含超时的 context 即可，执行时就会返回超时错误
	select {
	case r := <-result:
		return r.Val, r.Err
	case <-ctx.Done():
		return "", ctx.Err()
	}
}
