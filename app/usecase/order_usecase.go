package usecase

import (
	"order/domain"
	"order/pb"
	"strconv"
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
func (pu *OrderUsecase) FindOne(req *pb.OrderFindOneRequest) (res *pb.Order, err error) {
	res, err = pu.repository.FindOne(req)

	return
}

func (pu *OrderUsecase) ChangeStatus(req *pb.OrderChangeStatus) (affected bool, err error) {
	updatedTime := time.Now().Unix()
	affected, err = pu.repository.ChangeStatus(req, updatedTime)

	return
}

func (pu *OrderUsecase) Save(req *pb.OrderCreateRequest) (err error) {
	t := time.Now()
	createdTime := t.Unix()
	generatedId := strconv.Itoa(int(t.UnixMilli()))
	err = pu.repository.Save(req, generatedId, createdTime)

	return
}
