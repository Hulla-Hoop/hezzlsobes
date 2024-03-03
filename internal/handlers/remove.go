package handlers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (e *Handlers) Remove(c echo.Context) error {
	reqId := c.Get("reqId").(string)
	id := c.Param("id")
	idi, err := strconv.Atoi(id)
	if err != nil {
		e.logger.L.WithField("Handlers.Remove", reqId).Error(err)
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	del, err := e.service.Delete(reqId, idi)
	if err != nil {
		e.logger.L.WithField("Handlers.Remove", reqId).Error(err)
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	e.logger.L.WithField("Handlers.Remove", reqId).Debug("пользователь удален")
	return c.JSON(http.StatusAccepted, del)
}
