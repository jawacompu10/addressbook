package models

import (
	"time"
)

// Address is the data structure that holds an address
type Address struct {
	ID         string    `json:"_id,omitempty"`
	UserID     string    `json:"user_id"`
	Name       string    `json:"name"`
	IsDefault  bool      `json:"is_default,omitempty"`
	Note       string    `json:"note,omitempty"`
	Lat        string    `json:"lat"`
	Long       string    `json:"long"`
	Details    Details   `json:"details"`
	CreatedAt  time.Time `json:"created_at,omitempty"`
	ModifiedAt time.Time `json:"modified_at,omitempty"`
}

// Details is a map to store the address details like street, city, etc
type Details map[string]string
