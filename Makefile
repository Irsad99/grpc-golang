generate:
	protoc -Iprotos -Iprotos/libs \
	--go_opt=module=grpc --go_out=. \
	--go-grpc_opt=module=grpc --go-grpc_out=. \
	--grpc-gateway_opt=module=grpc --grpc-gateway_out=. --grpc-gateway_opt=logtostderr=true \
	--validate_opt=module=grpc --validate_out=. --validate_opt=lang=go \
	./protos/*.proto

run:
	go run cmd/main.go