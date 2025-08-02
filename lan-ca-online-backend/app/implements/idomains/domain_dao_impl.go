package idomains

import (
	"github.com/bitwormhole/lan-ca/backend/app/classes/domains"
	"github.com/bitwormhole/lan-ca/backend/app/data/dxo"
	"github.com/bitwormhole/lan-ca/backend/app/data/entity"
	"gorm.io/gorm"
)

type DomainDaoImpl struct {

	//starter:component
	_as func(domains.DAO) //starter:as('#')

	Agent dxo.DatabaseAgent //starter:inject("#")

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

	list1 := inst.modelList()

	db = inst.Agent.DB(db)
	db = q.ApplyPagination(db)
	db = q.ApplyWhere(db)

	res := db.Find(&list1)
	err := res.Error
	if err != nil {
		return nil, err
	}
	return list1, nil
}
