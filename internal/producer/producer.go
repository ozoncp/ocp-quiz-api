package producer

import (
	"github.com/Shopify/sarama"
	"github.com/rs/zerolog/log"
)

type Producer interface {
	Send(msg ...EventMsg)
}

// NewProducer Returns new kafka producer
func NewProducer(topic string, kafkaProducer sarama.SyncProducer) Producer {
	p := &producer{topic: topic, kafkaProducer: kafkaProducer}
	return p
}

type producer struct {
	topic         string
	kafkaProducer sarama.SyncProducer
}

// Send sends a batch of message to Kafka broker
func (p *producer) Send(msgs ...EventMsg) {
	if len(msgs) == 0 {
		return
	}
	prepared := make([]*sarama.ProducerMessage, len(msgs))
	for i, m := range msgs {
		prepared[i] = &sarama.ProducerMessage{
			Topic:     p.topic,
			Partition: -1,
			Value:     m,
		}

		log.Info().Msgf("%v", m)
	}
	err := p.kafkaProducer.SendMessages(prepared)
	if err != nil {
		log.Error().Msgf("failed to send messages to Kafka: %v", err)
	}
}
