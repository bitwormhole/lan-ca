package iscenes

import (
	"context"

	"github.com/bitwormhole/lan-ca/backend/app/classes/scenes"
	"github.com/bitwormhole/lan-ca/backend/app/data/dxo"
	"github.com/bitwormhole/lan-ca/backend/app/data/entity"
	"github.com/bitwormhole/lan-ca/backend/app/web/dto"
)

type SceneServiceImpl struct {

	//starter:component
	_as func(scenes.Service) //starter:as('#')

	Dao scenes.DAO //starter:inject('#')

}

func (inst *SceneServiceImpl) _impl() scenes.Service {
	return inst
}

func (inst *SceneServiceImpl) Find(ctx context.Context, id dxo.SceneID) (*dto.Scene, error) {
	o1, err := inst.Dao.Find(nil, id)
	if err != nil {
		return nil, err
	}
	o2 := new(dto.Scene)
	err = scenes.ConvertE2D(o1, o2)
	return o2, err
}

func (inst *SceneServiceImpl) List(ctx context.Context, q *scenes.Query) ([]*dto.Scene, error) {

	want := q.Want
	if want == nil {
		want = new(entity.Scene)
		want.Owner = 17
		q.Want = want
	}

	list1, err := inst.Dao.List(nil, q)
	if err != nil {
		return nil, err
	}
	return scenes.ConvertListE2D(list1, nil)
}
