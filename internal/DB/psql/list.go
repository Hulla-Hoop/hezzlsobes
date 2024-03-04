package psql

import "hezzl/internal/model"

//TODO Stop here

func (db *SqlPostgres) List(reqId string, page uint, limit int) (*model.List, error) {
	db.logger.L.WithField("psql.List", reqId).Debug("db insert page полученные данные---", page, limit)

	var List model.List
	List.Limit = limit
	List.Offset = int(page)

	cashPage := page*uint(limit) - 1

	rows, err := db.dB.Query(
		`	SELECT * 
			FROM goods 
			WHERE id > $1 
			ORDER BY id ASC 
			LIMIT $2 `, cashPage, limit)

	if err != nil {
		db.logger.L.WithField("psql.List", reqId).Error(err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var goods model.Goods

		err := rows.Scan(&goods.ID, &goods.ProjectID, &goods.Name, &goods.Description, &goods.Priority, &goods.Removed, &goods.Created_at)
		if err != nil {
			db.logger.L.WithField("psql.List", reqId).Error(err)
			continue
		}
		List.GoodsSL = append(List.GoodsSL, &goods)
	}

	meta, err := db.Meta(reqId, page, limit)
	if err != nil {
		return nil, err
	}
	List.Meta = *meta

	db.logger.L.WithField("psql.List", reqId).Debug("данные на выходе db list page ", &List)
	return &List, nil
}

func (db *SqlPostgres) Meta(reqID string, page uint, limit int) (*model.Meta, error) {
	var meta model.Meta
	meta.Limit = limit
	meta.Offset = int(page)
	row, err := db.dB.Query(
		`	SELECT 
			(SELECT COUNT(*) FROM goods) AS total,
			(SELECT COUNT(*) FROM goods WHERE removed=true) AS delete 
			FROM goods 
			GROUP BY total `)

	if err != nil {
		return nil, err
	}
	defer row.Close()

	for row.Next() {

		err := row.Scan(&meta.Total, &meta.Removed)
		if err != nil {
			db.logger.L.WithField("psql.Meta", reqID).Error(err)
			continue
		}
	}
	return &meta, nil
}
