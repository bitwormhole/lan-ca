package iscenes

import (
	"github.com/bitwormhole/lan-ca/backend/app/classes/scenes"
	"github.com/bitwormhole/lan-ca/backend/app/data/dxo"
	"github.com/bitwormhole/lan-ca/backend/app/data/entity"
	"github.com/starter-go/base/lang"
	"github.com/starter-go/security/random"
	"gorm.io/gorm"
)

type SceneDaoImpl struct {

	//starter:component
	_as func(scenes.DAO) //starter:as('#')

	Agent   dxo.DatabaseAgent  //starter:inject("#")
	UUIDGen random.UUIDService //starter:inject("#")
}

func (inst *SceneDaoImpl) _impl() scenes.DAO {
	return inst
}

func (inst *SceneDaoImpl) modelItem() *entity.Scene {
	return new(entity.Scene)
}

func (inst *SceneDaoImpl) modelList() []*entity.Scene {
	return make([]*entity.Scene, 0)
}

func (inst *SceneDaoImpl) makeUUID() lang.UUID {
	builder := inst.UUIDGen.Build().Class("SceneDaoImpl")
	return builder.Generate()
}

func (inst *SceneDaoImpl) makeResult(item *entity.Scene, res *gorm.DB) (*entity.Scene, error) {
	err := res.Error
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (inst *SceneDaoImpl) Find(db *gorm.DB, id dxo.SceneID) (*entity.Scene, error) {
	db = inst.Agent.DB(db)
	o1 := inst.modelItem()
	res := db.First(o1, id)
	return inst.makeResult(o1, res)

}

func (inst *SceneDaoImpl) List(db *gorm.DB, q *scenes.Query) ([]*entity.Scene, error) {

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

func (inst *SceneDaoImpl) Insert(db *gorm.DB, item *entity.Scene) (*entity.Scene, error) {
	item.ID = 0
	item.UUID = inst.makeUUID()
	db = inst.Agent.DB(db)
	res := db.Create(item)
	err := res.Error
	if err != nil {
		return nil, err
	}
	return item, nil
}
