package domains

import (
	"github.com/bitwormhole/lan-ca/backend/app/data/dxo"
	"github.com/bitwormhole/lan-ca/backend/app/data/entity"
	"gorm.io/gorm"
)

type Query struct {
	dxo.CommonQuery

	Want *entity.Domain // 可选的查询条件
}

func (inst *Query) ApplyWhere(db *gorm.DB) *gorm.DB {
	want := inst.Want
	if want == nil {
		return db
	}
	return db.Where(want)
}
