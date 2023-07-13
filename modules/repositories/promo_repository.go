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

func (p *PromoRepository) TotalData(table string) (int64, error) {
	var total int64
	query := `SELECT COUNT(*) FROM ` + table
	err := p.db.QueryRowx(query).Scan(&total)
	if err != nil {
		return 0, err
	}
	return total, nil
}

func (p *PromoRepository) GetAllOrder(offset int64, limit int64) ([]*promoProto.ListOrder, error) {
	var orders []*promoProto.ListOrder
	query := `
		SELECT * FROM orders LIMIT $1 OFFSET $2
	`
	rows, err := p.db.Queryx(query, limit, offset)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var order promoProto.ListOrder
		err = rows.Scan(&order.OrderId, &order.OrderAttr, &order.PromoCode)
		if err != nil {
			return nil, err
		}
		orders = append(orders, &order)
	}
	return orders, nil
}
