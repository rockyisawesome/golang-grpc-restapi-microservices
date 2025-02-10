package main

import (
	protos "currency/protos/currency"
	sercur "currency/server"
	"net"
	"os"

	"github.com/hashicorp/go-hclog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	// logger
	var loggs hclog.Logger = hclog.Default()

	serverPort := "localhost:9090"

	loggs.Info("Starting the application in port ", "port", serverPort)

	//create a grpc server
	grpcServer := grpc.NewServer()

	//create an instance of the currency server
	currencyServer := sercur.NewCurrency(loggs)
	// create HelloWorld Server
	helloServer := sercur.NewHelloServer(loggs)
	//Register the currency Server
	protos.RegisterCurrencyServer(grpcServer, currencyServer)
	protos.RegisterHelloWorldServer(grpcServer, helloServer)

	// reginter the reflection service which allows client to determine
	// methods for grpc service
	reflection.Register(grpcServer)

	// create a TCP socket for an Inbound server connection
	tcpSocket, err := net.Listen("tcp", serverPort)
	if err != nil {
		loggs.Error("Unable to create listener", "error", err)
		os.Exit(1)
	}

	grpcServer.Serve(tcpSocket)
}
