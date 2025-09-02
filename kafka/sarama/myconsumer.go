package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/IBM/sarama"
	"github.com/jili/pkg-practice/kafka/sarama/base"
	"github.com/jili/pkg-practice/kafka/sarama/types"
	"go.uber.org/zap"
	"log"
	"os"
	"os/signal"
	"sync"
	"time"
)

type MyConsumerCallBack func(message *types.MyConsumerInfo)

type MyConsumer struct {
	ready        chan bool
	ConsumerCall MyConsumerCallBack
}

func (m *MyConsumer) Setup(session sarama.ConsumerGroupSession) error {
	close(m.ready)
	return nil
}

func (m *MyConsumer) Cleanup(session sarama.ConsumerGroupSession) error {
	return nil
}

func (m *MyConsumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		time.Sleep(1 * time.Second)
		fmt.Println(fmt.Sprintf("--%s---重写父级消费方法------", session.GenerationID()))
		m.callBack(msg)
		// 手动提交偏移量（标记为已消费）
		session.MarkMessage(msg, "")
		session.Commit()
	}
	return nil
}

func NewMyConsumer(callBack MyConsumerCallBack) *MyConsumer {
	c := &MyConsumer{ConsumerCall: callBack}
	return c
}

func (c *MyConsumer) callBack(message *sarama.ConsumerMessage) {
	// todo 自定义方法1,msg string->json->struct{}
	fmt.Println(fmt.Sprintf("callBack topic:%s,消息所在partition：%d,offset;%d", message.Topic, message.Partition, message.Offset))
	var temp types.MyConsumerInfo
	if err := json.Unmarshal(message.Value, &temp); err != nil {
		fmt.Println("err:", err)
		return
	}
	c.ConsumerCall(&temp)
}

func main() {
	// 组一
	client := base.NewConsumerGroup([]string{"localhost:29092"}, "group")
	ctx, cancel := context.WithCancel(context.Background())
	wg := &sync.WaitGroup{}
	for i := 0; i < 2; i++ {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			handle := NewMyConsumer(func(info *types.MyConsumerInfo) {
				fmt.Println("自定义消费方法:start", time.Now().Format("2006-01-02 15:04:05"))
				fmt.Println("自定义消费方法:", info.Name, time.Now().Format("2006-01-02 15:04:05"))
				fmt.Println("自定义消费方法:end", time.Now().Format("2006-01-02 15:04:05"))
			})
			for {
				err := client.Consume(ctx, []string{"purchases"}, handle)
				if err != nil {
					log.Fatal("kafka消费错误", zap.Error(err), zap.String("err", err.Error()))
					return
				}
				// check if context was cancelled, signaling that the consumer should stop
				if ctx.Err() != nil {
					return
				}
			}
		}(i)
	}
	log.Println("Consumer up and running!")
	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, os.Interrupt)
	<-sigterm
	cancel()
	wg.Wait()
	if err := client.Close(); err != nil {
		log.Fatalf("Error closing consumer: %v", err)
	}
}
