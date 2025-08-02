package iscenes

import (
	"github.com/bitwormhole/lan-ca/backend/app/classes/scenes"
	"github.com/bitwormhole/lan-ca/backend/app/data/dxo"
	"github.com/bitwormhole/lan-ca/backend/app/data/entity"
	"gorm.io/gorm"
)

type SceneDaoImpl struct {

	//starter:component
	_as func(scenes.DAO) //starter:as('#')

	Agent dxo.DatabaseAgent //starter:inject("#")

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
