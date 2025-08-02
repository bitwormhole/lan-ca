package icertificates

import (
	"github.com/bitwormhole/lan-ca/backend/app/classes/certificates"
	"github.com/bitwormhole/lan-ca/backend/app/data/dxo"
	"github.com/bitwormhole/lan-ca/backend/app/data/entity"
	"gorm.io/gorm"
)

type CertificateDaoImpl struct {

	//starter:component
	_as func(certificates.DAO) //starter:as('#')

	Agent dxo.DatabaseAgent //starter:inject("#")

}

func (inst *CertificateDaoImpl) _impl() certificates.DAO {
	return inst
}

func (inst *CertificateDaoImpl) modelItem() *entity.Certificate {
	return new(entity.Certificate)
}

func (inst *CertificateDaoImpl) modelList() []*entity.Certificate {
	return make([]*entity.Certificate, 0)
}

func (inst *CertificateDaoImpl) makeResult(item *entity.Certificate, res *gorm.DB) (*entity.Certificate, error) {
	err := res.Error
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (inst *CertificateDaoImpl) Find(db *gorm.DB, id dxo.CertificateID) (*entity.Certificate, error) {
	db = inst.Agent.DB(db)
	o1 := inst.modelItem()
	res := db.First(o1, id)
	return inst.makeResult(o1, res)
}

func (inst *CertificateDaoImpl) List(db *gorm.DB, q *certificates.Query) ([]*entity.Certificate, error) {

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
