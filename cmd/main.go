// Package addressbook is the entry point of the address book service (of demand)
package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	repo "bitbucket.org/jawacompu10/addressbook/repo/mongo"
	"bitbucket.org/jawacompu10/addressbook/service"
	"bitbucket.org/jawacompu10/addressbook/transport"
	"bitbucket.org/jawacompu10/addressbook/transport/grpc"
	"bitbucket.org/jawacompu10/addressbook/transport/http"
)

var (
	httpTransport, grpcTransport transport.Transport
)

func main() {
	fmt.Println("main for address book service")
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	errorChan := make(chan error)

	db, err := repo.NewRepo(repo.DBInfo{
		URL:            "mongodb://localhost:27017/",
		DBName:         "demand",
		CollectionName: "addressbook",
	})
	if err != nil {
		log.Fatal("Cannot connect to DB")
	}

	go handleSignals(&errorChan)

	addressService := service.NewService(db)
	httpTransport = http.New(addressService)
	grpcTransport = grpc.New(addressService)
	go func() {
		errorChan <- httpTransport.Start()
	}()
	go func() {
		errorChan <- grpcTransport.Start()
	}()

	log.Println("Exit", <-errorChan)
}

func handleSignals(errorChan *chan error) {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	*errorChan <- fmt.Errorf("Signal received: %s", <-signalChan)
}
