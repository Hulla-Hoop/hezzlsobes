package service

import "hezzl/internal/model"

func (s *Service) Reprioritize(reqId string, id int, priority int) (model.PriorityGoodsSL, error) {

	pr, err := s.DB.Reprioritize(reqId, id)
	if err != nil {
		return nil, err
	}

	return pr, nil

}
