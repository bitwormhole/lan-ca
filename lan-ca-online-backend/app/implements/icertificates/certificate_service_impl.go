package icertificates

import (
	"context"

	"github.com/bitwormhole/lan-ca/backend/app/classes/certificates"
	"github.com/bitwormhole/lan-ca/backend/app/classes/permissions"
	"github.com/bitwormhole/lan-ca/backend/app/data/dxo"
	"github.com/bitwormhole/lan-ca/backend/app/data/entity"
	"github.com/bitwormhole/lan-ca/backend/app/web/dto"
)

type CertificateServiceImpl struct {

	//starter:component
	_as func(certificates.Service) //starter:as('#')

	Dao         certificates.DAO    //starter:inject('#')
	PermChecker permissions.Checker //starter:inject('#')

}

func (inst *CertificateServiceImpl) _impl() certificates.Service {
	return inst
}

func (inst *CertificateServiceImpl) Find(ctx context.Context, id dxo.CertificateID) (*dto.Certificate, error) {
	o1, err := inst.Dao.Find(nil, id)
	if err != nil {
		return nil, err
	}
	o2 := new(dto.Certificate)
	err = certificates.ConvertE2D(o1, o2)
	return o2, err
}

func (inst *CertificateServiceImpl) List(ctx context.Context, q *certificates.Query) ([]*dto.Certificate, error) {

	checking := inst.PermChecker.NewChecking(ctx)

	want := q.Want
	if want == nil {
		want = new(entity.Certificate)
		want.Owner = checking.CurrentUserID()
		q.Want = want
	}

	err := checking.Check() // 查询前,检查
	if err != nil {
		return nil, err
	}

	list1, err := inst.Dao.List(nil, q)
	if err != nil {
		return nil, err
	}

	for _, item := range list1 {
		checking.CheckEntity(&item.BaseEntity)
	}

	err = checking.Check() // 查询后,检查
	if err != nil {
		return nil, err
	}

	return certificates.ConvertListE2D(list1, nil)
}
