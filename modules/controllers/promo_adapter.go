package controllers

import (
	promoProto "promo/modules/protobuf/pb"
	promoRepository "promo/modules/repositories"
	promoUseCase "promo/modules/usecase"

	"github.com/jmoiron/sqlx"
)

type PromoUseCase struct {
	usecase *promoUseCase.PromoUseCase
	promoProto.UnimplementedPromoServiceServer
}

func NewPromoUseCase(db *sqlx.DB) *PromoUseCase {
	promoRepo := promoRepository.NewPromoRepository(db)
	promoUseCase := promoUseCase.NewPromoUseCase(promoRepo)
	return &PromoUseCase{
		usecase: promoUseCase,
	}
}
