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

func (pu *OrderUsecase) SumIncome(ctx context.Context, req *pb.OrderSumIncomeRequest) (res int64, err error) {
	res, err = pu.repository.SumIncome(ctx, req)

	return
}

func (pu *OrderUsecase) FindAll(ctx context.Context, req *pb.OrderFindAllRequest) (res *pb.OrderFindAllResponse, err error) {
	res, err = pu.repository.FindAll(ctx, req)

	return
}

func (pu *OrderUsecase) FindOne(ctx context.Context, req *pb.OrderFindOneRequest) (res *pb.Order, err error) {
	res, err = pu.repository.FindOne(ctx, req)

	return
}

func (pu *OrderUsecase) ChangeStatus(ctx context.Context, req *pb.OrderChangeStatus) (affected bool, err error) {
	updatedTime := time.Now().Unix()
	affected, err = pu.repository.ChangeStatus(ctx, req, updatedTime)

	return
}

func (pu *OrderUsecase) Save(ctx context.Context, req *pb.OrderCreateRequest) (err error) {
	t := time.Now()
	createdTime := t.Unix()
	err = pu.repository.Save(ctx, req, createdTime)

	return
}
