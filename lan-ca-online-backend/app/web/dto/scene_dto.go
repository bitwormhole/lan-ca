package dto

import "github.com/bitwormhole/lan-ca/backend/app/data/dxo"

// Scene ...
type Scene struct {
	ID dxo.SceneID `json:"id"`

	Base

	Foo string `json:"foo"`
	Bar int    `json:"bar"`
}
