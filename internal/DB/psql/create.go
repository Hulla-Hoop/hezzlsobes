package psql

import (
	"database/sql"
	"fmt"
	"hezzl/internal/model"
)

func (db *SqlPostgres) Create(reqId string, name string, project_id int) (*model.Goods, error) {

	var id int
	db.logger.L.WithField("psql.Create", reqId).Debug("db create полученные данные---", name, project_id)
	err := db.dB.QueryRow(
		`INSERT INTO goods(name,project_id) 
		 VALUES ($1,$2) returning id`,
		name,
		project_id,
	).Scan(&id)
	db.logger.L.WithField("psql.Create", reqId).Debug("db create выходные данные ---", id)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return nil, fmt.Errorf("пользователь добавлен но не удалось записать ID %s", err)
		default:
			return nil, fmt.Errorf("ошибка при создании пользователя %s", err)
		}
	}
	goods, err := db.Select(reqId, id)
	if err != nil {
		return nil, err
	}

	db.logger.L.WithField("psql.Create", reqId).Debug("db create выходные данные ----", goods)
	return goods, nil

}
