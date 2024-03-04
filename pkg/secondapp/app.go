package secondapp

import (
	"hezzl/internal/brocker/kafkacons"
	"hezzl/internal/logger"
)

type App struct {
	C *kafkacons.Consumer
}

func New(log *logger.Logger) *App {
	a := App{}
	a.C = kafkacons.NewConsumer(log)
	return &a
}

func (a *App) Close() {
	a.C.Close()
}
