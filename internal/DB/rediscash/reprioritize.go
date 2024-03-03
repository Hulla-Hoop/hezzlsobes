package rediscash

import "hezzl/internal/model"

func (c *Cash) Reprioritize(reqId string, id int, priority int) (model.PriorityGoodsSL, error) {

	goods, err := c.db.Reprioritize(reqId, id, priority)
	if err != nil {
		return nil, err
	}
	for _, i := range goods {
		err := c.DelVal(i.ID)
		if err != nil {
			return nil, err
		}
	}
	return goods, nil
}
