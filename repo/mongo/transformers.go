package repo

import (
	"bitbucket.org/jawacompu10/addressbook/service"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (da *dbAddress) fromServiceAddress(addr service.Address) error {
	if addr.ID != "" {
		var err error
		da.ID, err = primitive.ObjectIDFromHex(addr.ID)
		if err != nil {
			return err
		}
	} else {
		da.ID = primitive.NewObjectID()
	}
	da.Address = addr
	return nil
}

func (da *dbAddress) toServiceAddress() service.Address {
	addr := da.Address
	addr.ID = da.ID.Hex()
	return addr
}
