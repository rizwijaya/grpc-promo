package repository

import (
	promoProto "promo/modules/protobuf/pb"

	"github.com/jmoiron/sqlx"
)

type PromoRepositoryPresenter interface {
	GetPromoByCode(code string) (*promoProto.Promo, error)
	TotalData(table string) (int64, error)
	GetAllOrder(offset int64, limit int64) ([]*promoProto.ListOrder, error)
}

type PromoRepository struct {
	db *sqlx.DB
}

func NewPromoRepository(db *sqlx.DB) *PromoRepository {
	return &PromoRepository{db}
}
