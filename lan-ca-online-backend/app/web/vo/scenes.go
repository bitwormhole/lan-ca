package vo

import "github.com/bitwormhole/lan-ca/backend/app/web/dto"

// Scenes VO
type Scenes struct {
	Base

	Items []*dto.Scene `json:"scenes"`
}
