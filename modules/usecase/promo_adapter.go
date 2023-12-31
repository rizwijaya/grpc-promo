package usecase

import (
	promoProto "promo/modules/protobuf/pb"
	promoRepository "promo/modules/repositories"
)

type PromoAdapter interface {
	GetPromoByCode(code string) (*promoProto.Promo, error)
	GetAllOrder(offset int64, limit int64) ([]*promoProto.ListOrder, error)
	Pagination(table string, pages int64, limit int64) (int64, error)
	GetOrderById(id string) (*promoProto.Order, error)
	CreateOrder(orderAttr string, promoCode string) (*promoProto.ListOrder, error)
}

type PromoUseCase struct {
	repository *promoRepository.PromoRepository
}

func NewPromoUseCase(promoRepo *promoRepository.PromoRepository) *PromoUseCase {
	return &PromoUseCase{
		repository: promoRepo,
	}
}
