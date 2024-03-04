package service

import (
	"hezzl/internal/model"
	"time"
)

func (s *Service) Delete(reqId string, pid int, id int) (*model.DeleteGoods, error) {
	var log model.LogGoods
	del, err := s.DB.Delete(reqId, id)
	if err != nil {
		return nil, err
	}
	log.ID = del.ID
	log.ProjectID = pid
	log.Removed = del.Removed
	log.EventTime = time.Now()

	err = s.kafka.Send(reqId, log)
	if err != nil {
		s.Log.L.WithField("Service.Update", reqId).Error(err)
		return nil, err
	}
	return del, nil

}
