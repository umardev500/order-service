package delivery

import (
	"context"
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

func (pd *OrderDelivery) ChangeStatus(ctx context.Context, req *pb.OrderChangeStatus) (res *pb.OperationResponse, err error) {
	affected, err := pd.usecase.ChangeStatus(req)

	return &pb.OperationResponse{IsAffected: affected}, err
}

func (pd *OrderDelivery) Create(ctx context.Context, req *pb.OrderCreateRequest) (res *pb.Empty, err error) {
	err = pd.usecase.Save(req)
	res = &pb.Empty{}

	return
}
