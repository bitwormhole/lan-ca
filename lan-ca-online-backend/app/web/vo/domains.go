package vo

import "github.com/bitwormhole/lan-ca/backend/app/web/dto"

// Domains ... VO
type Domains struct {
	Base

	Items []*dto.Domain `json:"domains"`
}
