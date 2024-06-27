gen-product:
	protoc --go_out=. --go-grpc_out=. protos/protoduct/prtoduct.proto