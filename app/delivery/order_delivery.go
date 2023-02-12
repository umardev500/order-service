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

func (pd *OrderDelivery) SumIncome(ctx context.Context, req *pb.OrderSumIncomeRequest) (res *pb.OrderSumResponse, err error) {
	total, err := pd.usecase.SumIncome()
	if err == mongo.ErrNoDocuments {
		res = &pb.OrderSumResponse{IsEmpty: true}
		err = nil

		return
	}

	res = &pb.OrderSumResponse{Payload: &pb.OrderSumPayload{Total: total}}

	return
}

func (pd *OrderDelivery) FindAll(ctx context.Context, req *pb.OrderFindAllRequest) (res *pb.OrderFindAllResponse, err error) {
	res, err = pd.usecase.FindAll(req)

	return
}

func (pd *OrderDelivery) FindOne(ctx context.Context, req *pb.OrderFindOneRequest) (res *pb.Order, err error) {
	res, err = pd.usecase.FindOne(req)

	return
}

func (pd *OrderDelivery) ChangeStatus(ctx context.Context, req *pb.OrderChangeStatus) (res *pb.OperationResponse, err error) {
	affected, err := pd.usecase.ChangeStatus(req)

	return &pb.OperationResponse{IsAffected: affected}, err
}

func (pd *OrderDelivery) Create(ctx context.Context, req *pb.OrderCreateRequest) (res *pb.Empty, err error) {
	err = pd.usecase.Save(req)
	res = &pb.Empty{}

	return
}
