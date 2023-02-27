package delivery

import (
	"context"
	"order/domain"
	"order/pb"

	"go.mongodb.org/mongo-driver/mongo"
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
// func (pd *OrderDelivery) Delete(ctx context.Context, req *pb.) (res *pb., err error) {}

func (o *OrderDelivery) SumIncome(ctx context.Context, req *pb.OrderSumIncomeRequest) (res *pb.OrderSumResponse, err error) {
	total, err := o.usecase.SumIncome(ctx, req)
	if err == mongo.ErrNoDocuments {
		res = &pb.OrderSumResponse{IsEmpty: true}
		err = nil

		return
	}

	res = &pb.OrderSumResponse{Payload: &pb.OrderSumPayload{Total: total}}

	return
}

func (o *OrderDelivery) Cancel(ctx context.Context, req *pb.OrderCancelRequest) (res *pb.OperationResponse, err error) {
	res, err = o.usecase.Cancel(ctx, req)
	return
}

func (o *OrderDelivery) FindAll(ctx context.Context, req *pb.OrderFindAllRequest) (res *pb.OrderFindAllResponse, err error) {
	res, err = o.usecase.FindAll(ctx, req)

	return
}

func (o *OrderDelivery) FindOne(ctx context.Context, req *pb.OrderFindOneRequest) (res *pb.OrderFindOneResponse, err error) {
	res, err = o.usecase.FindOne(ctx, req)

	return
}

func (o *OrderDelivery) ChangeStatus(ctx context.Context, req *pb.OrderChangeStatus) (res *pb.OperationResponse, err error) {
	affected, err := o.usecase.ChangeStatus(ctx, req)

	return &pb.OperationResponse{IsAffected: affected}, err
}

func (o *OrderDelivery) Create(ctx context.Context, req *pb.OrderCreateRequest) (res *pb.Empty, err error) {
	err = o.usecase.Save(ctx, req)
	res = &pb.Empty{}

	return
}
