package main

import (
	"encoding/json"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/gofrs/uuid"
	"github.com/jili/pkg-practice/kafka/confluent/base"
	"github.com/jili/pkg-practice/kafka/confluent/types"
	"time"
)

func main() {
	aa, err := base.NewBaseProducer(nil)
	if err != nil {
		return
	}
	go func() {
		for {
			uuid, _ := uuid.NewV4()
			jsonstr, _ := json.Marshal(types.MyConsumerInfo{
				Name: uuid.String(),
			})
			aa.Produce("purchases2", [][]byte{jsonstr}, kafka.PartitionAny)
			time.Sleep(5 * time.Second)
		}
	}()
	time.Sleep(1 * time.Hour)
}
