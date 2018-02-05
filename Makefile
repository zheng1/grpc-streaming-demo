all: generate

generate:
	@echo "--> Generating Go files"
	protoc -I protobuf/ --go_out=plugins=grpc:protobuf/ protobuf/pilot.proto
	@echo ""
