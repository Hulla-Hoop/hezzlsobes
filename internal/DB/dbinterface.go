package DB

import (
	"hezzl/internal/model"
)

type DB interface {
	Create(user *model.Goods) (*int, error)
	Delete(id int) error
	Update(user *model.Goods, id int) error
	InsertPage(page uint, limit int) (model.GoodsSL, error)
}
