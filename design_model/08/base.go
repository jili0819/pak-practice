package main

import (
	"fmt"
)

// 观察者模式是一种行为设计模式，允许你定义一种订阅机制，可在对象事件发生时通知多个 “观察” 该对象的其他对象。
// 观察者模式提供了一种作用于任何实现了订阅者接口的对象的机制，可对其事件进行订阅和取消订阅。
// 观察者模式是最常用的模式之一，是事件总线，分布式消息中间件等各种事件机制的原始理论基础，常用于解耦多对一的对象依赖关系；

// 增强的实现功能包括：
// 当被观察者通过异步实现通知多个观察者时就相当于单进程实例的消息总线；
// 同时还可以根据业务需要，将被观察者所有数据状态变更进行分类为不同的主题，观察者通过不同主题进行订阅；
// 同一个主题又可分为增加，删除，修改事件行为；
// 每个主题可以实现一个线程池，多个主题通过不同的线程池进行处理隔离，线程池可以设置并发线程大小、缓冲区大小及调度策略，比如先进先出，优先级等策略；
// 观察者处理事件时有可能出现异常，所以也可以注册异常处理函数，异常处理也可以通过异常类型进行分类；
// 根据业务需求也可以实现通知异常重试，延迟通知等功能；

// 例子：信用卡业务消息提醒可通过观察者模式实现，业务消息包括日常消费，出账单，账单逾期
// 消息提醒包括短信、邮件及电话，根据不同业务的场景会采用不同的消息提醒方式或者多种消息提醒方式
// 这里信用卡相当于被观察者，观察者相当于不同的通知方式；日常消费通过短信通知，出账单通过邮件通知，账单逾期三种方式都会进行通知；

type MessageNotifyType int

const (
	MessageNotify MessageNotifyType = iota + 1 // 短信通知
	EmailNotify                                // 邮件通知
	TelNotify                                  // 手机通知
)

type (
	// Subscriber 订阅者接口
	Subscriber interface {
		Name() string            //订阅者名称
		Type() MessageNotifyType // 通知类型
		Update(message string)   //订阅更新方法
	}

	// shortMessage 信用卡消息短信订阅者
	shortMessage struct{}

	// email 信用卡消息邮箱订阅者
	email struct{}

	// telephone 信用卡消息电话订阅者
	telephone struct{}
)

func (s *shortMessage) Name() string {
	return "短信通知"
}

func (s *shortMessage) Type() MessageNotifyType {
	return MessageNotify
}

func (s *shortMessage) Update(message string) {
	fmt.Printf("通过【%s】发送消息:%s\n", s.Name(), message)
}

func (e *email) Name() string {
	return "电子邮件"
}

func (e *email) Type() MessageNotifyType {
	return EmailNotify
}

func (e *email) Update(message string) {
	fmt.Printf("通过【%s】发送消息:%s\n", e.Name(), message)
}

func (t *telephone) Name() string {
	return "电话"
}

func (t *telephone) Type() MessageNotifyType {
	return TelNotify
}

func (t *telephone) Update(message string) {
	fmt.Printf("通过【%s】告知:%s\n", t.Name(), message)
}

// 信用卡业务信息

// MsgType 信用卡消息类型
type MsgType int

const (
	ConsumeType MsgType = iota + 1 // 消费消息类型
	BillType                       // 账单消息类型
	ExpireType                     // 逾期消息类型
)

// CreditCard 信用卡
type CreditCard struct {
	holder          string                   // 持卡人
	consumeSum      float32                  // 消费总金额
	subscriberGroup map[MsgType][]Subscriber // 根据消息类型分组订阅者
}

// NewCreditCard 指定持卡人创建信用卡
func NewCreditCard(holder string) *CreditCard {
	return &CreditCard{
		holder:          holder,
		subscriberGroup: make(map[MsgType][]Subscriber),
	}
}

// Subscribe 支持订阅多种消息类型
func (c *CreditCard) Subscribe(subscriber Subscriber, msgTypes ...MsgType) {
	for _, msgType := range msgTypes {
		c.subscriberGroup[msgType] = append(c.subscriberGroup[msgType], subscriber)
	}
}

// Unsubscribe 解除订阅多种消息类型
func (c *CreditCard) Unsubscribe(subscriber Subscriber, msgTypes ...MsgType) {
	for _, msgType := range msgTypes {
		if subs, ok := c.subscriberGroup[msgType]; ok {
			c.subscriberGroup[msgType] = removeSubscriber(subs, subscriber)
		}
	}
}

func removeSubscriber(subscribers []Subscriber, toRemove Subscriber) []Subscriber {
	length := len(subscribers)
	for i, subscriber := range subscribers {
		if toRemove.Name() == subscriber.Name() {
			subscribers[length-1], subscribers[i] = subscribers[i], subscribers[length-1]
			return subscribers[:length-1]
		}
	}
	return subscribers
}

// Consume 信用卡消费
func (c *CreditCard) Consume(money float32) {
	c.consumeSum += money
	c.notify(ConsumeType, fmt.Sprintf("尊敬的持卡人%s,您当前消费%.2f元;", c.holder, money))
}

// SendBill 发送信用卡账单
func (c *CreditCard) SendBill() {
	c.notify(BillType, fmt.Sprintf("尊敬的持卡人%s,您本月账单已出，消费总额%.2f元;", c.holder, c.consumeSum))
}

// Expire 逾期通知
func (c *CreditCard) Expire() {
	c.notify(ExpireType, fmt.Sprintf("尊敬的持卡人%s,您本月账单已逾期，请及时还款，总额%.2f元;", c.holder, c.consumeSum))
}

// notify 根据消息类型通知订阅者
func (c *CreditCard) notify(msgType MsgType, message string) {
	if subs, ok := c.subscriberGroup[msgType]; ok {
		for _, sub := range subs {
			sub.Update(message)
		}
	}
}

func main() {
	// 创建张三的信用卡
	creditCard := NewCreditCard("张三")
	// 短信通知订阅信用卡消费及逾期消息
	creditCard.Subscribe(new(shortMessage), ConsumeType, ExpireType)
	// 电子邮件通知订阅信用卡账单及逾期消息
	creditCard.Subscribe(new(email), BillType, ExpireType)
	// 电话通知订阅信用卡逾期消息，同时逾期消息通过三种方式通知
	creditCard.Subscribe(new(telephone), ExpireType)

	creditCard.Consume(500.00) // 信用卡消费
	creditCard.Consume(800.00) // 信用卡消费
	creditCard.SendBill()      // 信用卡发送账单
	creditCard.Expire()        // 信用卡逾期

	// 信用卡逾期消息取消电子邮件及短信通知订阅
	creditCard.Unsubscribe(new(email), ExpireType)
	creditCard.Unsubscribe(new(shortMessage), ExpireType)
	creditCard.Consume(300.00) // 信用卡消费
	creditCard.Expire()        // 信用卡逾期
}
