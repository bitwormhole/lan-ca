package vo

import "github.com/bitwormhole/lan-ca/backend/app/web/dto"

// Example ... VO
type Example struct {
	Base

	Items []*dto.Example `json:"items"`
}
