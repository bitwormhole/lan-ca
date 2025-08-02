package permissions

import (
	"context"

	"github.com/starter-go/rbac"
	"github.com/starter-go/security-gorm/rbacdb"
)

type Checking interface {
	Context() context.Context   // 获取与这个 checking 相对应的 context
	HasError() bool             // 判断是否发生了错误
	Error() error               // 获取之前发生的错误
	CurrentUserID() rbac.UserID // 取当前的用户ID

	Reset() Checking                          // 复位这个 checking 对象
	SetAuthIsRequired(required bool) Checking // 设置是否必须登录

	CheckDTO(o *rbac.BaseDTO) Checking         // 检查 DTO
	CheckEntity(o *rbacdb.BaseEntity) Checking // 检查 entity
	Check() error                              // 完成最终的检查
}

type Checker interface {
	NewChecking(ctx context.Context) Checking
}
