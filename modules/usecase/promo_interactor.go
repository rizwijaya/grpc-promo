package usecase

import promoProto "promo/modules/protobuf/pb"

func (p *PromoUseCase) GetPromoByCode(code string) (*promoProto.Promo, error) {
	return p.repository.GetPromoByCode(code)
}
