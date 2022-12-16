package repository

import (
	"context"
	"order/domain"
	"order/pb"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type OrderRepository struct {
	db     *mongo.Database
	orders *mongo.Collection
}

func NewOrderRepository(db *mongo.Database) domain.OrderRepository {
	return &OrderRepository{
		db:     db,
		orders: db.Collection("orders"),
	}
}

// Template
// func (pr *OrderRepository) {}

func (pr *OrderRepository) Save(req *pb.OrderCreateRequest, createdTime int64) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	buyer := bson.D{
		{Key: "customer_id", Value: req.Buyer.CustomerId},
		{Key: "name", Value: req.Buyer.Name},
		{Key: "user", Value: req.Buyer.User},
	}

	product := bson.D{
		{Key: "product_id", Value: req.Product.ProductId},
		{Key: "name", Value: req.Product.Name},
		{Key: "price", Value: req.Product.Price},
		{Key: "duration", Value: req.Product.Duration},
		{Key: "description", Value: req.Product.Description},
	}

	data := bson.D{
		{Key: "buyer", Value: buyer},
		{Key: "product", Value: product},
		{Key: "status", Value: "pending"},
		{Key: "created_at", Value: createdTime},
	}

	_, err = pr.orders.InsertOne(ctx, data)

	return
}
