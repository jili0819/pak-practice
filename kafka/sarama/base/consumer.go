package base

import "github.com/IBM/sarama"

func NewConsumerGroup(adds []string, groupId string) sarama.ConsumerGroup {
	config := sarama.NewConfig()
	config.Version = sarama.DefaultVersion
	// 手动提交
	config.Consumer.Offsets.AutoCommit.Enable = false
	config.Consumer.Offsets.Initial = sarama.OffsetOldest
	c, err := sarama.NewConsumerGroup(adds, groupId, config)
	if err != nil {
		panic(err)
	}
	return c
}
