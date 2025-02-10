package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"product-api/handlers"
	"time"

	"github.com/hashicorp/go-hclog"
	"github.com/nicholasjackson/env"
)

var bindAddress = env.String("BIND_ADDRESS", false, ":9090", "Bind address for the server")

func main() {

	env.Parse()

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

	// create the handler
	hello := handlers.NewHello(&loggs)

	serveMux := http.NewServeMux()
	serveMux.Handle("/", hello)

	//cerate a new server
	opts := hclog.StandardLoggerOptions{
		InferLevels: true,
	}
	httpServer := http.Server{
		Addr:         *bindAddress,
		Handler:      serveMux,
		ErrorLog:     loggs.StandardLogger(&opts),
		ReadTimeout:  5 * time.Second,   // max time to read request from the client
		WriteTimeout: 10 * time.Second,  // max time to write response to the client
		IdleTimeout:  120 * time.Second, // max time for connections using TCP Keep-Alive
	}

	// start the server in go routine

	go func() {
		loggs.Info("starting server on port 8090")
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

	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	httpServer.Shutdown(ctx)

	http.HandleFunc("/", hello.ServeHTTP) //

	http.ListenAndServe(":9092", nil)
}
