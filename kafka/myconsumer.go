package main

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	base "github.com/jili/pkg-practice/kafka/base/customer"
	"go.uber.org/zap"
	"log"
)

type MyConsumerCallBack func(*MyConsumerInfo)

type MyConsumerInfo struct {
	Name string `json:"name"`
}

type MyConsumer struct {
	Topics       []string
	Consumer     *base.BaseConsumer
	ConsumerCall MyConsumerCallBack
}

func NewMyConsumer(topics []string, groupId, offset string, callBack MyConsumerCallBack) *MyConsumer {
	c := &MyConsumer{Topics: topics, ConsumerCall: callBack}
	c.Consumer = base.NewBaseConsumer(&kafka.ConfigMap{}, c.callBack)
	return c
}

func (c *MyConsumer) callBack(partition kafka.TopicPartition, msg []byte) {
	// todo 自定义方法,msg string->json->struct{}
	c.ConsumerCall(&MyConsumerInfo{})
}

func main() {
	aa := NewMyConsumer([]string{"1"}, "1", "12", func(info *MyConsumerInfo) {
		fmt.Println(info.Name)
	})

	go func() {
		if err := aa.Consumer.BaseConsumer.SubscribeTopics(aa.Topics, nil); err != nil {
			log.Fatalf("kafka消费错误", zap.Error(err), zap.String("err", err.Error()))
		}
		for {
			ev := <-aa.Consumer.BaseConsumer.Events()
			switch e := ev.(type) {
			case *kafka.Message:
				aa.Consumer.Callback(e.TopicPartition, e.Value)
			}
		}
	}()
}
