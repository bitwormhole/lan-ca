package icertificates

import (
	"context"
	"fmt"

	"github.com/bitwormhole/lan-ca/backend/app/classes/certificates"
	"github.com/bitwormhole/lan-ca/backend/app/data/dxo"
	"github.com/bitwormhole/lan-ca/backend/app/data/entity"
	"github.com/bitwormhole/lan-ca/backend/app/web/dto"
)

type CertificateServiceImpl struct {

	//starter:component
	_as func(certificates.Service) //starter:as('#')

	Dao certificates.DAO //starter:inject('#')

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

	want := q.Want
	if want == nil {
		want = new(entity.Certificate)
		want.Owner = q.User
		q.Want = want
	}
	if want.Owner < 1 {
		return nil, fmt.Errorf("owner is 0")
	}

	list1, err := inst.Dao.List(nil, q)
	if err != nil {
		return nil, err
	}

	return certificates.ConvertListE2D(list1, nil)
}
