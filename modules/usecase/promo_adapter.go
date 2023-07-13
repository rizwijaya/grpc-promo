package usecase

import (
	promoProto "promo/modules/protobuf/pb"
	promoRepository "promo/modules/repositories"
)

type PromoAdapter interface {
	GetPromoByCode(code string) (*promoProto.Promo, error)
}

type PromoUseCase struct {
	repository *promoRepository.PromoRepository
}

func NewPromoUseCase(promoRepo *promoRepository.PromoRepository) *PromoUseCase {
	return &PromoUseCase{
		repository: promoRepo,
	}
}
