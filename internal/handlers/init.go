package handlers

import (
	"hezzl/internal/logger"
	"hezzl/internal/service"
)

type Handlers struct {
	service *service.Service
	logger  *logger.Logger
}

func New(logger *logger.Logger, service *service.Service) *Handlers {
	return &Handlers{
		logger:  logger,
		service: service}
}

func (e *Handlers) Close() {
	e.logger.L.Info("Close handlers")
	e.service.Close()
}
