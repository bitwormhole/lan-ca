package permissions

import (
	"context"

	"github.com/starter-go/rbac"
	"github.com/starter-go/security-gorm/rbacdb"
)

// 定义一些角色
const (
	RoleUserDoNothing rbac.RoleName = "user_do_nothing" // 表示访问过程中没有执行有效操作的用户
)

type Checker interface {
	Reset() Checker // 复位这个 checking 对象
	Close() error   // 完成最终的检查

	Context() context.Context   // 获取与这个 checking 相对应的 context
	HasError() bool             // 判断是否发生了错误
	Error() error               // 获取之前发生的错误
	CurrentUserID() rbac.UserID // 取当前的用户ID

	CheckOwnerForDTO(o *rbac.BaseDTO) Checker         // 检查 DTO
	CheckOwnerForEntity(o *rbacdb.BaseEntity) Checker // 检查 entity

	AddRoles(roles ...rbac.RoleName) Checker // 为当前用户赋予指定的角色
	HandleError(err error) Checker           // 处理错误
}

type CheckerBuilder interface {
	SetAuthIsRequired(required bool) CheckerBuilder // 设置是否必须登录

	AcceptRoles(roles ...rbac.RoleName) CheckerBuilder

	Open() Checker
}

type CheckerService interface {
	NewBuilder(ctx context.Context) CheckerBuilder
}

////////////////////////////////////////////////////////////////////////////////

type CheckerHolder struct {
	checker          Checker
	builder          CheckerBuilder
	context          context.Context
	uid              rbac.UserID
	accept_roles     []rbac.RoleName
	accept_anonymous bool // 是否接受匿名访问
}

func (inst *CheckerHolder) Init(ctx context.Context, ser CheckerService) error {

	builder := ser.NewBuilder(ctx)
	builder.AcceptRoles(inst.accept_roles...)
	builder.SetAuthIsRequired(!inst.accept_anonymous)

	checker := builder.Open()

	inst.builder = builder
	inst.checker = checker
	inst.context = ctx
	inst.uid = checker.CurrentUserID()

	return checker.Error()
}

func (inst *CheckerHolder) AcceptAnonymous(can_anonymous bool) {
	inst.accept_anonymous = can_anonymous
}

func (inst *CheckerHolder) AcceptRoles(roles ...rbac.RoleName) {
	inst.accept_roles = append(inst.accept_roles, roles...)
}

func (inst *CheckerHolder) Done(err error) error {
	inst.checker.HandleError(err)
	return inst.checker.Close()
}

func (inst *CheckerHolder) GetChecker() Checker {
	return inst.checker
}

func (inst *CheckerHolder) GetCheckerBuilder() CheckerBuilder {
	return inst.builder
}
