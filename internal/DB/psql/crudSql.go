package psql

import (
	"fmt"
	"hezzl/internal/model"
)

func (db *sqlPostgres) Select(id int) (*model.Goods, error) {

	rows, err := db.dB.Query(
		`SELECT * 
	FROM goods
	WHERE id=$1`, id)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	goods := model.NewGoods()

	for rows.Next() {

		err := rows.Scan(goods.ID, goods.ProjectID, goods.Name, goods.Description, goods.Priority, goods.Removed, goods.Created_at)
		if err != nil {
			db.logger.Error(err)
			continue
		}
	}

	return goods, nil

}

func (db *sqlPostgres) SelectDel(id int) (*model.DeleteGoods, error) {

	rows, err := db.dB.Query(
		`SELECT id,project_id,remove 
		 FROM goods 
		 WHERE id=$1`)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	goods := model.NewDeleteGoods()

	for rows.Next() {

		err := rows.Scan(goods.ID, goods.ProjectID, goods.Removed)
		if err != nil {
			db.logger.Error(err)
			continue
		}
	}

	return goods, nil

}

func (db *sqlPostgres) Check(id int) error {
	w, err := db.dB.Query("SELECT EXISTS(SELECT * FROM goods WHERE id=$1)", id)
	if err != nil {
		return err
	}
	defer w.Close()
	for w.Next() {
		var ok bool

		err := w.Scan(&ok)
		if err != nil {
			db.logger.Error(err)
			continue
		}
		db.logger.Debug("Значение OK--", ok)
		if !ok {
			return fmt.Errorf("Пользователь с таким ID не существует")
		}
	}
}
