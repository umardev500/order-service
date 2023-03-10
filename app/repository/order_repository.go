package repository

import (
	"context"
	"math"
	"order/domain"
	"order/helper"
	"order/pb"
	"order/variable"
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
	product := &pb.OrderProduct{
		ProductId:   each.Product.ProductId,
		Name:        each.Product.Name,
		Price:       each.Product.Price,
		Duration:    each.Product.Duration,
		Description: each.Product.Name,
	}

	payment := &pb.OrderPayment{
		PaymentType: each.Payment.PaymentType,
		OrderId:     each.Payment.PaymentType,
		Bank:        each.Payment.Bank,
		VaNumber:    each.Payment.VaNumber,
		GrossAmount: each.Payment.GrossAmount,
	}

	order = &pb.Order{
		OrderId: each.OrderId,
		Buyer: &pb.OrderBuyer{
			CustomerId: each.Buyer.CustomerId,
			Name:       each.Buyer.Name,
			User:       each.Buyer.User,
		},
		Product:        product,
		Payment:        payment,
		Status:         each.Status,
		CreatedAt:      each.CreatedAt,
		UpdatedAt:      each.UpdatedAt,
		SettlementTime: each.SettlementTime,
		TrxTime:        each.TrxTime,
		PayExp:         each.PayExp,
	}

	return
}

// Template
// func (pr *OrderRepository) {
// }

func (o *OrderRepository) Cancel(ctx context.Context, req *pb.OrderCancelRequest) (res *pb.OperationResponse, err error) {
	isAffected := true
	updatedTime := time.Now().UTC().Unix()

	filter := bson.M{
		"status":            "pending",
		"order_id":          req.OrderId,
		"buyer.customer_id": req.UserId,
	}
	payload := bson.M{"status": "cancel", "updated_at": updatedTime}
	set := bson.M{"$set": payload}
	resp, err := o.orders.UpdateOne(ctx, filter, set)
	if resp.ModifiedCount < 1 {
		isAffected = false
	}

	res = &pb.OperationResponse{IsAffected: isAffected}

	return
}

func (o *OrderRepository) SumIncome(ctx context.Context, req *pb.OrderSumIncomeRequest) (res int64, err error) {
	status := "settlement"
	if req.Status != "" {
		status = req.Status
	}

	pipeline := []bson.M{
		{
			"$match": bson.M{
				"status": status,
			},
		},
		{
			"$group": bson.M{
				"_id": nil,
				"total": bson.M{
					"$sum": "$payment.gross_amount",
				},
			},
		},
	}

	cur, err := o.orders.Aggregate(ctx, pipeline)
	if err != nil {
		return
	}

	defer cur.Close(ctx)

	var results []struct {
		Total int64 `bson:"total"`
	}
	cur.All(ctx, &results)

	if len(results) == 0 {
		err = mongo.ErrNoDocuments
		return
	}

	res = results[0].Total

	return
}

func (o *OrderRepository) FindAll(ctx context.Context, req *pb.OrderFindAllRequest) (result *pb.OrderFindAllResponse, err error) {
	result = &pb.OrderFindAllResponse{}
	s := req.Search

	status := bson.M{"status": req.Status}
	if (req.Status == "none" || req.Status == "") && req.Status != "deleted" {
		status = bson.M{"status": bson.M{"$ne": nil}}
	}

	expired := bson.M{}
	if req.Status == "pending" {
		expired = bson.M{"pay_exp": bson.M{"$gt": helper.GetTime(nil)}}
	}

	if req.Status == "cancel" {
		status = bson.M{"$or": []bson.M{
			{
				"status": "cancel",
			},
			{
				"$and": []bson.M{
					{"pay_exp": bson.M{"$lt": helper.GetTime(nil)}},
					{"status": bson.M{"$ne": "settlement"}},
				},
			},
		}}
	}

	userId := bson.M{}
	if req.UserId != "" {
		userId = bson.M{"buyer.customer_id": req.UserId}
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
			{
				"status": bson.M{
					"$regex": primitive.Regex{
						Pattern: s,
						Options: "i",
					},
				},
			},
			{
				"buyer.name": bson.M{
					"$regex": primitive.Regex{
						Pattern: s,
						Options: "i",
					},
				},
			},
			{
				"buyer.user": bson.M{
					"$regex": primitive.Regex{
						Pattern: s,
						Options: "i",
					},
				},
			},
			{
				"product.name": bson.M{
					"$regex": primitive.Regex{
						Pattern: s,
						Options: "i",
					},
				},
			},
		},
		"$and": []bson.M{
			status,
			userId,
			expired,
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

	var orders []*pb.Order
	if !req.CountOnly {
		cur, err := o.orders.Find(ctx, filter, findOpt)
		if err != nil {
			return nil, err
		}

		defer cur.Close(ctx)

		for cur.Next(ctx) {
			each := domain.Order{}
			err = cur.Decode(&each)
			if err != nil {
				return nil, err
			}

			order := o.parseOrderResponse(each)

			orders = append(orders, order)
		}

		if len(orders) < 1 {
			result.IsEmpty = true

			return result, nil
		}
	}

	result.Payload = &pb.OrderFindAllPayload{
		Orders: orders,
	}

	rows, _ := o.orders.CountDocuments(ctx, filter)

	dataSize := int64(len(result.Payload.Orders))
	result.Payload.Rows = rows
	result.Payload.Pages = int64(math.Ceil(float64(rows) / float64(perPage)))
	if dataSize < 1 {
		result.Payload.Pages = 0
	} else if perPage == 0 {
		result.Payload.Pages = 1
	}

	result.Payload.PerPage = perPage
	result.Payload.ActivePage = page + 1
	if dataSize < 1 {
		result.Payload.ActivePage = 0
	}
	result.Payload.Total = dataSize

	return
}

func (o *OrderRepository) FindOne(ctx context.Context, req *pb.OrderFindOneRequest) (res *pb.OrderFindOneResponse, err error) {
	var order domain.Order

	filter := bson.M{"order_id": req.OrderId}
	err = o.orders.FindOne(ctx, filter).Decode(&order)
	if err == mongo.ErrNoDocuments {
		res = &pb.OrderFindOneResponse{IsEmpty: true}
		err = nil
		return
	}

	res = &pb.OrderFindOneResponse{
		Payload: o.parseOrderResponse(order),
	}

	return
}

func (o *OrderRepository) ChangeStatus(ctx context.Context, req *pb.OrderChangeStatus, updatedTime int64) (affected bool, err error) {
	filter := bson.M{"payment.order_id": req.OrderId}
	if req.Status == variable.PayementStatusExpire {
		filter["status"] = bson.M{"$ne": variable.PayementStatusCancel}
	}
	data := bson.M{"status": req.Status, "updated_at": updatedTime}
	if req.Status == "settlement" {
		data["settlement_time"] = req.SettlementTime
	}

	set := bson.M{"$set": data}

	resp, err := o.orders.UpdateOne(ctx, filter, set)

	if resp.ModifiedCount > 0 {
		affected = true
	}

	return
}

func (o *OrderRepository) Save(ctx context.Context, req *pb.OrderCreateRequest, createdTime int64) (err error) {
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

	payment := bson.D{}
	if req.Payment != nil {
		payment = bson.D{
			{Key: "payment_type", Value: req.Payment.PaymentType},
			{Key: "order_id", Value: req.Payment.OrderId},
			{Key: "bank", Value: req.Payment.Bank},
			{Key: "va_number", Value: req.Payment.VaNumber},
			{Key: "gross_amount", Value: req.Payment.GrossAmount},
		}
	}

	data := bson.D{
		{Key: "order_id", Value: req.Payment.OrderId},
		{Key: "buyer", Value: buyer},
		{Key: "product", Value: product},
		{Key: "payment", Value: payment},
		{Key: "status", Value: "pending"},
		{Key: "created_at", Value: createdTime},
		{Key: "trx_time", Value: req.TrxTime},
		{Key: "pay_exp", Value: req.PayExp},
	}

	filteredData := bson.D{}
	for _, x := range data {
		if x.Key == "payment" && req.Payment == nil {
			continue
		}

		filteredData = append(filteredData, x)
	}

	_, err = o.orders.InsertOne(ctx, filteredData)

	return
}
