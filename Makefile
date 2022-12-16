run:
	go run main.go

proto:
	protoc --proto_path=pb --go_out=pb --go_opt=paths=source_relative \
	--go-grpc_out=pb --go-grpc_opt=paths=source_relative \
	pb/*.proto

clean:
	rm pb/*.pb.go

# grpcurl execution
create:
	grpcurl --plaintext -d '{"buyer": {"customer_id":"16678237623", "name": "SMK Walisongo", "user": "walisongo"}, "product": {"product_id": "16672232323", "name": "Trial 30 days", "price": 10, "duration": 30, "description": "This is the description"}}' localhost:5011 OrderService.Create

changeStatus:
	grpcurl --plaintext -d '{"order_id": "1671193878472"}' localhost:5011 OrderService.Delete
