package certificates

import (
	"github.com/bitwormhole/lan-ca/backend/app/data/entity"
	"github.com/bitwormhole/lan-ca/backend/app/web/dto"
	"github.com/starter-go/security-gorm/rbacdb"
)

func ConvertD2E(src *dto.Certificate, dst *entity.Certificate) error {

	rbacdb.CopyBaseFieldsFromDtoToEntity(&src.BaseDTO, &dst.BaseEntity)
	dst.ID = src.ID

	return nil
}

func ConvertE2D(src *entity.Certificate, dst *dto.Certificate) error {

	rbacdb.CopyBaseFieldsFromEntityToDTO(&src.BaseEntity, &dst.BaseDTO)
	dst.ID = src.ID

	return nil
}

func ConvertListE2D(src []*entity.Certificate, dst []*dto.Certificate) ([]*dto.Certificate, error) {
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
