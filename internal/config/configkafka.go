package config

import (
	"os"
)

type Configkafka struct {
	BootstrapService string
	GroupID          string
	AutoOffsetReset  string
	Topic            string
	TopicErr         string
}

func NewKafka() *Configkafka {

	return &Configkafka{
		BootstrapService: os.Getenv("BOOTSTRAPSERVER"),
		AutoOffsetReset:  os.Getenv("AUTOOFFSETRESET"),
		GroupID:          os.Getenv("GROUPID"),
		Topic:            os.Getenv("TOPIC"),
		TopicErr:         os.Getenv("TOPICERR"),
	}
}
