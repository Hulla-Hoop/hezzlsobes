package service

import (
	"hezzl/internal/model"
	"time"
)

func (s *Service) Update(reqId string, project_id int, id int, name string, description string) (*model.Goods, error) {

	var log model.LogGoods

	err := s.Chek(name)

	if err != nil {
		s.Log.L.WithField("Service.Update", reqId).Error(err)
		return nil, err
	}

	goods, err := s.DB.Update(reqId, id, name, description)
	if err != nil {
		s.Log.L.WithField("Service.Update", reqId).Error(err)
		return nil, err
	}
	log.ID = goods.ID
	log.Name = goods.Name
	log.Description = goods.Description
	log.ProjectID = goods.ProjectID
	log.Removed = goods.Removed
	log.EventTime = time.Now()

	err = s.kafka.Send(reqId, log)
	if err != nil {
		s.Log.L.WithField("Service.Update", reqId).Error(err)
		return nil, err
	}

	return goods, nil
}
