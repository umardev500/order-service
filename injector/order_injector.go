package injector

import (
	"order/app/delivery"
	"order/app/repository"
	"order/app/usecase"

	"go.mongodb.org/mongo-driver/mongo"
)

func NewOrderInjector(db *mongo.Database) *delivery.OrderDelivery {
	repo := repository.NewOrderRepository(db)
	usecase := usecase.NewOrderUsecase(repo)

	return delivery.NewOrderDelivery(usecase)
}
