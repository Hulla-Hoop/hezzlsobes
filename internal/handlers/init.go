package handlers

import (
	"hezzl/internal/DB"
	"hezzl/internal/logger"
	"hezzl/internal/service"
)

type Handlers struct {
	service *service.Service
	logger  *logger.Logger
}

func New(db DB.DB, logger *logger.Logger, service *service.Service) *Handlers {
	return &Handlers{
		logger:  logger,
		service: service}
}
