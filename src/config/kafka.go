package config

import (
	"context"
	"errors"
	"github.com/Shopify/sarama"
	"sync"
	"time"
)

func InitKaProducer(c *ProducerConfig) (err error) {
	KafkaProducer, err = NewProducer(c, Log)
	return
}

func (p *Producer) Stop() {
	p.cancel()

	if err := p.client.Close(); err != nil {
		p.log.LogInfo("kafka producer close failed | brokers: %+v | error: %s", p.c.Brokers, err)
	}

	p.wg.Wait()
}

func (p *Producer) Send(topic string, payload []byte) error {
	select {
	case <-p.ctx.Done():
		return errors.New("producer is stoped")
	default:
		p.client.Input() <- &sarama.ProducerMessage{
			Topic: topic,
			Value: sarama.ByteEncoder(payload),
		}

		return nil
	}
}

func (p *Producer) logErr() {
	defer p.wg.Done()

	for {
		select {
		case <-p.ctx.Done():
			return
		case err := <-p.client.Errors():
			p.log.LogInfo("kafka producer revice error | brokers: %+v | error: %s", p.c.Brokers, err)
		}
	}
}

func NewProducer(c *ProducerConfig, logger *Logger) (producer *Producer, err error) {
	config := sarama.NewConfig()
	config.Net.KeepAlive = 60 * time.Second
	config.Producer.Return.Successes = false
	config.Producer.Flush.Frequency = time.Second
	config.Producer.Flush.MaxMessages = 10

	client, err := sarama.NewAsyncProducer(c.Brokers, config)

	if err != nil {
		return
	}

	producer = &Producer{
		c:      c,
		client: client,
		log:    logger,
		wg:     &sync.WaitGroup{},
	}

	producer.ctx, producer.cancel = context.WithCancel(context.Background())

	producer.wg.Add(1)
	go producer.logErr()

	return
}

func GetKafka() *ProducerConfig {
	return config.KaProducer
}
