package service

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// GetAddress is the RPC for the get address API endpoint
func (as *AddressbookService) GetAddress(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	addrID := vars["id"]

	addr, err := as.repo.GetAddressByID(addrID)
	if err != nil {
		log.Println(err)
		http.Error(w, "failed to get record with given ID", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(addr)
}

// GetUserAddresses is the RPC for get user addresses endpoint
func (as *AddressbookService) GetUserAddresses(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	userID := vars["id"]

	addresses, err := as.repo.GetUserAddresses(userID)
	if err != nil {
		http.Error(w, "failed to get records for given user ID", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(addresses)
}

// CreateAddress is the RPC for the create address API endpoint
func (as *AddressbookService) CreateAddress(w http.ResponseWriter, req *http.Request) {
	addr, err := decodeCreateAddressRequest(req.Body)
	if err != nil {
		http.Error(w, "Invalid payload", http.StatusBadRequest)
		return
	}
	addr, err = as.repo.AddAddress(addr)
	if err != nil {
		http.Error(w, "DB write failed"+err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(addr)
}

// UpdateAddress is the RPC for the create address API endpoint
func (as *AddressbookService) UpdateAddress(w http.ResponseWriter, req *http.Request) {
	addr, err := decodeCreateAddressRequest(req.Body)
	if err != nil {
		http.Error(w, "Invalid payload", http.StatusBadRequest)
		return
	}
	addr, err = as.repo.UpdateAddress(addr)
	if err != nil {
		http.Error(w, "DB write failed"+err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(addr)
}
