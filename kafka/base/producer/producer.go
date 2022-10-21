package base

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"math/rand"
	"time"
)

var producer *kafka.Producer

func main() {
	ch1 := make(chan bool)
	// 开多个协程一直输入Kafka信息
	go func() {
		topic := "purchases1"
		users := [...]string{"eabara", "jsmith", "sgarcia", "jbernard", "htanaka", "awalther"}
		items := [...]string{"book", "alarm clock", "t-shirts", "gift card", "batteries"}
		for {
			key := users[rand.Intn(len(users))]
			data := items[rand.Intn(len(items))]
			if err := producer.Produce(&kafka.Message{
				TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
				Key:            []byte(key),
				Value:          []byte(data),
			}, nil); err != nil {
				fmt.Println("producer err:", err)
			}
			time.Sleep(100 * time.Millisecond)
		}
	}()
	<-ch1
}

func init() {
	var err error
	if producer, err = kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
	}); err != nil {
		fmt.Println("pro")
		return
	}
}
