package base

import "github.com/confluentinc/confluent-kafka-go/kafka"

type BaseConsumer struct {
	BaseConsumer *kafka.Consumer
	Callback     func(partition kafka.TopicPartition, msg []byte)
}

func NewBaseConsumer(params *kafka.ConfigMap, callBack func(partition kafka.TopicPartition, msg []byte)) *BaseConsumer {
	c, _ := kafka.NewConsumer(params)
	return &BaseConsumer{BaseConsumer: c, Callback: callBack}
}
