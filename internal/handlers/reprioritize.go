package handlers

import (
	"errors"
	"hezzl/internal/DB/psql"
	"hezzl/internal/model"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (e *Handlers) Reprioritize(c echo.Context) error {
	reqId := c.Get("reqId").(string)

	var priority model.PriorityGoods

	valueStr, err := c.FormParams()
	if err != nil {
		e.logger.L.WithField("Handlers.Update", reqId).Error(err)
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	idStr := valueStr["id"]

	Id, err := strconv.Atoi(idStr[0])
	if err != nil {
		e.logger.L.WithField("Handlers.Reprioritize", reqId).Error(err)
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	e.logger.L.WithField("Handlers.Reprioritize", reqId).Debug(Id)

	prIdStr := valueStr["projectId"]

	pid, err := strconv.Atoi(prIdStr[0])
	if err != nil {
		e.logger.L.WithField("Handlers.Update", reqId).Error(err)
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	err = c.Bind(&priority)
	if err != nil {
		e.logger.L.WithField("Handlers.Reprioritize", reqId).Error(err)
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	var psqlErr *psql.ErrorNotFound
	list, err := e.service.Reprioritize(reqId, pid, Id, priority.Priority)
	if err != nil {
		if errors.As(err, &psqlErr) {
			e.logger.L.WithField("Handlers.Update", reqId).Error(err)
			psqlErr.Code = 3
			psqlErr.Msg = "errors.good.notFound"
			psqlErr.Details = struct{}{}
			return c.JSON(http.StatusNotFound, psqlErr)
		}
		e.logger.L.WithField("Handlers.Reprioritize", reqId).Error(err)
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	e.logger.L.WithField("Handlers.Reprioritize", reqId).Debug("товар добавлен")
	return c.JSON(http.StatusOK, list)
}
