package delivery

import (
	"order/domain"
	"order/pb"
)

type OrderDelivery struct {
	usecase domain.OrderUsecase
	pb.UnimplementedOrderServiceServer
}

func NewOrderDelivery(usecase domain.OrderUsecase) *OrderDelivery {
	return &OrderDelivery{
		usecase: usecase,
	}
}

// Template
//func (pd *OrderDelivery) Delete(ctx context.Context, req *pb.) (res *pb., err error) {}
