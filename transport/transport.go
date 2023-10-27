package transport

import "addressbook/models"

// Transport is the interface that defines the transport layer
type Transport interface {
	Start() error
}

// AddressService defines the address book service
type AddressService interface {
	GetAddressByID(string) (models.Address, error)
	GetUserAddresses(string) ([]models.Address, error)
	CreateAddress(models.Address) (models.Address, error)
	UpdateAddress(models.Address) (models.Address, error)
}
