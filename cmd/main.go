// Package addressbook is the entry point of the address book service (of demand)
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gorilla/mux"

	repo "bitbucket.org/jawacompu10/addressbook/repo/mongo"
	"bitbucket.org/jawacompu10/addressbook/service"
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
	r := buildRouter(addressService)
	http.Handle("/", r)
	go func() {
		errorChan <- http.ListenAndServe(":8080", nil)
	}()

	log.Println("Exit", <-errorChan)
}

func buildRouter(addressService *service.AddressbookService) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/address/{id}", addressService.GetAddress).Methods("GET")
	r.HandleFunc("/address/user/{id}", addressService.GetUserAddresses).Methods("GET")
	r.HandleFunc("/address", addressService.CreateAddress).Methods("POST")
	r.HandleFunc("/address", addressService.UpdateAddress).Methods("PUT")

	r.Use(addJSONTypeHeader)

	return r
}

func addJSONTypeHeader(nextHandler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		nextHandler.ServeHTTP(w, req)
	})
}

func handleSignals(errorChan *chan error) {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	*errorChan <- fmt.Errorf("Signal received: %s", <-signalChan)
}
