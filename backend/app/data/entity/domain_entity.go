package entity

import "github.com/bitwormhole/lan-ca/backend/app/data/dxo"

// Domain ...
type Domain struct {
	ID dxo.DomainID

	Base

	Name              dxo.DomainName `gorm:"unique"` // 系统唯一的域名
	PersonalizedDN    dxo.DomainName // 个性化域名
	UsePersonalizedDN bool           // 是否使用个性化域名
	Scene             dxo.SceneID    // 使用场景

	IPv4 dxo.IPv4Address `gorm:"column:ipv4_address"`
	IPv6 dxo.IPv6Address `gorm:"column:ipv6_address"`

	Label       string
	Comment     string
	Description string
}
