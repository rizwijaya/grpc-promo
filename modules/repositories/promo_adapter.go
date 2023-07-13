package repository

import "github.com/jmoiron/sqlx"

type PromoRepositoryPresenter interface {
}

type PromoRepository struct {
	db *sqlx.DB
}

func NewPromoRepository(db *sqlx.DB) *PromoRepository {
	return &PromoRepository{db}
}
