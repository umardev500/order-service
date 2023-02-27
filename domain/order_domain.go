package domain

import (
	"context"
	"order/pb"
)

type OrderBuyer struct {
	CustomerId string `bson:"customer_id"`
	Name       string `bson:"name"`
	User       string `bson:"user"`
}

type OrderProduct struct {
	ProductId   string `bson:"product_id"`
	Name        string `bson:"name"`
	Price       int64  `bson:"price"`
	Duration    int64  `bson:"duration"`
	Description string `bson:"description"`
}

type OrderPayment struct {
	PaymentType string `bson:"payment_type"`
	OrderID     string `bson:"order_id"`
	Bank        string `bson:"bank"`
	VaNumber    string `bson:"va_number"`
	GrossAmount int64  `bson:"gross_amount"`
}

type Order struct {
	OrderId        string       `bson:"order_id"`
	Buyer          OrderBuyer   `bson:"buyer"`
	Product        OrderProduct `bson:"product"`
	Payment        OrderPayment `bson:"payment"`
	Status         string       `bson:"status"`
	CreatedAt      int64        `bson:"created_at"`
	UpdatedAt      int64        `bson:"updated_at"`
	SettlementTime int64        `bson:"settlement_time"`
}

type OrderUsecase interface {
	Save(ctx context.Context, req *pb.OrderCreateRequest) error
	ChangeStatus(ctx context.Context, req *pb.OrderChangeStatus) (affected bool, err error)
	FindOne(ctx context.Context, req *pb.OrderFindOneRequest) (res *pb.OrderFindOneResponse, err error)
	FindAll(ctx context.Context, req *pb.OrderFindAllRequest) (res *pb.OrderFindAllResponse, err error)
	SumIncome(ctx context.Context, req *pb.OrderSumIncomeRequest) (res int64, err error)
	Cancel(ctx context.Context, req *pb.OrderCancelRequest) (res *pb.OperationResponse, err error)
}

type OrderRepository interface {
	Save(ctx context.Context, req *pb.OrderCreateRequest, createdTime int64) error
	ChangeStatus(ctx context.Context, req *pb.OrderChangeStatus, updatedTime int64) (affected bool, err error)
	FindOne(ctx context.Context, req *pb.OrderFindOneRequest) (res *pb.OrderFindOneResponse, err error)
	FindAll(ctx context.Context, req *pb.OrderFindAllRequest) (orders *pb.OrderFindAllResponse, err error)
	SumIncome(ctx context.Context, req *pb.OrderSumIncomeRequest) (res int64, err error)
	Cancel(ctx context.Context, req *pb.OrderCancelRequest) (res *pb.OperationResponse, err error)
}
