package entity

import "github.com/bitwormhole/lan-ca/backend/app/data/dxo"

// Domain ...
type Domain struct {
	ID dxo.DomainID

	Base

	Name        dxo.DomainName `gorm:"unique"`
	AddressIPv4 dxo.IPv4Address
	AddressIPv6 dxo.IPv6Address
	Comment     string

	Scene dxo.SceneID
}
