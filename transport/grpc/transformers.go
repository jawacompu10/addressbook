package grpc

import (
	"fmt"

	"addressbook/models"
	"addressbook/transport/grpc/addressbook"

	"github.com/golang/protobuf/ptypes"
)

func modelFromTransportAddress(addr addressbook.Address) (models.Address, error) {
	outAddr := models.Address{
		ID:        addr.GetID(),
		UserID:    addr.GetUserID(),
		Name:      addr.GetName(),
		IsDefault: addr.GetIsDefault(),
		Note:      addr.GetNote(),
		Lat:       addr.GetLat(),
		Long:      addr.GetLong(),
		Details:   addr.GetDetails(),
	}
	var err error
	outAddr.CreatedAt, err = ptypes.Timestamp(addr.GetCreatedAt())
	if err != nil {
		return outAddr, fmt.Errorf("Could not parse CreatedAt date: %w", err)
	}
	outAddr.ModifiedAt, err = ptypes.Timestamp(addr.GetModifiedAt())
	if err != nil {
		return outAddr, fmt.Errorf("Could not parse ModifiedAt date: %w", err)
	}

	return outAddr, err
}

func transportAddressFromModel(addr models.Address) (*addressbook.Address, error) {
	outAddr := addressbook.Address{
		ID:        addr.ID,
		UserID:    addr.UserID,
		Name:      addr.Name,
		IsDefault: addr.IsDefault,
		Note:      addr.Note,
		Lat:       addr.Lat,
		Long:      addr.Long,
		Details:   addr.Details,
	}
	var err error
	outAddr.CreatedAt, err = ptypes.TimestampProto(addr.CreatedAt)
	if err != nil {
		return &outAddr, fmt.Errorf("Could not parse CreatedAt date: %w", err)
	}
	outAddr.ModifiedAt, err = ptypes.TimestampProto(addr.ModifiedAt)
	if err != nil {
		return &outAddr, fmt.Errorf("Could not parse ModifiedAt date: %w", err)
	}

	return &outAddr, err
}
