package repository

import (
	promoProto "promo/modules/protobuf/pb"

	"github.com/google/uuid"
)

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
	query := `
		SELECT
			o.order_id, o.order_attr, o.promo_code, p.promo_id, 
			p.promo_code as promo_active, p.promo_attr 
		FROM
			orders o 
		INNER JOIN
			order_promo op
		ON o.order_id = op.order_id 
		INNER JOIN 
			promo p
		ON op.promo_id = p.promo_id 
		WHERE
			o.order_id = $1::uuid;
	`
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

func (p *PromoRepository) CreateOrder(orderAttr string, promoCode string) (*promoProto.ListOrder, error) {
	var order promoProto.ListOrder
	query := `
		INSERT INTO orders (
			order_id, order_attr, promo_code
		) VALUES ($1, $2, $3) 
		RETURNING order_id, order_attr, promo_code
	`
	err := p.db.QueryRowx(query, uuid.New(), orderAttr, promoCode).Scan(&order.OrderId, &order.OrderAttr, &order.PromoCode)
	if err != nil {
		return nil, err
	}
	return &order, nil
}
