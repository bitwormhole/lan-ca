package scenes

import (
	"github.com/bitwormhole/lan-ca/backend/app/data/entity"
	"github.com/bitwormhole/lan-ca/backend/app/web/dto"
	"github.com/starter-go/security-gorm/rbacdb"
)

func ConvertD2E(src *dto.Scene, dst *entity.Scene) error {

	uidstr := src.Owner.String()

	rbacdb.CopyBaseFieldsFromDtoToEntity(&src.BaseDTO, &dst.BaseEntity)
	dst.ID = src.ID

	dst.Name = src.Name
	dst.QName = uidstr + ":" + src.Name

	dst.Label = src.Label
	dst.Comment = src.Comment
	dst.Description = src.Description

	return nil
}

func ConvertE2D(src *entity.Scene, dst *dto.Scene) error {

	rbacdb.CopyBaseFieldsFromEntityToDTO(&src.BaseEntity, &dst.BaseDTO)
	dst.ID = src.ID

	dst.Name = src.Name
	dst.Label = src.Label
	dst.Comment = src.Comment
	dst.Description = src.Description

	return nil
}

func ConvertListE2D(src []*entity.Scene, dst []*dto.Scene) ([]*dto.Scene, error) {
	for _, item1 := range src {
		item2 := new(dto.Scene)
		err := ConvertE2D(item1, item2)
		if err != nil {
			return nil, err
		}
		dst = append(dst, item2)
	}
	return dst, nil
}
