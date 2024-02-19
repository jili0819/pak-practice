package main

import (
	"fmt"
	"time"
)

type sliceT[T int | float64] []T                  // 切片范型
type mapT[K int | string, V int | string] map[K]V // map范型
type chanT[T int | string | A] chan T             // 通道范型
type structT[T int | string] struct {             // 结构体范型
	Name    string
	Content T
}
type funcT[T int | string] func(T) // 函数范型
type A struct {
	Name string
}

func main() {
	// 切片范型使用
	var aa sliceT[int] = []int{1, 2, 3}
	fmt.Println(aa)
	var bb sliceT[float64] = []float64{1.1, 2.2, 3.3}
	fmt.Println(bb)
	// map范型使用
	var cc mapT[int, string] = map[int]string{1: "a", 2: "b"}
	fmt.Println(cc)
	var dd mapT[int, int] = map[int]int{1: 1, 2: 2}
	fmt.Println(dd)
	var ee mapT[string, int] = map[string]int{"a": 1, "b": 2}
	fmt.Println(ee)
	var ff mapT[string, string] = map[string]string{"a": "a", "b": "b"}
	fmt.Println(ff)
	// 通道范型使用
	var gg chanT[int] = make(chan int)
	go func() {
		fmt.Println("<-gg:", <-gg)
	}()
	gg <- 1
	var hh chanT[string] = make(chan string)
	go func() {
		fmt.Println("<-hh:", <-hh)
	}()
	hh <- "hello"
	var ii chanT[A] = make(chan A)
	go func() {
		fmt.Println("<-ii:", <-ii)
	}()
	ii <- A{Name: "A"}
	// 结构体范型使用
	var jj = structT[int]{Name: "jj", Content: 1}
	fmt.Println(jj)
	var kk = structT[string]{Name: "kk", Content: "kk"}
	fmt.Println(kk)
	// 函数范型使用
	var ll = funcT[int](func(a int) {
		fmt.Println("ll:", a)
	})
	go func() {
		ll(1)
	}()
	var mm = funcT[string](func(a string) {
		fmt.Println("mm:", a)
	})
	go func() {
		mm("hello")
	}()

	time.Sleep(10 * time.Second)
}
