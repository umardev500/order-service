package domain

import "order/pb"

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

type Order struct {
	OrderId   string         `bson:"order_id"`
	Buyer     OrderBuyer     `bson:"buyer"`
	Product   []OrderProduct `bson:"product"`
	Status    string         `bson:"status"`
	CreatedAt int64          `bson:"created_at"`
	UpdatedAt int64          `bson:"updated_at"`
}

type OrderUsecase interface {
	Save(req *pb.OrderCreateRequest) error
	ChangeStatus(req *pb.OrderChangeStatus) (affected bool, err error)
	FindOne(req *pb.OrderFindOneRequest) (res *pb.Order, err error)
	FindAll(req *pb.OrderFindAllRequest) (res *pb.OrderFindAllResponse, err error)
}

type OrderRepository interface {
	Save(req *pb.OrderCreateRequest, generatedId string, createdTime int64) error
	ChangeStatus(req *pb.OrderChangeStatus, updatedTime int64) (affected bool, err error)
	FindOne(req *pb.OrderFindOneRequest) (res *pb.Order, err error)
	FindAll(req *pb.OrderFindAllRequest) (orders *pb.OrderFindAllResponse, err error)
}
