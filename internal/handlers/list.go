package handlers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (e *Handlers) UserPagination(c echo.Context) error {
	reqId := c.Get("reqId").(string)

	valueStr, err := c.FormParams()
	if err != nil {
		e.logger.L.WithField("Handlers.Update", reqId).Error(err)
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	offsetStr := valueStr["offset"]
	limitStr := valueStr["limit"]

	e.logger.L.WithField("Handlers.Create", reqId).Debug(offsetStr, limitStr)

	offset, err := strconv.Atoi(offsetStr[0])
	if err != nil {
		e.logger.L.WithField("Handlers.Update", reqId).Error(err)
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	limit, err := strconv.Atoi(limitStr[0])
	if err != nil {
		e.logger.L.WithField("Handlers.Update", reqId).Error(err)
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	u, err := e.service.List(reqId, uint(offset), limit)
	if err != nil {
		e.logger.L.WithField("Handlers.Update", reqId).Error(err)
		return c.JSON(http.StatusInternalServerError, err.Error())

	}

	return c.JSON(http.StatusOK, u)

}
