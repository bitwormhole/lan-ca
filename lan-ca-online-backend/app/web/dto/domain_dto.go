package dto

import "github.com/bitwormhole/lan-ca/backend/app/data/dxo"

// Domain ...
type Domain struct {
	ID dxo.DomainID `json:"id"`

	Base

	Name        dxo.DomainName  `json:"name"`
	IPv4        dxo.IPv4Address `json:"ipv4addr"`
	IPv6        dxo.IPv6Address `json:"ipv6addr"`
	Comment     string          `json:"comment"`
	Certificate string          `json:"cert_info"` // 表示证书状态的文本
}
