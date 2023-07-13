package usecase

import (
	promoProto "promo/modules/handler/pb"
	promoRepository "promo/modules/repositories"

	"github.com/jmoiron/sqlx"
)

type PromoAdapter interface {
}

type PromoUseCase struct {
	repository *promoRepository.PromoRepository
	promoProto.UnimplementedPromoServiceServer
}

func NewPromoUseCase(db *sqlx.DB) *PromoUseCase {
	promoRepo := promoRepository.NewPromoRepository(db)
	return &PromoUseCase{
		repository: promoRepo,
	}
}
