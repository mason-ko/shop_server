package service

import "atos.com/domain"

type ShopService struct {
	shopRepository domain.ShopRepository
}

func NewShopService(shopRepository domain.ShopRepository) domain.ShopService {
	svc := &ShopService{
		shopRepository: shopRepository,
	}

	return svc
}

func (s *ShopService) Get(id string) (domain.Shop, error) {
	return s.shopRepository.Get(id)
}

func (s *ShopService) Search(x, y string, limit int) ([]domain.Shop, error) {
	return s.shopRepository.Search(x, y, limit)
}
