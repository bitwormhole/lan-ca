package ipermissions

import (
	"context"
	"net/http"

	"github.com/bitwormhole/lan-ca/backend/app/classes/permissions"
	"github.com/starter-go/libgin/web"
	"github.com/starter-go/rbac"
	"github.com/starter-go/security-gorm/rbacdb"
)

////////////////////////////////////////////////////////////////////////////////
// CheckerServiceImpl

type CheckerServiceImpl struct {

	//starter:component

	_as func(permissions.CheckerService) //starter:as("#")

	Subjects rbac.SubjectService //starter:inject("#")

}

func (inst *CheckerServiceImpl) _impl() permissions.CheckerService {
	return inst
}

func (inst *CheckerServiceImpl) NewBuilder(ctx context.Context) permissions.CheckerBuilder {
	builder := &innerCheckerBuilder{
		context:        ctx,
		service:        inst,
		isAuthRequired: true,
	}
	return builder
}

////////////////////////////////////////////////////////////////////////////////
// innerCheckerBuilder

type innerCheckerBuilder struct {
	service          *CheckerServiceImpl
	context          context.Context
	isAuthRequired   bool
	acceptedRoleList []rbac.RoleName
}

func (inst *innerCheckerBuilder) _impl() permissions.CheckerBuilder {
	return inst
}

func (inst *innerCheckerBuilder) SetAuthIsRequired(required bool) permissions.CheckerBuilder {
	inst.isAuthRequired = required
	return inst
}

func (inst *innerCheckerBuilder) AcceptRoles(roles ...rbac.RoleName) permissions.CheckerBuilder {
	inst.acceptedRoleList = append(inst.acceptedRoleList, roles...)
	return inst
}

func (inst *innerCheckerBuilder) Open() permissions.Checker {
	checker := &innerChecker{
		service: inst.service,
		context: inst.context,
		builder: inst,
	}
	checker.init()
	return checker
}

////////////////////////////////////////////////////////////////////////////////
// innerChecker

type innerChecker struct {
	service *CheckerServiceImpl
	builder *innerCheckerBuilder
	context context.Context

	err          error
	subject      rbac.SubjectDTO
	haveRoleList []rbac.RoleName

	countEntityWithOwnership    int
	countEntityWithoutOwnership int
}

func (inst *innerChecker) _impl() permissions.Checker {
	return inst
}

// init == reset == reload

func (inst *innerChecker) Reset() permissions.Checker {
	err := inst.reload()
	inst.HandleError(err)
	return inst
}

func (inst *innerChecker) Context() context.Context {
	return inst.context
}

func (inst *innerChecker) HasError() bool {
	return (inst.err != nil)
}

func (inst *innerChecker) Error() error {
	return inst.err
}

func (inst *innerChecker) CurrentUserID() rbac.UserID {
	uid := inst.subject.CurrentUser.User
	auth := inst.subject.Authenticated
	if uid < 1 {
		return 0
	}
	if !auth {
		return 0
	}
	return uid
}

func (inst *innerChecker) CheckOwnerForDTO(o *rbac.BaseDTO) permissions.Checker {
	if o == nil {
		return inst
	}
	u1 := o.Owner
	u2 := inst.CurrentUserID()
	if u1 == u2 && u1 > 0 {
		inst.countEntityWithOwnership++
	} else {
		inst.countEntityWithoutOwnership++
	}
	return inst
}

func (inst *innerChecker) CheckOwnerForEntity(o *rbacdb.BaseEntity) permissions.Checker {
	if o == nil {
		return inst
	}
	u1 := o.Owner
	u2 := inst.CurrentUserID()
	if u1 == u2 && u1 > 0 {
		inst.countEntityWithOwnership++
	} else {
		inst.countEntityWithoutOwnership++
	}
	return inst
}

func (inst *innerChecker) Close() error {
	inst.innerCheckOwnership()
	inst.innerCheckRoles()
	return inst.err
}

func (inst *innerChecker) HandleError(err error) permissions.Checker {
	if err == nil {
		return inst // ignore nil
	}
	older := inst.err
	if older == nil {
		inst.err = err // 只接受第一个错误
	}
	return inst
}

func (inst *innerChecker) AddRoles(roles ...rbac.RoleName) permissions.Checker {
	inst.haveRoleList = append(inst.haveRoleList, roles...)
	return inst
}

///////////////////////////////////
// private

func (inst *innerChecker) init() permissions.Checker {
	err := inst.reload()
	inst.HandleError(err)
	return inst
}

func (inst *innerChecker) reload() error {
	inst.err = nil
	inst.loadSubjectInfo()
	inst.checkAuthState()
	return inst.err
}

func (inst *innerChecker) checkAuthState() {

	required := inst.builder.isAuthRequired
	if !required {
		return // 如果不是必须的, 就忽略这个检查
	}

	sub := &inst.subject
	uid := inst.CurrentUserID()
	if uid < 1 {
		inst.handleWebErrorCode(http.StatusUnauthorized)
	}
	if !sub.Authenticated {
		inst.handleWebErrorCode(http.StatusUnauthorized)
	}
}

func (inst *innerChecker) loadSubjectInfo() {

	ctx := inst.context
	subjects := inst.service.Subjects
	sub1, err := subjects.GetCurrent(ctx)

	if err == nil && sub1 != nil {
		inst.subject = *sub1
	} else {
		sub0 := rbac.SubjectDTO{}
		inst.subject = sub0
	}
}

func (inst *innerChecker) makeWebError(code int) error {
	if code < 1 {
		code = http.StatusForbidden
	}
	msg := http.StatusText(code)
	return web.NewError(code, msg)
}

func (inst *innerChecker) handleWebErrorCode(code int) {
	err := inst.makeWebError(code)
	inst.HandleError(err)
}

func (inst *innerChecker) innerCheckOwnership() {

	c1 := inst.countEntityWithOwnership
	c2 := inst.countEntityWithoutOwnership

	if c1 == 0 && c2 == 0 {
		r1 := rbac.RoleOwner
		r2 := permissions.RoleUserDoNothing
		inst.haveRoleList = append(inst.haveRoleList, r1, r2)
	}

	if c1 > 0 && c2 == 0 {
		owner := rbac.RoleOwner
		inst.haveRoleList = append(inst.haveRoleList, owner)
	}
}

func (inst *innerChecker) innerCheckRoles() {

	want := inst.builder.acceptedRoleList
	have := inst.haveRoleList
	table := make(map[rbac.RoleName]bool)

	for _, role := range have {
		table[role] = true
	}

	for _, role := range want {
		b := table[role]
		if b {
			return
		}
	}

	inst.handleWebErrorCode(http.StatusForbidden)
}

////////////////////////////////////////////////////////////////////////////////
// EOF
