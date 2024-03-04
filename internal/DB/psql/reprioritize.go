package psql

import (
	"fmt"
	"hezzl/internal/model"
)

func (db *SqlPostgres) Reprioritize(reqId string, id int, priority int) (model.PriorityGoodsSL, error) {

	err := db.Check(reqId, id)
	if err != nil {
		return nil, err
	}

	db.logger.L.WithField("psql.Reprioritize", reqId).Debug("полученные данные---", id, "--", priority)

	zpr := fmt.Sprintf(`
	BEGIN;

	SELECT * FROM goods WHERE id > %d  FOR UPDATE;

	SELECT setval('priority_seq',%d,false);

	UPDATE goods 
	SET priority = nextval('priority_seq'::regclass) 
	WHERE id>%d ;

	COMMIT;
	
	`, id-1, priority, id-1)

	_, err = db.dB.Exec(zpr)

	if err != nil {
		db.logger.L.WithField("psql.Reprioritize", reqId).Error(err)
		return nil, err
	}

	dG, err := db.SelectPriority(reqId, id)
	if err != nil {
		db.logger.L.WithField("psql.Reprioritize", reqId).Error(err)
		return nil, err
	}

	db.logger.L.WithField("psql.Reprioritize", reqId).Debug("Выходные данные", dG)
	return dG, nil
}
