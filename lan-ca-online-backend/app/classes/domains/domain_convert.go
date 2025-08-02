package domains

import (
	"github.com/bitwormhole/lan-ca/backend/app/data/entity"
	"github.com/bitwormhole/lan-ca/backend/app/web/dto"
	"github.com/starter-go/security-gorm/rbacdb"
)

func ConvertD2E(src *dto.Domain, dst *entity.Domain) error {

	rbacdb.CopyBaseFieldsFromDtoToEntity(&src.BaseDTO, &dst.BaseEntity)
	dst.ID = src.ID

	dst.Name = src.Name
	dst.AddressIPv4 = src.IPv4
	dst.AddressIPv6 = src.IPv6
	dst.Comment = src.Comment

	return nil
}

func ConvertE2D(src *entity.Domain, dst *dto.Domain) error {

	rbacdb.CopyBaseFieldsFromEntityToDTO(&src.BaseEntity, &dst.BaseDTO)
	dst.ID = src.ID

	dst.Name = src.Name
	dst.IPv4 = src.AddressIPv4
	dst.IPv6 = src.AddressIPv6
	dst.Comment = src.Comment

	return nil
}

func ConvertListE2D(src []*entity.Domain, dst []*dto.Domain) ([]*dto.Domain, error) {
	for _, item1 := range src {
		item2 := new(dto.Domain)
		err := ConvertE2D(item1, item2)
		if err != nil {
			return nil, err
		}
		dst = append(dst, item2)
	}
	return dst, nil
}
