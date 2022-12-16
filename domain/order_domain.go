package domain

import "order/pb"

type OrderUsecase interface {
	Save(req *pb.OrderCreateRequest) error
}

type OrderRepository interface {
	Save(req *pb.OrderCreateRequest, createdTime int64) error
}
