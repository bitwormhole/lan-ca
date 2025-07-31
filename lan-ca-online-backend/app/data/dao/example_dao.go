package dao

import (
	"github.com/bitwormhole/lan-ca/backend/app/data/dxo"
	"github.com/bitwormhole/lan-ca/backend/app/data/entity"
	"gorm.io/gorm"
)

// ExampleDAO ...
type ExampleDAO interface {
	Find(db *gorm.DB, id dxo.ExampleID) (*entity.Example, error)
}
