package entity

import "github.com/bitwormhole/lan-ca/backend/app/data/dxo"

// Solution ... 表示一种订阅方案(aka.套餐)
type Solution struct {
	ID dxo.SolutionID

	Base
}
