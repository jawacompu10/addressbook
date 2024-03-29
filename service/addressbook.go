// This file contains functions/methods that do only the business logic.
// The code in this file is not aware of the transport/endpoint layers.

package service

import "addressbook/models"

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
	GetAddressByID(string) (models.Address, error)
	GetUserAddresses(string) ([]models.Address, error)
	AddAddress(models.Address) (models.Address, error)
	UpdateAddress(models.Address) (models.Address, error)
}
