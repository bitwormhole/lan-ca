package entity

import (
	"time"

	"github.com/bitwormhole/lan-ca/backend/app/data/dxo"
)

// CertificateBase ...
type CertificateBase struct {

	// ID dxo.CertificateID

	Base

	SN string

	NotAfter   time.Time
	NotBefore  time.Time
	DomainName dxo.DomainName
	DomainID   dxo.DomainID

	Subject dxo.CertificateUserInfoJSON
	Signer  dxo.CertificateUserInfoJSON

	PublicKeyFingerprint   dxo.Fingerprint // in SHA256SUM
	CertificateFingerprint dxo.Fingerprint `gorm:"unique"` // in SHA256SUM

	Label   string
	Comment string
	State   dxo.CertificateState
}

// Certificate ...
type Certificate struct {
	ID dxo.CertificateID

	CertificateBase
}

// CertificateTemplate ...
type CertificateTemplate struct {
	ID dxo.CertificateTemplateID

	CertificateBase
}
