package handlers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (e *Handlers) Create(c echo.Context) error {
	reqId := c.Get("reqId").(string)

	var name string

	str := c.Param("projectId")

	projectId, err := strconv.Atoi(str)
	if err != nil {
		e.logger.L.WithField("Handlers.Create", reqId).Error(err)
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	e.logger.L.WithField("Handlers.Create", reqId).Debug(projectId)

	err = c.Bind(name)

	if err != nil {
		e.logger.L.WithField("Handlers.Create", reqId).Error(err)
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	goods, err := e.service.Create(reqId, name, projectId)
	if err != nil {
		e.logger.L.WithField("Handlers.Create", reqId).Error(err)
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	e.logger.L.WithField("Handlers.Create", reqId).Debug("товар добавлен")
	return c.JSON(http.StatusCreated, goods)
}
