package rediscash

import (
	"context"
	"encoding/json"
	"errors"
	"hezzl/internal/DB/psql"
	"hezzl/internal/model"
	"strconv"
	"time"
)

func (c *Cash) List(reqId string, page uint, limit int) (*model.List, error) {

	var List model.List

	meta, err := c.db.Meta(reqId, page, limit)
	if err != nil {
		return nil, err
	}
	List.Meta = *meta

	ctx := context.Background()
	for i := page; i < page+uint(limit); i++ {
		key := strconv.Itoa(int(i))
		v, err := c.r.Get(ctx, key).Bytes()
		if v == nil {
			c.log.L.WithField("List", reqId).Debug("Значения нет в Кэше---", key)
		} else if err != nil {
			c.log.L.WithField("List", reqId).Debug("Ошибка при получении значения из Кэша---", err)
		} else {
			var goods model.Goods
			err = json.Unmarshal(v, &goods)
			if err != nil {
				return nil, err
			}
			List.GoodsSL = append(List.GoodsSL, &goods)
			c.log.L.WithField("List", reqId).Debug("Значения взято из Кэша---", key, "----", goods)
			continue
		}

		goods, err := c.db.Select(reqId, int(i))
		var psqlErr *psql.ErrorNotFound
		if err != nil {
			if errors.As(err, &psqlErr) {
				c.log.L.WithField("Handlers.Update", reqId).Error(err)
				continue
			}
			return nil, err
		}
		c.log.L.WithField("List", reqId).Debug("Значения взято из db", key, "----", goods)
		List.GoodsSL = append(List.GoodsSL, goods)

		err = c.r.Set(ctx, key, goods, time.Minute).Err()
		if err != nil {
			c.log.L.WithField("List", reqId).Debug("Ошибка при записи значения в Кэш---", err)
		}
	}

	return &List, nil
}
