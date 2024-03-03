package psql

import (
	"fmt"
	"hezzl/internal/model"
)

func (db *SqlPostgres) Select(reqId string, id int) (*model.Goods, error) {

	rows, err := db.dB.Query(
		`SELECT * 
	FROM goods
	WHERE id=$1`, id)

	if err != nil {
		db.logger.L.WithField("psql.Select", reqId).Error(err)
		return nil, err
	}
	defer rows.Close()

	var goods model.Goods

	for rows.Next() {

		err := rows.Scan(&goods.ID, &goods.ProjectID, &goods.Name, &goods.Description, &goods.Priority, &goods.Removed, &goods.Created_at)
		if err != nil {
			db.logger.L.WithField("psql.Select", reqId).Error(err)
			continue
		}
	}

	return &goods, nil

}

func (db *SqlPostgres) SelectDel(reqId string, id int) (*model.DeleteGoods, error) {

	rows, err := db.dB.Query(
		`SELECT id,project_id,removed 
		 FROM goods 
		 WHERE id=$1`, id)

	if err != nil {
		db.logger.L.WithField("psql.SelectDel", reqId).Error(err)
		return nil, err
	}
	defer rows.Close()

	var goods model.DeleteGoods

	for rows.Next() {

		err := rows.Scan(goods.ID, goods.ProjectID, goods.Removed)
		if err != nil {
			db.logger.L.WithField("psql.SelectDel", reqId).Error(err)
			continue
		}
	}

	return &goods, nil

}

func (db *SqlPostgres) Check(reqId string, id int) error {
	w, err := db.dB.Query("SELECT EXISTS(SELECT * FROM goods WHERE id=$1)", id)
	if err != nil {
		return err
	}
	defer w.Close()
	for w.Next() {
		var ok bool

		err := w.Scan(&ok)
		if err != nil {
			db.logger.L.WithField("psql.Check", reqId).Error(err)
			continue
		}
		db.logger.L.WithField("psql.Check", reqId).Debug("Значение OK--", ok)
		if !ok {
			return fmt.Errorf("пользователь с таким ID не существует")
		}
	}
	return nil
}

func (db *SqlPostgres) SelectPriority(reqId string, id int) (model.PriorityGoodsSL, error) {

	rows, err := db.dB.Query(
		`SELECT id,priority 
		 FROM goods 
		 WHERE id=$1`, id-1)

	if err != nil {
		db.logger.L.WithField("psql.SelectPriority", reqId).Error(err)
		return nil, err
	}
	defer rows.Close()

	var goodsSL model.PriorityGoodsSL

	for rows.Next() {
		var goods model.PriorityGoods
		err := rows.Scan(&goods.ID, &goods.Priority)
		if err != nil {
			db.logger.L.WithField("psql.SelectPriority", reqId).Error(err)
			continue
		}

		goodsSL = append(goodsSL, &goods)
	}

	return goodsSL, nil

}
