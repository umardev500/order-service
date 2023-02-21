run:
	go run main.go

proto:
	protoc --go_out=. --go_opt=paths=source_relative \
	--go-grpc_out=. --go-grpc_opt=paths=source_relative \
	pb/*.proto

clean:
	rm pb/*.pb.go

# grpcurl execution
create:
	grpcurl --plaintext -d '{"product": {"product_id": "16672232323", "name": "Trial 30 days", "price": 10, "duration": 30, "description": "This is the description"}}' localhost:5011 OrderService.Create

changeStatus:
	grpcurl --plaintext -d '{"order_id": "1671193878472", "status": "settlement"}' localhost:5011 OrderService.ChangeStatus

findOne:
	grpcurl --plaintext -d '{"order_id": "1676961217179686338"}' localhost:5011 OrderService.FindOne

findAll:
	grpcurl --plaintext -d '{"status": "pending"}' localhost:5011 OrderService.FindAll

findAllByCustomer:
	grpcurl --plaintext -d '{"user_id": "1667292823233", "status": "pending"}' localhost:5011 OrderService.FindAll

sum:
	grpcurl --plaintext -d '{"status": "settlement"}' localhost:5011 OrderService.SumIncome
