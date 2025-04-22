package base

import "github.com/confluentinc/confluent-kafka-go/v2/kafka"

type BaseConsumer struct {
	Base     *kafka.Consumer
	Callback func(partition kafka.TopicPartition, msg []byte)
}

func NewBaseConsumer(params *kafka.ConfigMap, callBack func(partition kafka.TopicPartition, msg []byte)) *BaseConsumer {
	c, _ := kafka.NewConsumer(params)
	return &BaseConsumer{Base: c, Callback: callBack}
}
