package domains

import (
	"github.com/bitwormhole/lan-ca/backend/app/data/dxo"
	"github.com/bitwormhole/lan-ca/backend/app/data/entity"
	"gorm.io/gorm"
)

// DomainDAO ...
type DAO interface {
	Find(db *gorm.DB, id dxo.DomainID) (*entity.Domain, error)
	List(db *gorm.DB, q *Query) ([]*entity.Domain, error)
}
