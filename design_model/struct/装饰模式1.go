package main

import (
	"fmt"
	"time"
)

// 为函数类型设置别名提高代码可读性
type addFunc func(int, int) int

// 加法运算函数
func add(a, b int) int {
	return a + b
}

// 通过高阶函数在不侵入原有函数实现的前提下计算加法函数执行时间
func execTime(f addFunc) addFunc {
	return func(a, b int) int {
		start := time.Now()      // 起始时间
		c := f(a, b)             // 执行加法运算函数
		end := time.Since(start) // 函数执行完毕耗时
		fmt.Printf("--- 执行耗时: %v ---\n", end)
		return c // 返回计算结果
	}
}
func main() {
	a := 2
	b := 8
	// 通过修饰器调用加法函数，返回的是一个匿名函数
	decorator := execTime(add)
	// 执行修饰器返回函数
	c := decorator(a, b)
	fmt.Printf("%d + %d = %d\n", a, b, c)
}
