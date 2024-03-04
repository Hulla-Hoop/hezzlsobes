package handlers

import (
	"errors"
	"hezzl/internal/DB/psql"
	"hezzl/internal/model"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (e *Handlers) Update(c echo.Context) error {
	reqId := c.Get("reqId").(string)
	var upd model.UpdateGoods
	err := c.Bind(&upd)
	if err != nil {
		e.logger.L.WithField("Handlers.Update", reqId).Error(err)
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	valueStr, err := c.FormParams()
	if err != nil {
		e.logger.L.WithField("Handlers.Update", reqId).Error(err)
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	idStr := valueStr["id"]

	id, err := strconv.Atoi(idStr[0])
	if err != nil {
		e.logger.L.WithField("Handlers.Update", reqId).Error(err)
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	prIdStr := valueStr["projectId"]

	pid, err := strconv.Atoi(prIdStr[0])
	if err != nil {
		e.logger.L.WithField("Handlers.Update", reqId).Error(err)
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	goods, err := e.service.Update(reqId, pid, id, upd.Name, upd.Description)

	var psqlErr *psql.ErrorNotFound

	if err != nil {
		if errors.As(err, &psqlErr) {
			e.logger.L.WithField("Handlers.Update", reqId).Error(err)
			psqlErr.Code = 3
			psqlErr.Msg = "errors.good.notFound"
			psqlErr.Details = struct{}{}
			return c.JSON(http.StatusNotFound, psqlErr)
		}
		e.logger.L.WithField("Handler s.Update", reqId).Error(err)
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, goods)
}
