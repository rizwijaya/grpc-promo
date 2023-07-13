package usecase

import (
	promoProto "promo/modules/protobuf/pb"
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
