package transport

// Transport is the interface that defines the transport layer
type Transport interface {
	Start() error
}

// AddressService defines the address book service
type AddressService interface {
	GetAddressByID(string) (Address, error)
	GetUserAddresses(string) ([]Address, error)
	CreateAddress(Address) (Address, error)
	UpdateAddress(Address) (Address, error)
}
