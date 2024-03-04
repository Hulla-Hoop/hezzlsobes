package service

import (
	"hezzl/internal/model"
	"time"
)

func (s *Service) Reprioritize(reqId string, pid int, id int, priority int) (model.PriorityGoodsSL, error) {
	var log model.LogGoods
	pr, err := s.DB.Reprioritize(reqId, id, priority)
	if err != nil {
		return nil, err
	}
	for _, i := range pr {
		log.ID = i.ID
		log.EventTime = time.Now()
		log.Priority = i.Priority + 1
		err = s.kafka.Send(reqId, log)
		if err != nil {
			s.Log.L.WithField("Service.Reprioritize", reqId).Error(err)
			return nil, err
		}
	}
	return pr, nil

}
