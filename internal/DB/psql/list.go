package psql

import "hezzl/internal/model"

//TODO Stop here

func (db *sqlPostgres) List(page uint, limit int) (*model.List, error) {
	db.logger.Debug("db insert page полученные данные---", page, limit)

	List := model.NewList()

	cashPage := page*uint(limit) - 1

	rows, err := db.dB.Query(
		`	SELECT * 
			FROM goods 
			WHERE id > $1 
			ORDER BY id ASC 
			LIMIT $2 `, cashPage, limit)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		goods := model.NewGoods()

		err := rows.Scan(goods.ID, goods.ProjectID, goods.Name, goods.Description, goods.Priority, goods.Removed, goods.Created_at)
		if err != nil {
			db.logger.Error(err)
			continue
		}
		cashPage = uint(goods.ID)
		List.GoodsSL = append(List.GoodsSL, goods)
	}

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

		err := row.Scan(List.Meta.Total, List.Meta.Removed)
		if err != nil {
			db.logger.Error(err)
			continue
		}
	}

	db.logger.Debug("данные на выходе db insert page ", &List)
	return List, nil
}
