package dto

import "github.com/bitwormhole/lan-ca/backend/app/data/dxo"

// Certificate ...
type Certificate struct {
	ID dxo.CertificateID `json:"id"`

	Base

	Foo string `json:"foo"`
	Bar int    `json:"bar"`
}
