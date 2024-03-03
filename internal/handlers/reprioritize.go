package handlers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (e *Handlers) Reprioritize(c echo.Context) error {
	reqId := c.Get("reqId").(string)

	var priority int

	str := c.Param("id")

	Id, err := strconv.Atoi(str)
	if err != nil {
		e.logger.L.WithField("Handlers.Reprioritize", reqId).Error(err)
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	e.logger.L.WithField("Handlers.Reprioritize", reqId).Debug(Id)

	err = c.Bind(priority)
	if err != nil {
		e.logger.L.WithField("Handlers.Reprioritize", reqId).Error(err)
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	list, err := e.service.Reprioritize(reqId, Id, priority)
	if err != nil {
		e.logger.L.WithField("Handlers.Reprioritize", reqId).Error(err)
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	e.logger.L.WithField("Handlers.Reprioritize", reqId).Debug("товар добавлен")
	return c.JSON(http.StatusOK, list)
}
