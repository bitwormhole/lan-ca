package entity

import (
	"github.com/bitwormhole/lan-ca/backend/app/data/dxo"
	"github.com/starter-go/trading"
)

// Solution ... 表示一种订阅方案(aka.套餐)
type Solution struct {
	ID dxo.SolutionID

	Base

	Discount  float32               // 折扣 (0.5 表示 5 折)
	UnitPrice trading.UnitPriceText // 单价

}
