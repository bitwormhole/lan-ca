package icertificates

import (
	"context"
	"fmt"

	"github.com/bitwormhole/lan-ca/backend/app/classes/certificates"
	"github.com/bitwormhole/lan-ca/backend/app/data/dxo"
	"github.com/bitwormhole/lan-ca/backend/app/data/entity"
	"github.com/bitwormhole/lan-ca/backend/app/web/dto"
	"github.com/starter-go/base/lang"
	"github.com/starter-go/security/random"
)

type CertificateServiceImpl struct {

	//starter:component
	_as func(certificates.Service) //starter:as('#')

	Dao           certificates.DAO //starter:inject('#')
	RandomService random.Service   //starter:inject('#')

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

func (inst *CertificateServiceImpl) makeRandomNewCertFingerprint() dxo.Fingerprint {
	bin := inst.RandomService.NextBytes(128 / 8)
	hex := lang.HexFromBytes(bin)
	return dxo.Fingerprint(hex)
}

// computeMaxAgeAB 重新计算证书的有效期
func (inst *CertificateServiceImpl) computeMaxAgeAB(src, dst *dto.Certificate) {
	t1 := src.NotBefore
	t2 := src.NotAfter
	now := lang.Now()
	age := t2 - t1
	dst.NotBefore = now
	dst.NotAfter = now + age
}

func (inst *CertificateServiceImpl) Insert(ctx context.Context, it1 *dto.Certificate) (*dto.Certificate, error) {

	inst.computeMaxAgeAB(it1, it1)

	it2 := new(entity.Certificate)
	err := certificates.ConvertD2E(it1, it2)
	if err != nil {
		return nil, err
	}

	fp0 := inst.makeRandomNewCertFingerprint()
	it2.CertificateFingerprint = fp0 // 先用一个随机数填充,避免 index 冲突
	it2.State = dxo.CertificateStateInit

	it3, err := inst.Dao.Insert(nil, it2)
	if err != nil {
		return nil, err
	}

	it4 := new(dto.Certificate)
	err = certificates.ConvertE2D(it3, it4)
	if err != nil {
		return nil, err
	}

	return it4, nil
}
