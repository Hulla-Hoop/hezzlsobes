package psql

import (
	"database/sql"
	"fmt"
	"hezzl/internal/model"
)

func (db *sqlPostgres) Create(goods *model.Goods) (*model.Goods, error) {

	var id int
	db.logger.Debug("db create полученные данные---", goods)
	err := db.dB.QueryRow(
		`INSERT INTO goods(name,project_id) 
		 VALUES ($1,$2) returning id`,
		goods.Name,
		goods.ProjectID,
	).Scan(&id)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return goods, fmt.Errorf("Пользователь добавлен но не удалось записать ID %s", err)
		default:
			return nil, fmt.Errorf("Ошибка при создании пользователя %s", err)
		}
	}
	goods, err = db.Select(id)
	if err != nil {
		return nil, err
	}

	db.logger.Debug("id созданого пользователя ----", &goods)
	return goods, nil

}
