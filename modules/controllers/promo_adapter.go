package controllers

import (
	promoProto "promo/modules/protobuf/pb"
	promoRepository "promo/modules/repositories"
	promoUseCase "promo/modules/usecase"

	"github.com/jmoiron/sqlx"
)

type PromoController struct {
	usecase *promoUseCase.PromoUseCase
	promoProto.UnimplementedPromoServiceServer
}

func NewPromoController(db *sqlx.DB) *PromoController {
	promoRepo := promoRepository.NewPromoRepository(db)
	promoUseCase := promoUseCase.NewPromoUseCase(promoRepo)
	return &PromoController{
		usecase: promoUseCase,
	}
}
