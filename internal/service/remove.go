package service

import "hezzl/internal/model"

func (s *Service) Delete(reqId string, id int) (*model.DeleteGoods, error) {
	del, err := s.DB.Delete(reqId, id)
	if err != nil {
		return nil, err
	}
	return del, nil

}
