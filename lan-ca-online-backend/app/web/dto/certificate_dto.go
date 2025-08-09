package dto

import (
	"github.com/bitwormhole/lan-ca/backend/app/data/dxo"
	"github.com/starter-go/base/lang"
)

// Certificate ...
type Certificate struct {
	ID dxo.CertificateID `json:"id"`

	Base

	SN          string           `json:"sn"`
	DomainNames []dxo.DomainName `json:"domain_names"`

	PublicKeyFingerprint   dxo.Fingerprint `json:"public_key_fingerprint"`  // in SHA256SUM
	CertificateFingerprint dxo.Fingerprint `json:"certificate_fingerprint"` // in SHA256SUM

	NotBefore lang.Time `json:"not_before"`
	NotAfter  lang.Time `json:"not_after"`

	Subject dxo.CertificateUserInfo `json:"subject"`
	Signer  dxo.CertificateUserInfo `json:"signer"`

	Label   string               `json:"label"`
	Comment string               `json:"comment"`
	State   dxo.CertificateState `json:"state"`
}
