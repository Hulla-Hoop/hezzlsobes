package service

import (
	"hezzl/internal/model"
)

func (s *Service) Create(reqId string, name string, projectId int) (*model.Goods, error) {
	err := s.Chek(name)

	if err != nil {
		s.Log.L.WithField("Service.Create", reqId).Error(err)
		return nil, err
	}
	goods, err := s.DB.Create(reqId, name, projectId)
	if err != nil {
		return nil, err
	}
	return goods, nil
}
