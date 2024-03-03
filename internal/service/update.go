package service

import "hezzl/internal/model"

func (s *Service) Update(reqId string, id int, name string, description string) (*model.Goods, error) {
	goods, err := s.DB.Update(reqId, id, name, description)
	if err != nil {
		return nil, err
	}

	return goods, nil
}
