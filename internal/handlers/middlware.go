package handlers

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func (e *Handlers) ReqID(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		reqId := c.Get("reqID")
		if reqId == nil {
			reqId = uuid.New().String()
		}
		c.Set("reqId", reqId)

		return next(c)
	}
}
