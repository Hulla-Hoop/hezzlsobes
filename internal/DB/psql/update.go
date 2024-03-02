package psql

import (
	"fmt"
	"hezzl/internal/model"
)

func (db *sqlPostgres) Update(project_id int, id int, name string, description string) (*model.Goods, error) {

	err := db.Check(id)
	if err != nil {
		return nil, err
	}
	db.logger.Debug("db update полученные данные---", project_id, name, description, "--id----", id)
	var desc string
	if description == "" {
		desc = " "
	} else {
		desc = fmt.Sprintf("patronymic = '%s',", description)
	}

	update := fmt.Sprintf("UPDATE goods SET %s  name=$1  WHERE id=$2 ", desc)
	_, err = db.dB.Exec(update,
		name,
		id)
	if err != nil {
		return nil, err
	}
	goods, err := db.Select(id)
	if err != nil {
		return nil, err
	}

	return goods, nil
}
