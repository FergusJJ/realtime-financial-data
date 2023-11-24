generate:
	protoc --proto_path=internal/proto --go_out=internal/proto/ --go-grpc_out=internal/proto/ internal/proto/*.proto

cert:
	cd cert; ./gen.sh; cd ..

.PHONY: gen clean server client test cert