package main

import (
	"runtime"
	"sync"
)

// Task 任务接口。
type Task interface {
	// PoolExecute
	// 参数为自定义参数
	PoolExecute(interface{})
}

// Pool 协程池。
type Pool struct {
	Size        int
	TaskChannel chan interface{} // 任务队列
	fc          func(interface{})
	wg          *sync.WaitGroup
}

func NewPool(cap ...int) *Pool {
	// 获取 worker 数量
	p := &Pool{
		TaskChannel: make(chan interface{}),
		wg:          &sync.WaitGroup{},
	}
	if len(cap) > 0 {
		p.Size = cap[0]
	}
	if p.Size == 0 {
		p.Size = runtime.NumCPU()
	}
	return p
}

func (p *Pool) AddFunc(f func(interface{})) {
	p.wg.Add(1)
	defer p.wg.Done()
	p.fc = f
}

func (p *Pool) Wait() {
	// 创建指定数量 worker 从任务队列取出任务执行。
	for i := 0; i < p.Size; i++ {
		p.wg.Add(1)
		go func() {
			defer p.wg.Done()
			for task := range p.TaskChannel {
				p.fc(task)
			}
		}()
	}
	p.wg.Wait()
}
