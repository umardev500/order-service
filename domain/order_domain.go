package domain

import "order/pb"

type OrderUsecase interface {
	Save(req *pb.OrderCreateRequest) error
	ChangeStatus(req *pb.OrderChangeStatus) (affected bool, err error)
}

type OrderRepository interface {
	Save(req *pb.OrderCreateRequest, generatedId string, createdTime int64) error
	ChangeStatus(req *pb.OrderChangeStatus, updatedTime int64) (affected bool, err error)
}
