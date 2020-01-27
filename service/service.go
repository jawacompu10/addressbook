package service

import "bitbucket.org/jawacompu10/addressbook/transport"

// GetAddressByID gets an address specified by the given ID
func (as *AddressbookService) GetAddressByID(id string) (transport.Address, error) {
	return as.repo.GetAddressByID(id)
}

// GetUserAddresses gets all the addresses of the specified user
func (as *AddressbookService) GetUserAddresses(id string) ([]transport.Address, error) {
	return as.repo.GetUserAddresses(id)
}

// CreateAddress adds a new address to the system
func (as *AddressbookService) CreateAddress(addr transport.Address) (transport.Address, error) {
	return as.repo.AddAddress(addr)
}

// UpdateAddress updates an existing address
func (as *AddressbookService) UpdateAddress(addr transport.Address) (transport.Address, error) {
	return as.repo.UpdateAddress(addr)
}
