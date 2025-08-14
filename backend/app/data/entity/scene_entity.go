package entity

import "github.com/bitwormhole/lan-ca/backend/app/data/dxo"

// Scene ...
type Scene struct {
	ID dxo.SceneID

	Base

	Name  string
	QName string `gorm:"unique"`

	Label       string
	Comment     string
	Description string
}
