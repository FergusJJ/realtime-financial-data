generate:
	protoc --proto_path=internal/proto --go_out=internal/proto/ --go-grpc_out=internal/proto/ internal/proto/*.proto
