package firstapp

import (
	"context"
	"hezzl/internal/DB/psql"
	"hezzl/internal/DB/rediscash"
	"hezzl/internal/brocker/kafkaprod"
	"hezzl/internal/handlers"
	"hezzl/internal/logger"
	"hezzl/internal/service"

	"github.com/labstack/echo/v4"
)

type App struct {
	h   *handlers.Handlers
	e   echo.Echo
	log *logger.Logger
}

func New(log *logger.Logger) *App {
	a := App{}
	a.log = log
	db, err := psql.InitDb(a.log)
	if err != nil {
		a.log.L.Info("База данных не поднялась", err)
		return nil
	}

	prod := kafkaprod.NewProducer(a.log)
	if prod == nil {
		a.log.L.Info("Не удалось подключиться к кафка")
		return nil
	}

	r := rediscash.Init(db, a.log)

	s := service.New(a.log, r, prod)

	a.h = handlers.New(a.log, s)

	a.e = *echo.New()

	a.e.Use(a.h.ReqID)

	good := a.e.Group("/good")

	good.POST("/create", a.h.Create)

	good.DELETE("/remove", a.h.Remove)

	good.PATCH("/update", a.h.Update)

	good.GET("/list", a.h.List)

	good.PATCH("/reprioritize", a.h.Reprioritize)

	return &a

}

func (a *App) Start() error {
	a.log.L.Info("Запуск сервера на localhost:1234")
	err := a.e.Start(":1234")
	if err != nil {
		return err
	}
	return nil
}

func (a *App) Stop() {
	a.log.L.Info("Остановка сервера")
	ctx := context.Background()
	a.h.Close()
	a.e.Shutdown(ctx)

}
