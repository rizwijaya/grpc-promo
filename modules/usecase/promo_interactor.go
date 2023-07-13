package usecase

import (
	"errors"
	promoProto "promo/modules/protobuf/pb"
	errVal "promo/pkg/errorVal"
	"promo/pkg/pagination"
)

func (p *PromoUseCase) GetPromoByCode(code string) (*promoProto.Promo, error) {
	return p.repository.GetPromoByCode(code)
}

func (p *PromoUseCase) Pagination(table string, pages int64, limit int64) (int64, *promoProto.Pagination, error) {
	total, err := p.repository.TotalData(table)
	if err != nil {
		return 0, nil, err
	}
	offset, pagination := pagination.Pages(total, pages, limit)
	if err != nil {
		return 0, nil, err
	}
	return offset, pagination, nil
}

func (p *PromoUseCase) GetAllOrder(offset int64, limit int64) ([]*promoProto.ListOrder, error) {
	return p.repository.GetAllOrder(offset, limit)
}

func (p *PromoUseCase) GetOrderById(id string) (*promoProto.Order, error) {
	order, err := p.repository.GetOrderById(id)
	if err != nil {
		return nil, err
	}

	if order.OrderId == "" {
		return nil, errors.New(errVal.ErrDataNotFound)
	}

	return order, nil
}

func (p *PromoUseCase) CreateOrder(orderAttr string, promoCode string) (*promoProto.ListOrder, error) {
	return p.repository.CreateOrder(orderAttr, promoCode)
}
