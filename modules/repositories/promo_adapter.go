package repository

import (
	promoProto "promo/modules/protobuf/pb"

	"github.com/jmoiron/sqlx"
)

type PromoRepositoryPresenter interface {
	GetPromoByCode(code string) (*promoProto.Promo, error)
}

type PromoRepository struct {
	db *sqlx.DB
}

func NewPromoRepository(db *sqlx.DB) *PromoRepository {
	return &PromoRepository{db}
}
