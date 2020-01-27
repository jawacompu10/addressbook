package http

import (
	"encoding/json"
	"io"
	"time"

	"bitbucket.org/jawacompu10/addressbook/transport"
)

func decodeCreateAddressRequest(rdr io.Reader) (transport.Address, error) {
	var addr transport.Address
	err := json.NewDecoder(rdr).Decode(&addr)
	if err != nil {
		return addr, err
	}
	addr.CreatedAt = time.Now()
	fillDefaultValues(&addr)
	return addr, err
}

func decodeUpdateAddressRequest(rdr io.Reader) (transport.Address, error) {
	var addr transport.Address
	err := json.NewDecoder(rdr).Decode(&addr)
	if err != nil {
		return addr, err
	}
	fillDefaultValues(&addr)
	return addr, err
}

func fillDefaultValues(addr *transport.Address) {
	addr.ModifiedAt = time.Now()
}
