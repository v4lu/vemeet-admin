package services

import "github.com/valu/vemeet-admin-api/internal/data"

type BlockedService struct {
	blockedRepo data.BlockedRepositoryInterface
}

type BlockedServiceInterface interface {
	GetBlockedById(id int64) (*data.Blocked, error)
	GetBlockeds(page int64, limit int64, sort, order, search string) (*data.BlockedPagination, error)
	CreateBlocked(blocked *data.Blocked) (*data.Blocked, error)
	UpdateBlocked(blocked *data.Blocked) (*data.Blocked, error)
	DeleteBlocked(id int64) (bool, error)
}

func NewBlockedService(blockedRepo data.BlockedRepositoryInterface) BlockedServiceInterface {
	return &BlockedService{blockedRepo}
}

func (s *BlockedService) GetBlockedById(id int64) (*data.Blocked, error) {
	blocked, err := s.blockedRepo.FindById(id)
	if err != nil {
		return nil, err
	}

	return blocked, nil
}

func (s *BlockedService) GetBlockeds(page int64, limit int64, sort, order, search string) (*data.BlockedPagination, error) {
	blockeds, err := s.blockedRepo.FindAll(page, limit, sort, order, search)
	if err != nil {
		return nil, err
	}

	return blockeds, nil
}

func (s *BlockedService) CreateBlocked(blocked *data.Blocked) (*data.Blocked, error) {
	blocked, err := s.blockedRepo.Create(blocked)
	if err != nil {
		return nil, err
	}

	return blocked, nil
}

func (s *BlockedService) UpdateBlocked(blocked *data.Blocked) (*data.Blocked, error) {
	blocked, err := s.blockedRepo.Update(blocked)
	if err != nil {
		return nil, err
	}

	return blocked, nil
}

func (s *BlockedService) DeleteBlocked(id int64) (bool, error) {
	deleted, err := s.blockedRepo.Delete(id)
	if err != nil {
		return false, err
	}

	return deleted, nil
}
