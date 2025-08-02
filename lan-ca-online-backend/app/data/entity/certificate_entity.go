package entity

import (
	"time"

	"github.com/bitwormhole/lan-ca/backend/app/data/dxo"
)

// Certificate ...
type Certificate struct {
	ID dxo.CertificateID

	Base

	StartedAt  time.Time
	StoppedAt  time.Time
	DomainName dxo.DomainName
	DomainID   dxo.DomainID

	Label       string
	Description string
}
