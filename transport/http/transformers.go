package http

import (
	"encoding/json"
	"io"
	"time"

	"bitbucket.org/jawacompu10/addressbook/models"
)

func decodeCreateAddressRequest(rdr io.Reader) (models.Address, error) {
	var addr models.Address
	err := json.NewDecoder(rdr).Decode(&addr)
	if err != nil {
		return addr, err
	}
	addr.CreatedAt = time.Now()
	fillDefaultValues(&addr)
	return addr, err
}

func decodeUpdateAddressRequest(rdr io.Reader) (models.Address, error) {
	var addr models.Address
	err := json.NewDecoder(rdr).Decode(&addr)
	if err != nil {
		return addr, err
	}
	fillDefaultValues(&addr)
	return addr, err
}

func fillDefaultValues(addr *models.Address) {
	addr.ModifiedAt = time.Now()
}
