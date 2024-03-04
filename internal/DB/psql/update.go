package psql

import (
	"fmt"
	"hezzl/internal/model"
)

func (db *SqlPostgres) Update(reqId string, id int, name string, description string) (*model.Goods, error) {

	err := db.Check(reqId, id)
	if err != nil {
		return nil, err
	}
	db.logger.L.WithField("psql.Update", reqId).Debug("db update полученные данные---", name, description, "--id----", id)

	update := fmt.Sprintf(`
	BEGIN;

	SELECT * FROM goods WHERE id = %d FOR UPDATE;

	UPDATE goods 
	SET description='%s',name='%s'  
	WHERE id=%d;
	
	COMMIT;
	`, id, description, name, id)
	db.logger.L.WithField("psql.Update", reqId).Debug("db update запрос---", update)
	_, err = db.dB.Exec(update)
	if err != nil {
		db.logger.L.WithField("psql.Update", reqId).Error(err)
		return nil, err
	}
	goods, err := db.Select(reqId, id)
	if err != nil {
		db.logger.L.WithField("psql.Update", reqId).Error(err)
		return nil, err
	}

	return goods, nil
}
