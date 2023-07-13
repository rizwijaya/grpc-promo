package controllers

import (
	"context"
	"log"
	promoProto "promo/modules/protobuf/pb"
	errVal "promo/pkg/errorVal"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (p *PromoController) GetPromoByCode(ctx context.Context, code *promoProto.Code) (*promoProto.Promo, error) {
	//Validate if input invalid
	if code.PromoCode == "" {
		return nil, status.Errorf(codes.InvalidArgument, errVal.ErrInvalidArgument("Kode Promo"))
	}

	promo, err := p.usecase.GetPromoByCode(code.PromoCode)
	if err != nil {
		if err.Error() == errVal.ErrDataNotFound {
			return nil, status.Errorf(codes.NotFound, errVal.ErrNotFound("Promo"))
		}
		log.Println(err)
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return promo, err
}

func (p *PromoController) GetOrderList(ctx context.Context, pageParam *promoProto.Pages) (*promoProto.Orders, error) {
	var (
		page  int64 = 1
		limit int64 = 5
	)
	if pageParam.GetPage() != 0 || pageParam.GetLimit() != 0 {
		page = pageParam.GetPage()
		limit = pageParam.GetLimit()
	}
	offset, pagination, err := p.usecase.Pagination("orders", page, limit)
	if err != nil {
		log.Println(err)
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	order, err := p.usecase.GetAllOrder(offset, limit)
	if err != nil {
		log.Println(err)
		if err.Error() == errVal.ErrDataNotFound {
			return nil, status.Errorf(codes.NotFound, errVal.ErrNotFound("Order"))
		}
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	response := &promoProto.Orders{
		Orders:     order,
		Pagination: pagination,
	}
	return response, nil
}

func (p *PromoController) GetOrderDetail(ctx context.Context, id *promoProto.OrderId) (*promoProto.Order, error) {
	if id.OrderId == "" {
		return nil, status.Errorf(codes.InvalidArgument, errVal.ErrInvalidArgument("Id Order"))
	}

	order, err := p.usecase.GetOrderById(id.OrderId)
	if err != nil {
		if err.Error() == errVal.ErrDataNotFound || err.Error() == errVal.ErrInvalidUUID(id.OrderId) {
			return nil, status.Errorf(codes.NotFound, errVal.ErrNotFound("Order"))
		}

		log.Println(err)
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return order, err
}

func (p *PromoController) CreateOrder(ctx context.Context, orderParam *promoProto.CreateOrderRequest) (*promoProto.ListOrder, error) {
	if orderParam.OrderAttr == "" || orderParam.PromoCode == "" {
		return nil, status.Errorf(codes.InvalidArgument, errVal.ErrInvalidArgument("Order Attribut/Kode Promo"))
	}
	order, err := p.usecase.CreateOrder(orderParam.OrderAttr, orderParam.PromoCode)
	if err != nil {
		if err.Error() == errVal.ErrPromoCodeinOrderNotFound {
			return nil, status.Errorf(codes.NotFound, errVal.ErrNotFound("Kode Promo"))
		}
		log.Println(err)
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return order, nil
}
