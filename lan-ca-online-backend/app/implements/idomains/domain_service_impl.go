package idomains

import (
	"context"

	"github.com/bitwormhole/lan-ca/backend/app/classes/domains"
	"github.com/bitwormhole/lan-ca/backend/app/data/dxo"
	"github.com/bitwormhole/lan-ca/backend/app/data/entity"
	"github.com/bitwormhole/lan-ca/backend/app/web/dto"
	"github.com/starter-go/rbac"
)

type DomainServiceImpl struct {

	//starter:component
	_as func(domains.Service) //starter:as('#')

	Dao domains.DAO //starter:inject('#')

	Subjects rbac.SubjectService //starter:inject('#')

}

func (inst *DomainServiceImpl) _impl() domains.Service {
	return inst
}

func (inst *DomainServiceImpl) Find(ctx context.Context, id dxo.DomainID) (*dto.Domain, error) {

	o1, err := inst.Dao.Find(nil, id)
	if err != nil {
		return nil, err
	}

	o2 := &dto.Domain{}
	err = domains.ConvertE2D(o1, o2)
	if err != nil {
		return nil, err
	}

	return o2, nil
}

func (inst *DomainServiceImpl) List(ctx context.Context, q *domains.Query) ([]*dto.Domain, error) {

	subject, err := inst.Subjects.GetCurrent(ctx)
	if err != nil {
		return nil, err
	}

	want := q.Want
	if want == nil {
		want = new(entity.Domain)
		want.Owner = subject.User
		q.Want = want
	}

	////////////////////////////////////////

	list1, err := inst.Dao.List(nil, q)
	if err != nil {
		return nil, err
	}
	return domains.ConvertListE2D(list1, nil)
}
