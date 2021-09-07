package repository

import (
	"context"
	"math"
	"order/domain"
	"order/pb"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func (pr *OrderRepository) parseOrderResponse(each domain.Order) (order *pb.Order) {
	products := []*pb.OrderProduct{}
	for _, val := range each.Product {
		products = append(products, &pb.OrderProduct{
			ProductId:   val.ProductId,
			Name:        val.Name,
			Price:       val.Price,
			Duration:    val.Duration,
			Description: val.Description,
		})
	}

	order = &pb.Order{
		OrderId: each.OrderId,
		Buyer: &pb.OrderBuyer{
			CustomerId: each.Buyer.CustomerId,
			Name:       each.Buyer.Name,
			User:       each.Buyer.User,
		},
		Product:   products,
		Status:    each.Status,
		CreatedAt: each.CreatedAt,
		UpdatedAt: each.UpdatedAt,
	}

	return
}

// Template
// func (pr *OrderRepository) {
// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 	defer cancel()
// }

func (pr *OrderRepository) FindAll(req *pb.OrderFindAllRequest) (orders *pb.OrderFindAllResponse, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	orders = &pb.OrderFindAllResponse{}
	s := req.Search

	status := bson.M{"status": req.Status}
	if req.Status == "none" && req.Status != "deleted" {
		status = bson.M{"status": bson.M{"$ne": nil}}
	}

	filter := bson.M{
		"$or": []bson.M{
			{
				"order_id": bson.M{
					"$regex": primitive.Regex{
						Pattern: s,
						Options: "i",
					},
				},
			},
		},
		"$and": []bson.M{
			status,
		},
	}
	findOpt := options.Find()

	if req.Sort == "desc" {
		findOpt.SetSort(bson.M{"order_id": -1})
	}

	page := req.Page
	perPage := req.PerPage
	offset := page * perPage

	findOpt.SetSkip(offset)
	findOpt.SetLimit(perPage)

	cur, err := pr.orders.Find(ctx, filter, findOpt)
	if err != nil {
		return
	}

	defer cur.Close(ctx)

	for cur.Next(ctx) {
		each := domain.Order{}
		err = cur.Decode(&each)
		if err != nil {
			return
		}

		order := pr.parseOrderResponse(each)

		orders.Orders = append(orders.Orders, order)
	}

	rows, _ := pr.orders.CountDocuments(ctx, filter)

	dataSize := int64(len(orders.Orders))
	orders.Rows = rows
	orders.Pages = int64(math.Ceil(float64(rows) / float64(perPage)))
	if dataSize < 1 {
		orders.Pages = 0
	} else if perPage == 0 {
		orders.Pages = 1
	}

	orders.PerPage = perPage
	orders.ActivePage = page + 1
	if dataSize < 1 {
		orders.ActivePage = 0
	}
	orders.Total = dataSize

	return
}

func (pr *OrderRepository) FindOne(req *pb.OrderFindOneRequest) (res *pb.Order, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var order domain.Order

	filter := bson.M{"order_id": req.OrderId}
	err = pr.orders.FindOne(ctx, filter).Decode(&order)

	res = pr.parseOrderResponse(order)

	return
}

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

	products := []bson.D{}

	for _, val := range req.Product {
		each := bson.D{
			{Key: "product_id", Value: val.ProductId},
			{Key: "name", Value: val.Name},
			{Key: "price", Value: val.Price},
			{Key: "duration", Value: val.Duration},
			{Key: "description", Value: val.Description},
		}

		products = append(products, each)
	}

	data := bson.D{
		{Key: "order_id", Value: generatedId},
		{Key: "buyer", Value: buyer},
		{Key: "product", Value: products},
		{Key: "status", Value: "pending"},
		{Key: "created_at", Value: createdTime},
	}

	_, err = pr.orders.InsertOne(ctx, data)

	return
}
