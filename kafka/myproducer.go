package main

import (
	"encoding/json"
	"github.com/gofrs/uuid"
	"github.com/jili/pkg-practice/kafka/base"
	"github.com/jili/pkg-practice/kafka/types"
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
			aa.Produce("purchases2", [][]byte{jsonstr}, 0)
			time.Sleep(10 * time.Millisecond)
		}
	}()
	time.Sleep(1 * time.Hour)
}
