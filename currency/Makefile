.PHONY: protos

protos:
	 protoc -I protos/ protos/currency.proto --go_out=./ --go-grpc_out=./
grpcurl:
	grpcurl -plaintext localhost:9090 Currency.GetRate
getlist:
	grpcurl -plaintext localhost:9090 list