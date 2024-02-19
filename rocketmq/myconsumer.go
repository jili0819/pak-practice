package rocketmq

import (
	"context"
	"fmt"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/rlog"
)

func main() {
	c, _ := rocketmq.NewPushConsumer(
		consumer.WithNameServer([]string{"localhost:9876"}),
		consumer.WithConsumerModel(consumer.Clustering),
		consumer.WithGroupName("GID_XXXXXX"),
	)
	_ = c.Subscribe("test", consumer.MessageSelector{}, func(ctx context.Context,
		msgs ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
		rlog.Info("Subscribe Callback", map[string]interface{}{"msgs": msgs})
		for _, msg := range msgs {
			fmt.Println(msg.MsgId)
			return consumer.ConsumeRetryLater, nil
		}
		return consumer.ConsumeSuccess, nil
	})
}
