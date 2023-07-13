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

func (p *PromoRepository) GetOrderById(id string) (*promoProto.Order, error) {
	var order promoProto.Order
	query :=
		`select
		o.order_id, o.order_attr, o.promo_code, p.promo_id, p.promo_code as promo_active, p.promo_attr 
	from
		orders o 
	inner join
		order_promo op
	on o.order_id = op.order_id 
	inner join 
		promo p
	on op.promo_id = p.promo_id 
	where
		o.order_id = $1::uuid;`
	row, err := p.db.Queryx(query, id)
	if err != nil {
		return nil, err
	}
	for row.Next() {
		var promo promoProto.Promo
		if err := row.Scan(&order.OrderId, &order.OrderAttr, &order.PromoCode, &promo.PromoId, &promo.PromoCode, &promo.PromoAttr); err != nil {
			return nil, err
		}
		order.Promos = append(order.Promos, &promo)
	}
	return &order, nil
}
