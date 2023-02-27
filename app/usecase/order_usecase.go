package usecase

import (
	"context"
	"order/domain"
	"order/pb"
	"time"
)

// OrderUsecase defines the use case for managing Orders.
type OrderUsecase struct {
	// repository is the underlying repository for storing Orders.
	repository domain.OrderRepository
}

// NewOrderUsecase creates a new OrderUsecase with the given repository.
func NewOrderUsecase(repo domain.OrderRepository) domain.OrderUsecase {
	return &OrderUsecase{
		repository: repo,
	}
}

// Template
// func (pu *OrderUsecase) {}
func (o *OrderUsecase) Cancel(ctx context.Context, req *pb.OrderCancelRequest) (res *pb.OperationResponse, err error) {
	return o.repository.Cancel(ctx, req)
}

func (o *OrderUsecase) SumIncome(ctx context.Context, req *pb.OrderSumIncomeRequest) (res int64, err error) {
	res, err = o.repository.SumIncome(ctx, req)

	return
}

func (o *OrderUsecase) FindAll(ctx context.Context, req *pb.OrderFindAllRequest) (res *pb.OrderFindAllResponse, err error) {
	res, err = o.repository.FindAll(ctx, req)

	return
}

func (o *OrderUsecase) FindOne(ctx context.Context, req *pb.OrderFindOneRequest) (res *pb.OrderFindOneResponse, err error) {
	res, err = o.repository.FindOne(ctx, req)

	return
}

func (o *OrderUsecase) ChangeStatus(ctx context.Context, req *pb.OrderChangeStatus) (affected bool, err error) {
	updatedTime := time.Now().Unix()
	affected, err = o.repository.ChangeStatus(ctx, req, updatedTime)

	return
}

func (o *OrderUsecase) Save(ctx context.Context, req *pb.OrderCreateRequest) (err error) {
	t := time.Now()
	createdTime := t.Unix()
	err = o.repository.Save(ctx, req, createdTime)

	return
}
