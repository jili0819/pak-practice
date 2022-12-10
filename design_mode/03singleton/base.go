package main

import (
	"fmt"
	"sync"
)

// 单例模式

type singleton struct{}

func (s *singleton) Name() {
	fmt.Println("singleton.Name()")
}

var ins *singleton
var once sync.Once

func GetIns() *singleton {
	// 懒汉式单例模式
	once.Do(func() {
		ins = &singleton{}
	})
	// 饿汉式模式
	// 直接方法之外，全局初始化
	return ins
}

func main() {
	instance := GetIns()
	instance.Name()
}
