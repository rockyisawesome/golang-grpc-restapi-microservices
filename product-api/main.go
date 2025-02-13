package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"product-api/database"
	"product-api/handlers"
	"time"

	configData "product-api/configurations"

	"github.com/hashicorp/go-hclog"
)

func main() {

	//create or open the log file
	logFile, err := os.OpenFile("app.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Some error occured in creating or opeing a log file", err)
		os.Exit(1)
	}
	loggs := hclog.New(&hclog.LoggerOptions{
		Name:       "myapp",
		Output:     logFile,    // Write logs to the log file
		Level:      hclog.Info, // Log level: Info and higher will be logged
		JSONFormat: false,      // Set true if you prefer JSON format for logs
	})

	// mongo db server connection
	cfg, err := configData.NewMongoDbConfig()
	if err != nil {
		loggs.Error("Error occured in creating new MongoDB instance", err)
		os.Exit(1)
	}

	// getting mongodb struct
	mongodb := database.NewMongoDB(cfg, &loggs)

	// connection with mongodb
	ctx := context.Background()
	err = mongodb.Connect(ctx)
	if err != nil {
		loggs.Error("Not able to connect to MongoDB database", err)
		os.Exit(1)
	}
	loggs.Error("Connection to database established")

	appcfg, err := configData.NewAppConfig()
	if err != nil {
		loggs.Error("Error occured in environment variable", err)
		os.Exit(1)
	}

	loggs.Info("port received from configuration", appcfg.AppURI)

	// create the handler
	hello := handlers.NewHello(&loggs, mongodb, ctx)

	serveMux := http.NewServeMux()
	serveMux.Handle("/", hello)
	

	//cerate a new server
	opts := hclog.StandardLoggerOptions{
		InferLevels: true,
	}
	httpServer := http.Server{
		Addr:         appcfg.AppURI,
		Handler:      serveMux,
		ErrorLog:     loggs.StandardLogger(&opts),
		ReadTimeout:  5 * time.Second,   // max time to read request from the client
		WriteTimeout: 10 * time.Second,  // max time to write response to the client
		IdleTimeout:  120 * time.Second, // max time for connections using TCP Keep-Alive
	}

	// start the server in go routine

	go func() {
		loggs.Info("starting server on port " + appcfg.AppURI)
		err := httpServer.ListenAndServe()
		if err != nil {
			loggs.Info("Error starting server", err)
			os.Exit(1)
		}
	}()

	//gracefully shutting down the server
	signalChannel := make(chan os.Signal)
	signal.Notify(signalChannel, os.Interrupt)
	signal.Notify(signalChannel, os.Kill)

	//blocking statement
	waitingForChanel := <-signalChannel

	loggs.Info("Got killing or Interrupt signal", waitingForChanel)

	//gracefully shutting down the server

	newctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	httpServer.Shutdown(newctx)

}
