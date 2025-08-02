package entity

import "github.com/bitwormhole/lan-ca/backend/app/data/dxo"

// Scene ...
type Scene struct {
	ID dxo.SceneID

	Base

	Label       string
	Description string
}
