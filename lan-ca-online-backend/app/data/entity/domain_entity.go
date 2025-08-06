package entity

import "github.com/bitwormhole/lan-ca/backend/app/data/dxo"

// Domain ...
type Domain struct {
	ID dxo.DomainID

	Base

	Name  dxo.DomainName  `gorm:"unique"`
	IPv4  dxo.IPv4Address `gorm:"column:ipv4_address"`
	IPv6  dxo.IPv6Address `gorm:"column:ipv6_address"`
	Scene dxo.SceneID

	Label       string
	Comment     string
	Description string
}
