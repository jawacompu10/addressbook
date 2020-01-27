// This file contains functions/methods that do only the business logic.
// The code in this file is not aware of the transport/endpoint layers.

package service

import "bitbucket.org/jawacompu10/addressbook/transport"

// AddressbookService contains methods that act as the RPC for address book service
type AddressbookService struct {
	repo Repository
}

// NewService initializes and returns a new address book service value
func NewService(db Repository) *AddressbookService {
	return &AddressbookService{db}
}

// Repository is the interface that a DB wrapper for addressbook should implement
type Repository interface {
	GetAddressByID(string) (transport.Address, error)
	GetUserAddresses(string) ([]transport.Address, error)
	AddAddress(transport.Address) (transport.Address, error)
	UpdateAddress(transport.Address) (transport.Address, error)
}
