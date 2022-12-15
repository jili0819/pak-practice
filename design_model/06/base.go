package main

import (
	"fmt"
	"time"
)

// 中介者模式是一种行为设计模式，能让你减少对象之间混乱无序的依赖关系。
// 该模式会限制对象之间的直接交互，迫使它们通过一个中介者对象进行合作，将网状依赖变为星状依赖。

// 中介者能使得程序更易于修改和扩展，而且能更方便地对独立的组件进行复用，因为它们不再依赖于很多其他的类。

// 中介者模式与观察者模式之间的区别是，中介者模式解决的是同类或者不同类的多个对象之间多对多的依赖关系，
// 观察者模式解决的是多个对象与一个对象之间的多对一的依赖关系。

// 例子：机场塔台调度系统是一个体现中介者模式的典型示例，假设是一个小机场，每次只能同时允许一架飞机起降，
// 每架靠近机场的飞机需要先与塔台沟通是否可以降落，如果没有空闲的跑道，需要在天空盘旋等待，
// 如果有飞机离港，等待的飞机会收到塔台的通知，按先后顺序降落；
// 这种方式，免去多架飞机同时到达机场需要相互沟通降落顺序的复杂性，减少多个飞机间的依赖关系，简化业务逻辑，从而降低系统出问题的风险。

type (
	// Aircraft 飞机接口
	Aircraft interface {
		ApproachAirport() // 抵达机场空域
		DepartAirport()   // 飞离机场
	}

	// airliner 客机
	airliner struct {
		name            string          // 客机型号
		status          int             // 1-起飞中 2-停飞中
		airportMediator AirportMediator // 机场调度
	}

	// helicopter 直升机
	helicopter struct {
		name            string          // 客机型号
		status          int             // 1-起飞中 2-停飞中
		airportMediator AirportMediator // 机场调度
	}

	// AirportMediator 机场调度中介者
	AirportMediator interface {
		CanLandAirport(aircraft Aircraft) bool // 确认是否可以降落
		NotifyWaitingAircraft()                // 通知等待降落或者起飞的其他飞机
		AddStartAirport(aircraft Aircraft)     // 添加起飞飞机
	}

	// ApproachTower 机场塔台
	ApproachTower struct {
		hasFreeAirstrip   bool
		waitingQueue      []Aircraft // 等待降落的飞机队列
		waitingStartQueue []Aircraft // 等待起飞的飞机队列
	}
)

// NewAirliner 根据指定型号及机场调度创建客机
func NewAirliner(name string, status int, mediator AirportMediator) *airliner {
	return &airliner{
		name:            name,
		status:          status,
		airportMediator: mediator,
	}
}

func (a *airliner) ApproachAirport() {
	if a.status != 1 {
		return
	}
	if !a.airportMediator.CanLandAirport(a) { // 请求塔台是否可以降落
		fmt.Printf("机场繁忙，正有其它飞机在降落或起飞，%s等待降落;\n", a.name)
		return
	}
	fmt.Printf("%s开始降落机场;\n", a.name)
	time.Sleep(2 * time.Second)
	fmt.Printf("%s降落机场;\n", a.name)
	a.airportMediator.AddStartAirport(a)
	a.airportMediator.NotifyWaitingAircraft() // 通知等待的其他飞机
}

func (a *airliner) DepartAirport() {
	if a.status != 2 {
		return
	}

	if !a.airportMediator.CanLandAirport(a) { // 请求塔台是否可以起飞
		fmt.Printf("机场繁忙，%s等待起飞;\n", a.name)
		a.airportMediator.AddStartAirport(a)
		return
	}
	fmt.Printf("%s开始起飞;\n", a.name)
	time.Sleep(2 * time.Second)
	fmt.Printf("%s起飞，离开机场;\n", a.name)
	a.airportMediator.NotifyWaitingAircraft() // 通知等待的其他飞机
}

// NewHelicopter 根据指定型号及机场调度创建直升机
func NewHelicopter(name string, status int, mediator AirportMediator) *helicopter {
	return &helicopter{
		name:            name,
		status:          status,
		airportMediator: mediator,
	}
}

func (h *helicopter) ApproachAirport() {
	if h.status != 1 {
		return
	}
	if !h.airportMediator.CanLandAirport(h) { // 请求塔台是否可以降落
		fmt.Printf("机场繁忙，正有其它飞机在降落或起飞，%s等待降落;\n", h.name)
		return
	}
	fmt.Printf("%s开始降落机场;\n", h.name)
	time.Sleep(2 * time.Second)
	fmt.Printf("%s降落机场;\n", h.name)
	h.airportMediator.AddStartAirport(h)
	h.airportMediator.NotifyWaitingAircraft() // 通知等待的其他飞机
}

func (h *helicopter) DepartAirport() {
	if h.status != 2 {
		return
	}
	if !h.airportMediator.CanLandAirport(h) { // 请求塔台是否可以起飞
		fmt.Printf("机场繁忙，%s等待起飞;\n", h.name)
		h.airportMediator.AddStartAirport(h)
		return
	}
	fmt.Printf("%s开始起飞---\n", h.name)
	time.Sleep(2 * time.Second)
	fmt.Printf("%s起飞，离开机场---\n", h.name)
	h.airportMediator.NotifyWaitingAircraft() // 通知其他等待降落的飞机
}

func (a *ApproachTower) CanLandAirport(aircraft Aircraft) bool {
	if a.hasFreeAirstrip {
		a.hasFreeAirstrip = false
		return true
	}
	// 没有空余的跑道，加入等待队列
	a.waitingQueue = append(a.waitingQueue, aircraft)
	return false
}

func (a *ApproachTower) AddStartAirport(aircraft Aircraft) {
	a.waitingStartQueue = append(a.waitingStartQueue, aircraft)
}

func (a *ApproachTower) NotifyWaitingAircraft() {
	if !a.hasFreeAirstrip {
		a.hasFreeAirstrip = true
	}
	if len(a.waitingQueue) > 0 {
		// 如果存在等待降落的飞机，通知第一个降落
		first := a.waitingQueue[0]
		if len(a.waitingQueue) > 1 {
			a.waitingQueue = a.waitingQueue[1:]
		} else {
			a.waitingQueue = nil
		}
		first.ApproachAirport()
	}
	if len(a.waitingStartQueue) > 0 {
		// 如果存在等待起飞的飞机，通知第一个起飞
		first := a.waitingStartQueue[0]
		if len(a.waitingStartQueue) > 1 {
			a.waitingStartQueue = a.waitingStartQueue[1:]
		} else {
			a.waitingStartQueue = nil
		}
		first.DepartAirport()
	}
}

func main() {

	// 创建机场调度塔台
	airportMediator := &ApproachTower{hasFreeAirstrip: true}
	// 创建C919客机
	c919Airliner := NewAirliner("C919", 1, airportMediator)
	// 创建米-26重型运输直升机
	m26Helicopter := NewHelicopter("米-26", 1, airportMediator)
	go c919Airliner.ApproachAirport()  // c919进港降落
	go m26Helicopter.ApproachAirport() // 米-26进港等待

	c919Airliner1 := NewAirliner("stop-C919", 2, airportMediator)
	// 创建米-26重型运输直升机
	m26Helicopter2 := NewHelicopter("stop-米-26", 2, airportMediator)

	go c919Airliner1.DepartAirport()  // c919飞离，等待的米-26进港降落
	go m26Helicopter2.DepartAirport() // 最后米-26飞离
	time.Sleep(1 * time.Hour)
}
