package usecase

import (
	promoRepository "promo/modules/repositories"
)

type PromoAdapter interface {
}

type PromoUseCase struct {
	repository *promoRepository.PromoRepository
}

func NewPromoUseCase(promoRepo *promoRepository.PromoRepository) *PromoUseCase {
	return &PromoUseCase{
		repository: promoRepo,
	}
}
