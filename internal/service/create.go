package service

import "hezzl/internal/model"

func (s *Service) Create(reqId string, name string, projectId int) (*model.Goods, error) {
	goods, err := s.DB.Create(reqId, name, projectId)
	if err != nil {
		return nil, err
	}
	return goods, nil
}
