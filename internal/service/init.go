package service

import (
	"errors"
	"hezzl/internal/DB"
	"hezzl/internal/brocker/kafkaprod"
	"hezzl/internal/logger"
	"regexp"
)

type Service struct {
	Log   *logger.Logger
	DB    DB.DB
	kafka *kafkaprod.Producer
}

func New(log *logger.Logger, db DB.DB, kafka *kafkaprod.Producer) *Service {
	return &Service{
		Log:   log,
		DB:    db,
		kafka: kafka,
	}
}

func (s *Service) Chek(name string) error {
	r, err := regexp.MatchString("^[a-zA-Z0-9]+$", name)

	if err != nil {
		return err
	}
	if !r {
		err = errors.New("неверный формат поля имя")
		return err
	}
	return nil
}

func (s *Service) Close() {
	s.Log.L.Info("Close KafkaProducer")
	s.kafka.Close()
	s.Log.L.Info("Close DB")
	s.DB.Close()
}
