gen-product:
	protoc --go_out=. --go-grpc_out=. protos/protoduct/prtoduct.proto


DB_URL := "postgres://postgres:+_+diyor2005+_+@localhost:5432/productdb?sslmode=disable"

migrate-file:
	migrate create -ext sql -dir migrations/ -seq products_table


migrate-up:
	migrate -path migrations -database $(DB_URL) -verbose up

migrate-down:
	migrate -path migrations -database $(DB_URL) -verbose down

migrate-force:
	migrate -path migrations -database $(DB_URL) -verbose force 1