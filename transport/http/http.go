package http

import (
	"encoding/json"
	"log"
	"net/http"

	"bitbucket.org/jawacompu10/addressbook/transport"
	"github.com/gorilla/mux"
)

// StatusCoder is the interface that the custom errors should implement so that
// the handlers know what status code to return to the requests in case of the error
type StatusCoder interface {
	GetStatusCode() int
}

// Transport implements the Transport layer for http
type Transport struct {
	service transport.AddressService
}

// New creates a new value for the http Transport type
func New(service transport.AddressService) transport.Transport {
	return &Transport{
		service: service,
	}
}

// Start starts listen and serve via HTTP
func (ht *Transport) Start() error {
	r := ht.buildRouter()
	http.Handle("/", r)

	return http.ListenAndServe(":8080", nil)
}

func (ht *Transport) buildRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/address/{id}", ht.GetAddress).Methods("GET")
	r.HandleFunc("/address/user/{id}", ht.GetUserAddresses).Methods("GET")
	r.HandleFunc("/address", ht.CreateAddress).Methods("POST")
	r.HandleFunc("/address", ht.UpdateAddress).Methods("PUT")

	r.Use(addJSONTypeHeader)

	return r
}

func addJSONTypeHeader(nextHandler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		nextHandler.ServeHTTP(w, req)
	})
}

// GetAddress is the RPC for the get address API endpoint
func (ht *Transport) GetAddress(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	addrID := vars["id"]

	addr, err := ht.service.GetAddressByID(addrID)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), GetStatusCode(err))
		return
	}
	json.NewEncoder(w).Encode(addr)
}

// GetUserAddresses is the RPC for get user addresses endpoint
func (ht *Transport) GetUserAddresses(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	userID := vars["id"]

	addresses, err := ht.service.GetUserAddresses(userID)
	if err != nil {
		http.Error(w, "failed to get records for given user ID", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(addresses)
}

// CreateAddress is the RPC for the create address API endpoint
func (ht *Transport) CreateAddress(w http.ResponseWriter, req *http.Request) {
	log.Println("Request to CreateAddress")
	addr, err := decodeCreateAddressRequest(req.Body)
	if err != nil {
		http.Error(w, "Invalid payload", http.StatusBadRequest)
		return
	}
	addr, err = ht.service.CreateAddress(addr)
	if err != nil {
		http.Error(w, "DB write failed"+err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(addr)
}

// UpdateAddress is the RPC for the create address API endpoint
func (ht *Transport) UpdateAddress(w http.ResponseWriter, req *http.Request) {
	addr, err := decodeCreateAddressRequest(req.Body)
	if err != nil {
		http.Error(w, "Invalid payload", http.StatusBadRequest)
		return
	}
	addr, err = ht.service.UpdateAddress(addr)
	if err != nil {
		http.Error(w, err.Error(), GetStatusCode(err))
		return
	}
	json.NewEncoder(w).Encode(addr)
}

// GetStatusCode returns the HTTPS status code that should be returned
// in the event of the given error
func GetStatusCode(err error) int {
	code := http.StatusInternalServerError
	if sc, ok := getErrAsStatusCoder(err); ok {
		code = sc.GetStatusCode()
	}
	return code
}

func getErrAsStatusCoder(err error) (StatusCoder, bool) {
	var i interface{} = err
	v, ok := i.(StatusCoder)
	return v, ok
}
