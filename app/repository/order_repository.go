package repository

import (
	"order/domain"

	"go.mongodb.org/mongo-driver/mongo"
)

type OrderRepository struct {
	db     *mongo.Database
	Orders *mongo.Collection
}

func NewOrderRepository(db *mongo.Database) domain.OrderRepository {
	return &OrderRepository{
		db:     db,
		Orders: db.Collection("orders"),
	}
}

// Template
// func (pr *OrderRepository) {}
