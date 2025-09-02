package main

import (
	"encoding/json"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/jili/pkg-practice/kafka/confluent/base"
	"github.com/jili/pkg-practice/kafka/confluent/types"
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
	callBackFunc := c.callBack
	if groupId == "two" {
		callBackFunc = c.callBack1
	}
	c.Consumer = base.NewBaseConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:29092",
		"group.id":          groupId,
		"auto.offset.reset": offset,
	}, callBackFunc)
	return c
}

func (c *MyConsumer) callBack(partition kafka.TopicPartition, msg []byte) {
	// todo 自定义方法1,msg string->json->struct{}

	fmt.Println(fmt.Sprintf("callBack topic:%s,消息所在partition：%d", *partition.Topic, partition.Partition))
	var temp types.MyConsumerInfo
	if err := json.Unmarshal(msg, &temp); err != nil {
		fmt.Println("err:", err)
		return
	}
	c.ConsumerCall(&temp)
}

func (c *MyConsumer) callBack1(partition kafka.TopicPartition, msg []byte) {
	// todo 自定义方法,msg string->json->struct{}
	fmt.Println(fmt.Sprintf("callBack1 topic:%s,消息所在partition：%d", *partition.Topic, partition.Partition))
	var temp types.MyConsumerInfo
	if err := json.Unmarshal(msg, &temp); err != nil {
		fmt.Println("err:", err)
		return
	}
	c.ConsumerCall(&temp)
}

func main() {
	// 组一
	aa := NewMyConsumer([]string{"purchases2"}, "one", "earliest", func(info *types.MyConsumerInfo) {
		fmt.Println("自定义消费方法:start", time.Now().Format("2006-01-02 15:04:05"))
		fmt.Println("自定义消费方法:", info.Name, time.Now().Format("2006-01-02 15:04:05"))
		fmt.Println("自定义消费方法:end", time.Now().Format("2006-01-02 15:04:05"))

	})
	// 组二
	bb := NewMyConsumer([]string{"purchases2"}, "two", "earliest", func(info *types.MyConsumerInfo) {
		fmt.Println("自定义消费方法:start", time.Now().Format("2006-01-02 15:04:05"))
		fmt.Println("自定义消费方法：:", info.Name, time.Now().Format("2006-01-02 15:04:05"))
		fmt.Println("自定义消费方法:end", time.Now().Format("2006-01-02 15:04:05"))
	})

	for i := 0; i < 2; i++ {
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
					fmt.Println("消费者groupId:", "one")
					aa.Consumer.Callback(e.TopicPartition, e.Value)
					//aa.Consumer.Base.Commit()
				case kafka.Error:
					// Errors should generally be considered
					// informational, the client will try to
					// automatically recover.
					// But in this example we choose to terminate
					// the application if all brokers are down.
					fmt.Fprintf(os.Stderr, "%% Error: %v: %v\n", e.Code(), e)
					break
				case kafka.OffsetsCommitted:
					fmt.Printf("%v\n", e.String())
				default:
					fmt.Printf("default %v\n", e)
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

	for i := 0; i < 2; i++ {
		go func(index int) {
			if err := bb.Consumer.Base.SubscribeTopics(bb.Topics, nil); err != nil {
				log.Fatalf("kafka消费错误", zap.Error(err), zap.String("err", err.Error()))
			}
			for {
				//fmt.Println(aa.Consumer.BaseConsumer.ReadMessage(-1))

				ev := bb.Consumer.Base.Poll(100)
				if ev == nil {
					continue
				}

				switch e := ev.(type) {
				case *kafka.Message:
					fmt.Println("消费者groupId:", "two")
					bb.Consumer.Callback(e.TopicPartition, e.Value)
				case kafka.Error:
					// Errors should generally be considered
					// informational, the client will try to
					// automatically recover.
					// But in this example we choose to terminate
					// the application if all brokers are down.
					fmt.Fprintf(os.Stderr, "%% Error: %v: %v\n", e.Code(), e)
					break
				case kafka.OffsetsCommitted:
					fmt.Printf("%v\n", e.String())
				default:
					fmt.Printf("default %v\n", e)
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
