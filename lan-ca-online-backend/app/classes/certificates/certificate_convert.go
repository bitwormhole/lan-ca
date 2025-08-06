package certificates

import (
	"github.com/bitwormhole/lan-ca/backend/app/data/dxo"
	"github.com/bitwormhole/lan-ca/backend/app/data/entity"
	"github.com/bitwormhole/lan-ca/backend/app/web/dto"
	"github.com/starter-go/base/lang"
	"github.com/starter-go/security-gorm/rbacdb"
	"github.com/starter-go/vlog"
)

func ConvertD2E(src *dto.Certificate, dst *entity.Certificate) error {

	rbacdb.CopyBaseFieldsFromDtoToEntity(&src.BaseDTO, &dst.BaseEntity)
	dst.ID = src.ID

	var dn dxo.DomainName = ""
	for i, name := range src.DomainNames {
		if i == 0 {
			dn = name
		} else {
			vlog.Warn("这里只支持单个域名, 该域名将被忽略: ", name)
		}
	}

	dst.SN = src.SN
	dst.CertificateFingerprint = src.CertificateFingerprint
	dst.PublicKeyFingerprint = src.PublicKeyFingerprint
	dst.NotAfter = src.NotAfter.Time()
	dst.NotBefore = src.NotBefore.Time()
	dst.Subject = src.Subject.JSON()
	dst.Signer = src.Signer.JSON()
	dst.DomainName = dn
	dst.Label = src.Label
	dst.Comment = src.Comment

	return nil
}

func ConvertE2D(src *entity.Certificate, dst *dto.Certificate) error {

	rbacdb.CopyBaseFieldsFromEntityToDTO(&src.BaseEntity, &dst.BaseDTO)
	dst.ID = src.ID

	subject := src.Subject.Info()
	signer := src.Signer.Info()
	dn := src.DomainName

	dst.SN = src.SN
	dst.CertificateFingerprint = src.CertificateFingerprint
	dst.PublicKeyFingerprint = src.PublicKeyFingerprint
	dst.NotAfter = lang.NewTime(src.NotAfter)
	dst.NotBefore = lang.NewTime(src.NotBefore)
	dst.Subject = *subject
	dst.Signer = *signer
	dst.DomainNames = []dxo.DomainName{dn}
	dst.Label = src.Label
	dst.Comment = src.Comment

	return nil
}

func ConvertListE2D(src []*entity.Certificate, dst []*dto.Certificate) ([]*dto.Certificate, error) {

	if dst == nil {
		dst = make([]*dto.Certificate, 0)
	}

	for _, item1 := range src {
		item2 := new(dto.Certificate)
		err := ConvertE2D(item1, item2)
		if err != nil {
			return nil, err
		}
		dst = append(dst, item2)
	}
	return dst, nil
}
