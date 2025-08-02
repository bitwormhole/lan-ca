package certificates

import (
	"github.com/bitwormhole/lan-ca/backend/app/data/dxo"
	"github.com/bitwormhole/lan-ca/backend/app/data/entity"
	"gorm.io/gorm"
)

// CertificateDAO ...
type DAO interface {
	Find(db *gorm.DB, id dxo.CertificateID) (*entity.Certificate, error)
	List(db *gorm.DB, q *Query) ([]*entity.Certificate, error)
}
