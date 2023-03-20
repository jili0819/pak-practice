package main

import (
	"bytes"
	"fmt"
)

// 模板方法模式是一种行为设计模式，它在超类中定义了一个算法的框架，允许子类在不修改结构的情况下重写算法的特定步骤。
// 由于GO语言没有继承的语法，模板方法又是依赖继承实现的设计模式，因此GO语言实现模板方法比较困难
// GO语言支持隐式内嵌字段“继承”其他结构体的字段与方法，但是这个并不是真正意义上的继承语法
// 外层结构重写隐式字段中的算法特定步骤后，无法动态绑定到“继承”过来的算法的框架方法调用中，因此不能实现模板方法模式的语义。

// 将多个算法特定步骤组合成一个接口；
// 基类隐式内嵌算法步骤接口，同时调用算法步骤接口的各方法，实现算法的模板方法，此时基类内嵌的算法步骤接口并没有真正的处理行为；
// 子类隐式内嵌基类，并覆写算法步骤接口的方法；
// 通过工厂方法创建具体子类，并将自己的引用赋值给基类中算法步骤接口字段；
// 以演员装扮为例，演员的装扮是分为化妆，穿衣，配饰三步骤，三个步骤又根据不同角色的演员有所差别
// 因此演员基类实现装扮的模板方法，对于化妆，穿衣，配饰的三个步骤，在子类演员中具体实现，子类具体演员分为，男演员、女演员和儿童演员；

type (
	// IActor 演员接口
	IActor interface {
		DressUp() string // 装扮
	}

	// dressBehavior 装扮的多个行为，这里多个行为是私有的，通过DressUp模版方法调用
	dressBehavior interface {
		makeUp() string // 化妆
		clothe() string // 穿衣
		wear() string   // 配饰
	}
	// BaseActor 演员基类
	BaseActor struct {
		roleName      string // 扮演角色
		dressBehavior        // 装扮行为
	}
)

// DressUp 统一实现演员接口的DressUp模版方法，装扮过程通过不同装扮行为进行扩展
func (b *BaseActor) DressUp() string {
	buf := bytes.Buffer{}
	buf.WriteString(fmt.Sprintf("扮演%s的", b.roleName))
	buf.WriteString(b.makeUp())
	buf.WriteString(b.clothe())
	buf.WriteString(b.wear())
	return buf.String()
}

// womanActor 扩展装扮行为的女演员
type womanActor struct {
	BaseActor
}

// NewWomanActor 指定角色创建女演员
func NewWomanActor(roleName string) *womanActor {
	actor := new(womanActor)    // 创建女演员
	actor.roleName = roleName   // 设置角色
	actor.dressBehavior = actor // 将女演员实现的扩展装扮行为，设置给自己的装扮行为接口
	return actor
}

// 化妆
func (w *womanActor) makeUp() string {
	return "女演员涂着口红，画着眉毛；"
}

// 穿衣
func (w *womanActor) clothe() string {
	return "穿着连衣裙；"
}

// 配饰
func (w *womanActor) wear() string {
	return "带着耳环，手拎着包；"
}

// manActor 扩展装扮行为的男演员
type manActor struct {
	BaseActor
}

func NewManActor(roleName string) *manActor {
	actor := new(manActor)
	actor.roleName = roleName
	actor.dressBehavior = actor // 将男演员实现的扩展装扮行为，设置给自己的装扮行为接口
	return actor
}

func (m *manActor) makeUp() string {
	return "男演员刮净胡子，抹上发胶；"
}

func (m *manActor) clothe() string {
	return "穿着一身西装；"
}

func (m *manActor) wear() string {
	return "带上手表，抽着烟；"
}

// NewChildActor 扩展装扮行为的儿童演员
type childActor struct {
	BaseActor
}

func NewChildActor(roleName string) *childActor {
	actor := new(childActor)
	actor.roleName = roleName
	actor.dressBehavior = actor // 将儿童演员实现的扩展装扮行为，设置给自己的装扮行为接口
	return actor
}

func (c *childActor) makeUp() string {
	return "儿童演员抹上红脸蛋；"
}

func (c *childActor) clothe() string {
	return "穿着一身童装；"
}

func (c *childActor) wear() string {
	return "手里拿着一串糖葫芦；"
}

func main() {
	showActors(NewWomanActor("妈妈"), NewManActor("爸爸"), NewChildActor("孩子"))
}

// showActors 显示演员的装扮信息
func showActors(actors ...IActor) {
	for _, actor := range actors {
		fmt.Println(actor.DressUp())
	}
}
