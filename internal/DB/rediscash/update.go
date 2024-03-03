package rediscash

import (
	"hezzl/internal/model"
)

func (c *Cash) Update(reqId string, id int, name string, description string) (*model.Goods, error) {
	err := c.DelVal(id)
	if err != nil {
		return nil, err
	}
	goods, err := c.db.Update(reqId, id, name, description)
	if err != nil {
		return nil, err
	}
	return goods, nil
}
