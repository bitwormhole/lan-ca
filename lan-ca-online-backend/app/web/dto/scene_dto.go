package dto

import "github.com/bitwormhole/lan-ca/backend/app/data/dxo"

// Scene ...
type Scene struct {
	ID dxo.SceneID `json:"id"`

	Base

	Name        string `json:"name"`
	Label       string `json:"label"`
	Comment     string `json:"comment"`
	Description string `json:"description"`
}
