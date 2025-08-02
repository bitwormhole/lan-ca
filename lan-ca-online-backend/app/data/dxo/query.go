package dxo

import (
	"github.com/starter-go/rbac"
	"gorm.io/gorm"
)

// CommonQuery ... 是一个通用的基本查询对象
type CommonQuery struct {
	Pagination rbac.Pagination
	Limit      int64
	Offset     int64
}

func (inst *CommonQuery) ApplyPagination(db *gorm.DB) *gorm.DB {

	if inst == nil || db == nil {
		return db
	}

	page := &inst.Pagination
	limit := page.Limit()
	offset := page.Offset()

	inst.Limit = limit
	inst.Offset = offset

	if offset > 0 {
		db = db.Offset(int(offset))
	}
	if limit > 0 {
		db = db.Limit(int(limit))
	}

	return db
}
