package service

import "hezzl/internal/model"

func (s *Service) List(reqId string, offset uint, limit int) (*model.List, error) {
	list, err := s.DB.List(reqId, offset, limit)
	if err != nil {
		return nil, err
	}
	return list, nil
}
