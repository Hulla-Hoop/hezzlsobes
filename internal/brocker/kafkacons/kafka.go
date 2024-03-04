package kafkacons

import (
	"encoding/json"
	"hezzl/internal/clickhouse"
	"hezzl/internal/config"
	"hezzl/internal/logger"
	"hezzl/internal/model"
	"sync"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type Consumer struct {
	c     *kafka.Consumer
	log   *logger.Logger
	topic string
	click *clickhouse.ClickHouse
	mu    sync.Mutex
	Rows  map[time.Time]model.LogGoods
}

func NewConsumer(log *logger.Logger) *Consumer {
	cfg := config.NewKafka()
	click := clickhouse.Init(log)
	conf := make(kafka.ConfigMap)
	conf["bootstrap.servers"] = cfg.BootstrapService
	conf["group.id"] = cfg.GroupID
	conf["auto.offset.reset"] = cfg.AutoOffsetReset
	c, err := kafka.NewConsumer(&conf)
	if err != nil {
		log.L.Info("не удалось подключиться к кафка")
		return nil
	}
	return &Consumer{
		c:     c,
		log:   log,
		topic: cfg.Topic,
		click: click,
		Rows:  make(map[time.Time]model.LogGoods),
	}
}

func (c *Consumer) Close() {
	defer c.log.L.WithField("Consumer.Close", "").Info("Close")
	c.click.Close()
	defer c.log.L.Info("Close KafkaConsumer")
	c.c.Close()
}

func (c *Consumer) Read() {

	c.c.SubscribeTopics([]string{c.topic}, nil)
	for {
		msg, err := c.c.ReadMessage(-1)
		var Log model.LogGoods
		if err == nil {
			err = json.Unmarshal(msg.Value, &Log)
			if err != nil {
				c.log.L.WithField("Consumer.Read", "").Error(err)
			}
			c.log.L.Debug("Received message: ", Log)
			c.Append(Log)
			go func() {
				for {
					if len(c.Rows) == 5 {
						wer := c.Clear()
						c.log.L.WithField("Consumer.Read", "    ").Debug("Rows: ", wer, "LEN", len(wer))
						c.click.Create(wer)
					}
				}
			}()

		} else {
			c.log.L.Debug("Consumer error: ", err)
			break
		}
	}
}

func (c *Consumer) Append(log model.LogGoods) {
	c.mu.Lock()
	defer c.mu.Unlock()
	key := time.Now()
	c.Rows[key] = log
	c.log.L.Debug("Rows: ", c.Rows, len(c.Rows))
}

func (c *Consumer) Clear() map[time.Time]model.LogGoods {
	c.mu.Lock()
	defer c.mu.Unlock()
	rows := c.Rows
	c.Rows = make(map[time.Time]model.LogGoods)

	return rows
}
