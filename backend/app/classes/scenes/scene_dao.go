package scenes

import (
	"github.com/bitwormhole/lan-ca/backend/app/data/dxo"
	"github.com/bitwormhole/lan-ca/backend/app/data/entity"
	"gorm.io/gorm"
)

// SceneDAO ...
type DAO interface {
	Find(db *gorm.DB, id dxo.SceneID) (*entity.Scene, error)
	List(db *gorm.DB, q *Query) ([]*entity.Scene, error)
	Insert(db *gorm.DB, item *entity.Scene) (*entity.Scene, error)
}
