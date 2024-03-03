package rediscash

import "hezzl/internal/model"

func (c *Cash) Delete(reqId string, id int) (*model.DeleteGoods, error) {
	err := c.DelVal(id)
	if err != nil {
		return nil, err
	}
	del, err := c.db.Delete(reqId, id)
	if err != nil {
		return nil, err
	}
	return del, nil
}
