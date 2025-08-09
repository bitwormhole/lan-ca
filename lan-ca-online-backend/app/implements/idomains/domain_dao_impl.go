package idomains

import (
	"github.com/bitwormhole/lan-ca/backend/app/classes/domains"
	"github.com/bitwormhole/lan-ca/backend/app/data/dxo"
	"github.com/bitwormhole/lan-ca/backend/app/data/entity"
	"github.com/starter-go/base/lang"
	"github.com/starter-go/security/random"
	"gorm.io/gorm"
)

type DomainDaoImpl struct {

	//starter:component
	_as func(domains.DAO) //starter:as('#')

	Agent   dxo.DatabaseAgent  //starter:inject("#")
	UUIDGen random.UUIDService //starter:inject("#")

}

func (inst *DomainDaoImpl) _impl() domains.DAO {
	return inst
}

func (inst *DomainDaoImpl) modelItem() *entity.Domain {
	return new(entity.Domain)
}

func (inst *DomainDaoImpl) modelList() []*entity.Domain {
	return make([]*entity.Domain, 0)
}

func (inst *DomainDaoImpl) makeItemUUID() lang.UUID {
	builder := inst.UUIDGen.Build().Class("DomainDaoImpl")
	return builder.Generate()
}

func (inst *DomainDaoImpl) makeResult(item *entity.Domain, res *gorm.DB) (*entity.Domain, error) {
	err := res.Error
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (inst *DomainDaoImpl) Find(db *gorm.DB, id dxo.DomainID) (*entity.Domain, error) {
	db = inst.Agent.DB(db)
	o1 := inst.modelItem()
	res := db.First(o1, id)
	return inst.makeResult(o1, res)
}

func (inst *DomainDaoImpl) List(db *gorm.DB, q *domains.Query) ([]*entity.Domain, error) {

	model := inst.modelItem()
	list1 := inst.modelList()
	var count int64

	db = inst.Agent.DB(db)
	db = db.Model(model)
	db = q.ApplyWhere(db)
	db = db.Count(&count)
	db = q.ApplyPagination(db)

	res := db.Find(&list1)
	err := res.Error
	if err != nil {
		return nil, err
	}

	q.Pagination.Total = count
	return list1, nil
}

func (inst *DomainDaoImpl) Insert(db *gorm.DB, item *entity.Domain) (*entity.Domain, error) {
	item.ID = 0
	item.UUID = inst.makeItemUUID()
	db = inst.Agent.DB(db)
	res := db.Create(item)
	return inst.makeResult(item, res)
}
