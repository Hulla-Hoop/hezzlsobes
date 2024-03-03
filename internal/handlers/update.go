package handlers

import (
	"hezzl/internal/model"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (e *Handlers) Update(c echo.Context) error {
	reqId := c.Get("reqId").(string)
	var upd model.UpdateGoods
	err := c.Bind(upd)
	if err != nil {
		e.logger.L.WithField("Handlers.Update", reqId).Error(err)
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		e.logger.L.WithField("Handlers.Update", reqId).Error(err)
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	goods, err := e.service.Update(reqId, id, upd.Name, upd.Description)
	if err != nil {
		e.logger.L.WithField("Handler s.Update", reqId).Error(err)
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	e.logger.L.WithField("Handlers.Update", reqId).Error(err)
	return c.JSON(http.StatusOK, goods)
}
