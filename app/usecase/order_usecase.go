package usecase

import (
	"order/domain"
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
