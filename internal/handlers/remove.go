package handlers

import (
	"errors"
	"hezzl/internal/DB/psql"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (e *Handlers) Remove(c echo.Context) error {
	reqId := c.Get("reqId").(string)
	valueStr, err := c.FormParams()
	if err != nil {
		e.logger.L.WithField("Handlers.Update", reqId).Error(err)
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	idStr := valueStr["id"]
	idi, err := strconv.Atoi(idStr[0])
	if err != nil {
		e.logger.L.WithField("Handlers.Remove", reqId).Error(err)
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	prIdStr := valueStr["projectId"]

	pid, err := strconv.Atoi(prIdStr[0])
	if err != nil {
		e.logger.L.WithField("Handlers.Update", reqId).Error(err)
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	var psqlErr *psql.ErrorNotFound

	del, err := e.service.Delete(reqId, pid, idi)

	if err != nil {
		if errors.As(err, &psqlErr) {
			e.logger.L.WithField("Handlers.Update", reqId).Error(err)
			psqlErr.Code = 3
			psqlErr.Msg = "errors.good.notFound"
			psqlErr.Details = struct{}{}
			return c.JSON(http.StatusNotFound, psqlErr)
		}
		e.logger.L.WithField("Handlers.Remove", reqId).Error(err)
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	e.logger.L.WithField("Handlers.Remove", reqId).Debug("пользователь удален")
	return c.JSON(http.StatusAccepted, del)
}
