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
// func (pr *OrderRepository) {
// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 	defer cancel()
// }

func (pr *OrderRepository) ChangeStatus(req *pb.OrderChangeStatus, updatedTime int64) (affected bool, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"order_id": req.OrderId}
	data := bson.M{"status": req.Status, "updated_at": updatedTime}
	set := bson.M{"$set": data}

	resp, err := pr.orders.UpdateOne(ctx, filter, set)

	if resp.ModifiedCount > 0 {
		affected = true
	}

	return
}

func (pr *OrderRepository) Save(req *pb.OrderCreateRequest, generatedId string, createdTime int64) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	buyer := bson.D{
		{Key: "customer_id", Value: req.Buyer.CustomerId},
		{Key: "name", Value: req.Buyer.Name},
		{Key: "user", Value: req.Buyer.User},
	}

	product := []bson.D{}

	for _, val := range req.Product {
		each := bson.D{
			{Key: "product_id", Value: val.ProductId},
			{Key: "name", Value: val.Name},
			{Key: "price", Value: val.Price},
			{Key: "duration", Value: val.Duration},
			{Key: "description", Value: val.Description},
		}

		product = append(product, each)
	}

	data := bson.D{
		{Key: "order_id", Value: generatedId},
		{Key: "buyer", Value: buyer},
		{Key: "product", Value: product},
		{Key: "status", Value: "pending"},
		{Key: "created_at", Value: createdTime},
	}

	_, err = pr.orders.InsertOne(ctx, data)

	return
}
