package producer

import (
	"fmt"
	"github.com/Shopify/sarama"
)

type Producer struct {
	self sarama.SyncProducer
}

func (p *Producer) Init(addr []string) (err error) {
	config := sarama.NewConfig()
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.RequiredAcks = sarama.WaitForAll
    config.Producer.Return.Successes = true

	p.self, err = sarama.NewSyncProducer(addr, config)
	if err != nil {
		return err
	}

	return nil
}

func (p *Producer) SendMessage(topic string, msg string) (err error) {
	kafkaMsg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(msg),
	}

	partition, offset, err := p.self.SendMessage(kafkaMsg)
	if err != nil {
		return err
	}

	fmt.Printf("partition: %v, offset: %v", partition, offset)
	return nil
}
