package service

import (
	"hezzl/internal/DB"
	"hezzl/internal/logger"
)

type Service struct {
	Log *logger.Logger
	DB  DB.DB
}

func New(log *logger.Logger, db DB.DB) *Service {
	return &Service{
		Log: log,
		DB:  db,
	}
}
