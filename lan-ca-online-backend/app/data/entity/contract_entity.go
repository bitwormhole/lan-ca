package entity

import (
	"time"

	"github.com/bitwormhole/lan-ca/backend/app/data/dxo"
)

// Contract ... 表示一份订阅合同
type Contract struct {
	ID dxo.ContractID

	Base

	NotBefore time.Time // 生效日期
	NotAfter  time.Time // 失效日期

}
