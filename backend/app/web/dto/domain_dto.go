package dto

import "github.com/bitwormhole/lan-ca/backend/app/data/dxo"

// Domain ...
type Domain struct {
	ID dxo.DomainID `json:"id"`

	Base

	Name              dxo.DomainName  `json:"name"`
	IPv4              dxo.IPv4Address `json:"ipv4_address"`
	IPv6              dxo.IPv6Address `json:"ipv6_address"`
	Scene             dxo.SceneID     `json:"scene"`               // 所属应用场景的 id
	PersonalizedDN    dxo.DomainName  `json:"personalized_dn"`     // 个性化域名
	UsePersonalizedDN bool            `json:"use_personalized_dn"` // 是否使用个性化域名

	Certificate string `json:"cert_info"` // 表示证书状态的文本

	Label       string `json:"label"`
	Comment     string `json:"comment"`
	Description string `json:"description"`
}
