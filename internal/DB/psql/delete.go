package psql

import (
	"fmt"
	"hezzl/internal/model"
)

func (db *SqlPostgres) Delete(reqId string, id int) (*model.DeleteGoods, error) {

	err := db.Check(reqId, id)
	if err != nil {
		return nil, err
	}

	db.logger.L.WithField("psql.Delete", id).Debug("db delete полученные данные---", id)

	str := fmt.Sprintf(`
	BEGIN;

	SELECT * FROM goods WHERE id = %d FOR UPDATE;

	UPDATE goods 
	SET removed = true 
	WHERE id=%d;

	COMMIT;
	
	`, id, id)

	_, err = db.dB.Exec(str)

	if err != nil {
		return nil, err
	}

	dG, err := db.SelectDel(reqId, id)
	if err != nil {
		return nil, err
	}

	db.logger.L.WithField("psql.Delete", id).Info("Пользователь успешно удален ", dG)
	return dG, nil
}
