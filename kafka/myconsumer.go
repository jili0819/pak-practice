package main

import (
	"encoding/json"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/jili/pkg-practice/kafka/base"
	"github.com/jili/pkg-practice/kafka/types"
	"go.uber.org/zap"
	"log"
	"os"
	"time"
)

type MyConsumerCallBack func(*types.MyConsumerInfo)

type MyConsumer struct {
	Topics       []string
	Consumer     *base.BaseConsumer
	ConsumerCall MyConsumerCallBack
}

func NewMyConsumer(topics []string, groupId, offset string, callBack MyConsumerCallBack) *MyConsumer {
	c := &MyConsumer{Topics: topics, ConsumerCall: callBack}
	c.Consumer = base.NewBaseConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
		"group.id":          groupId,
		"auto.offset.reset": offset,
	}, c.callBack)
	return c
}

func (c *MyConsumer) callBack(partition kafka.TopicPartition, msg []byte) {
	// todo 自定义方法,msg string->json->struct{}

	fmt.Println("消息所在partition：", *partition.Topic, partition.Partition)
	var temp types.MyConsumerInfo
	if err := json.Unmarshal(msg, &temp); err != nil {
		fmt.Println("err:", err)
		return
	}
	c.ConsumerCall(&temp)
}

func main() {
	aa := NewMyConsumer([]string{"purchases2"}, "test", "earliest", func(info *types.MyConsumerInfo) {
		fmt.Println("------MyConsumer--------:", info.Name, time.Now().Format("2006-01-02 15:04:05"))
	})

	for i := 0; i < 10; i++ {
		go func(index int) {
			if err := aa.Consumer.Base.SubscribeTopics(aa.Topics, nil); err != nil {
				log.Fatalf("kafka消费错误", zap.Error(err), zap.String("err", err.Error()))
			}
			for {
				//fmt.Println(aa.Consumer.BaseConsumer.ReadMessage(-1))

				ev := aa.Consumer.Base.Poll(100)
				if ev == nil {
					continue
				}

				switch e := ev.(type) {
				case *kafka.Message:
					fmt.Println("消费者groupId:", index)
					aa.Consumer.Callback(e.TopicPartition, e.Value)
				case kafka.Error:
					// Errors should generally be considered
					// informational, the client will try to
					// automatically recover.
					// But in this example we choose to terminate
					// the application if all brokers are down.
					fmt.Fprintf(os.Stderr, "%% Error: %v: %v\n", e.Code(), e)
					break
				default:
					fmt.Printf("Ignored %v\n", e)
				}
				// aa.Consumer.Callback(msg.TopicPartition, msg.Value)
				/*ev := <-aa.Consumer.BaseConsumer.Events()
				switch e := ev.(type) {
				case *kafka.Message:
					aa.Consumer.Callback(e.TopicPartition, e.Value)
				}*/
			}
		}(i)
	}

	time.Sleep(1 * time.Hour)
}
