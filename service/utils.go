package service

import (
	"encoding/json"
	"io"
	"time"
)

func decodeCreateAddressRequest(rdr io.Reader) (Address, error) {
	var addr Address
	err := json.NewDecoder(rdr).Decode(&addr)
	if err != nil {
		return addr, err
	}
	addr.CreatedAt = time.Now()
	fillDefaultValues(&addr)
	return addr, err
}

func decodeUpdateAddressRequest(rdr io.Reader) (Address, error) {
	var addr Address
	err := json.NewDecoder(rdr).Decode(&addr)
	if err != nil {
		return addr, err
	}
	fillDefaultValues(&addr)
	return addr, err
}

func fillDefaultValues(addr *Address) {
	addr.ModifiedAt = time.Now()
}
