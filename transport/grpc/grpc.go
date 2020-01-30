package grpc

import (
	"context"
	"log"
	"net"

	"bitbucket.org/jawacompu10/addressbook/transport"
	"bitbucket.org/jawacompu10/addressbook/transport/grpc/addressbook"

	"google.golang.org/grpc"
)

// Transport implements the Transport layer for gRPC
type Transport struct {
	service transport.AddressService
}

// New creates a new value for the gRPC Transport type
func New(service transport.AddressService) transport.Transport {
	return &Transport{
		service: service,
	}
}

// Start starts listening for incoming requests
func (gt *Transport) Start() error {
	lis, err := net.Listen("tcp", ":9999")
	if err != nil {
		log.Fatal("Failed to listen", err)
	}
	grpcServer := grpc.NewServer()
	addressbook.RegisterAddressBookServer(grpcServer, gt)
	return grpcServer.Serve(lis)
}

// GetAddress gets the Address with the given ID
func (gt *Transport) GetAddress(ctx context.Context, addrID *addressbook.AddressID) (*addressbook.Address, error) {
	addr, err := gt.service.GetAddressByID(addrID.GetID())
	if err != nil {
		return &addressbook.Address{}, nil
	}
	return transportAddressFromModel(addr)
}

// GetUserAddresses gets the Addresses of the given user
func (gt *Transport) GetUserAddresses(userID *addressbook.UserID, stream addressbook.AddressBook_GetUserAddressesServer) error {
	addresses, err := gt.service.GetUserAddresses(userID.GetID())
	if err != nil {
		return err
	}
	for _, addr := range addresses {
		var trAddr *addressbook.Address
		if trAddr, err = transportAddressFromModel(addr); err != nil {
			return err
		}
		if err := stream.Send(trAddr); err != nil {
			return err
		}
	}
	return nil
}

// AddAddress adds the given Address
func (gt *Transport) AddAddress(ctx context.Context, addr *addressbook.Address) (*addressbook.Address, error) {
	mAddr, err := modelFromTransportAddress(*addr)
	if err != nil {
		return nil, err
	}
	mAddr, err = gt.service.CreateAddress(mAddr)
	if err != nil {
		return nil, err
	}
	return transportAddressFromModel(mAddr)
}

// UpdateAddress updates an existing Address
func (gt *Transport) UpdateAddress(ctx context.Context, addr *addressbook.Address) (*addressbook.Address, error) {
	mAddr, err := modelFromTransportAddress(*addr)
	if err != nil {
		return nil, err
	}
	mAddr, err = gt.service.UpdateAddress(mAddr)
	if err != nil {
		return nil, err
	}
	return transportAddressFromModel(mAddr)
}
