package entity

import (
	"time"

	"github.com/bitwormhole/lan-ca/backend/app/data/dxo"
	"github.com/starter-go/trading"
)

// Contract ... 表示一份订阅合同
type Contract struct {
	ID dxo.ContractID

	Base

	NotBefore time.Time // 生效日期
	NotAfter  time.Time // 失效日期

	Title   string
	Content string

	Discount    float32               // 折扣 (0.5 表示 5 折)
	Quantity    float32               // 数量
	UnitPrice   trading.UnitPriceText // 单价
	Total       trading.PriceText     // 总价
	PayingPrice trading.PriceText     // 需支付的总价

}
