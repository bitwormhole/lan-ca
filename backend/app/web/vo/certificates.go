package vo

import "github.com/bitwormhole/lan-ca/backend/app/web/dto"

// Certificates VO
type Certificates struct {
	Base

	Items []*dto.Certificate `json:"certificates"`
}
