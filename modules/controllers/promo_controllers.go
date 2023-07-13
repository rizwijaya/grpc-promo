package controllers

import (
	"context"
	"log"
	promoProto "promo/modules/protobuf/pb"
	errVal "promo/pkg/errorVal"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (p *PromoUseCase) GetPromoByCode(ctx context.Context, code *promoProto.Code) (*promoProto.Promo, error) {
	//Validate if input invalid
	if code.PromoCode == "" {
		return nil, status.Errorf(codes.InvalidArgument, errVal.ErrInvalidArgument("Kode Promo"))
	}

	promo, err := p.usecase.GetPromoByCode(code.PromoCode)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return nil, status.Errorf(codes.NotFound, errVal.ErrNotFound("Promo"))
		}
		log.Println(err)
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return promo, err
}
