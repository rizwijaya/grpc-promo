package repository

import promoProto "promo/modules/protobuf/pb"

func (p *PromoRepository) GetPromoByCode(code string) (*promoProto.Promo, error) {
	var promo promoProto.Promo
	query := `SELECT * FROM promo WHERE promo_code = $1`
	err := p.db.QueryRowx(query, code).Scan(&promo.PromoId, &promo.PromoAttr, &promo.PromoCode)
	if err != nil {
		return nil, err
	}
	return &promo, nil
}
