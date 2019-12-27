package Service

import "iris/config"

var Kafka = &kafka{}

type kafka struct{}

func (k *kafka) Send(topic string, payload string) error {
	return config.KafkaProducer.Send(topic, []byte(payload))
}
