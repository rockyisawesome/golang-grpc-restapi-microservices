.PHONY: protos

protos:
	 protoc -I protos/ protos/currency.proto --go_out=./ --go-grpc_out=./
grpcurl1:
	grpcurl -plaintext localhost:9090 Currency.GetRate
grpcurl2:
	grpcurl -plaintext localhost:9090 HelloWorld.GetHello
getlist:
	grpcurl -plaintext localhost:9090 list