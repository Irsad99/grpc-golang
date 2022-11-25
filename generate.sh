protoc --proto_path=./protos ./protos/product.proto \
    --proto_path=./protos/libs \
    --plugin=$(go env GOPATH)/bin/protoc-gen-go \
    --plugin=$(go env GOPATH)/bin/protoc-gen-govalidators \
    --go_out=./cmd/pb --go_opt paths=source_relative \
    --govalidators_out=./cmd/pb --govalidators_opt paths=source_relative
    
protoc --proto_path=./protos ./protos/product.proto \
    --proto_path=./protos/libs \
    --proto_path=./vendor \
    --plugin=$(go env GOPATH)/bin/protoc-gen-grpc-gateway \
    --plugin=$(go env GOPATH)/bin/protoc-gen-openapiv2 \
    --plugin=$(go env GOPATH)/bin/protoc-gen-go-grpc \
    --go-grpc_out=./cmd/pb --go-grpc_opt paths=source_relative \
    --grpc-gateway_out ./cmd/pb \
    --grpc-gateway_opt allow_delete_body=true,logtostderr=true,paths=source_relative,repeated_path_param_separator=ssv \
    --openapiv2_out ./protos \
    --openapiv2_opt logtostderr=true,repeated_path_param_separator=ssv