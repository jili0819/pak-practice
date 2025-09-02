package base

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"go.uber.org/zap"
	"time"
)

type BaseProducer struct {
	producer *kafka.Producer
	callback func(partition kafka.TopicPartition, msg []byte)
}

func NewBaseProducer(params *kafka.ConfigMap) (producer *BaseProducer, err error) {
	if params == nil {
		params = getDefaultParams()
	}
	var p *kafka.Producer
	if p, err = kafka.NewProducer(params); err != nil {
		fmt.Println("pro")
		return
	}
	return &BaseProducer{producer: p}, nil
}

func (base *BaseProducer) Produce(topic string, messages [][]byte, partition int32) {
	// 默认任意分区
	if partition == 0 {
		partition = kafka.PartitionAny
	}

	go func() {
		for e := range base.producer.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				m := ev
				if ev.TopicPartition.Error != nil {
					fmt.Println("kafka生产失败 Failed to write access log entry", zap.Error(ev.TopicPartition.Error), zap.String("error", ev.TopicPartition.Error.Error()))
				} else {
					fmt.Println("kafka生产数据", zap.Any("topic", m.String()))
				}
			}
		}
	}()

	// Produce messages to topic (asynchronously)
	for _, msg := range messages {
		msgNew := &kafka.Message{
			TopicPartition: kafka.TopicPartition{
				Topic:     &topic,
				Partition: partition,
			},
			Value: msg,
		}
		err := base.producer.Produce(msgNew, nil)
		if err != nil {
			fmt.Println("kafka生产数据失败", zap.String("msg", string(msg)))
		}
		time.Sleep(time.Duration(1) * time.Millisecond)
	}

	// Wait for message deliveries before shutting down
	base.producer.Flush(15 * 1000)
}

func getDefaultParams() *kafka.ConfigMap {
	return &kafka.ConfigMap{
		"bootstrap.servers": "localhost:29092", // kafka地址
	}
}
