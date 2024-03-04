package kafkaprod

import (
	"encoding/json"
	"hezzl/internal/config"
	"hezzl/internal/logger"
	"hezzl/internal/model"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type Producer struct {
	p     *kafka.Producer
	log   *logger.Logger
	topic string
}

func NewProducer(log *logger.Logger) *Producer {
	cfg := config.NewKafka()
	conf := make(kafka.ConfigMap)
	conf["bootstrap.servers"] = cfg.BootstrapService
	p, err := kafka.NewProducer(&conf)

	if err != nil {
		log.L.Info("не удалось подключиться к кафка")
		return nil
	}
	return &Producer{
		p:     p,
		log:   log,
		topic: cfg.Topic,
	}
}

func (p *Producer) Close() {
	p.p.Close()
}

func (p *Producer) Send(reqId string, value model.LogGoods) error {
	p.log.L.WithField("Producer.Send", reqId).Debug(value)
	a, err := json.Marshal(value)
	if err != nil {
		p.log.L.WithField("Producer.Send", reqId).Error(err)
		return err
	}
	p.p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &p.topic, Partition: kafka.PartitionAny},
		Key:            []byte(reqId),
		Value:          a,
	}, nil)

	return nil
}
