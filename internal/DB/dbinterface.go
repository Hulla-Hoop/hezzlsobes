package DB

import (
	"hezzl/internal/model"
)

type DB interface {
	Create(reqId string, name string, project_id int) (*model.Goods, error)
	Delete(reqId string, id int) (*model.DeleteGoods, error)
	Update(reqId string, id int, name string, description string) (*model.Goods, error)
	List(reqId string, page uint, limit int) (*model.List, error)
	Reprioritize(reqId string, id int, priority int) (model.PriorityGoodsSL, error)
	Close()
}
