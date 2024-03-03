package rediscash

import (
	"hezzl/internal/model"
)

func (c *Cash) Create(reqId string, name string, project_id int) (*model.Goods, error) {
	good, err := c.db.Create(reqId, name, project_id)
	if err != nil {
		return nil, err
	}
	return good, nil
}
