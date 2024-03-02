package psql

import "hezzl/internal/model"

func (db *sqlPostgres) Delete(id int) (*model.DeleteGoods, error) {

	err := db.Check(id)
	if err != nil {
		return nil, err
	}

	dG := model.NewDeleteGoods()

	db.logger.Debug("db delete полученные данные---", id)

	result, err := db.dB.Exec(
		`UPDATE goods SET remove=true WHERE id=$1`,
		id)
	if err != nil {
		return nil, err
	}

	dG, err = db.SelectDel(id)
	if err != nil {
		return nil, err
	}

	db.logger.Info("Пользователь успешно удален ", result)
	return dG, nil
}
