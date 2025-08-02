package ipermissions

import (
	"context"
	"fmt"

	"github.com/bitwormhole/lan-ca/backend/app/classes/permissions"
	"github.com/starter-go/rbac"
	"github.com/starter-go/security-gorm/rbacdb"
)

////////////////////////////////////////////////////////////////////////////////
// CheckerServiceImpl

type CheckerServiceImpl struct {

	//starter:component

	_as func(permissions.Checker) //starter:as("#")

	Subjects rbac.SubjectService //starter:inject("#")

}

func (inst *CheckerServiceImpl) _impl() permissions.Checker {
	return inst
}

func (inst *CheckerServiceImpl) NewChecking(ctx context.Context) permissions.Checking {
	checking := &innerChecking{
		context: ctx,
		service: inst,
	}
	return checking.init()
}

////////////////////////////////////////////////////////////////////////////////
// innerChecking

type innerChecking struct {
	service        *CheckerServiceImpl
	context        context.Context
	err            error
	isAuthRequired bool
	subject        rbac.SubjectDTO
}

func (inst *innerChecking) _impl() permissions.Checking {
	return inst
}

func (inst *innerChecking) handleError(err error) permissions.Checking {
	if err != nil {
		inst.err = err
	}
	return inst
}

func (inst *innerChecking) init() permissions.Checking {

	sub0 := rbac.SubjectDTO{}
	ctx := inst.context
	subjects := inst.service.Subjects

	inst.err = nil
	inst.isAuthRequired = true
	inst.subject = sub0

	sub1, err := subjects.GetCurrent(ctx)
	inst.handleError(err)

	if err == nil && sub1 != nil {
		inst.subject = *sub1
	}

	return inst
}

func (inst *innerChecking) Context() context.Context {
	return inst.context
}

func (inst *innerChecking) HasError() bool {
	return (inst.err != nil)
}

func (inst *innerChecking) Error() error {
	return inst.err
}

func (inst *innerChecking) CurrentUserID() rbac.UserID {
	return inst.subject.Owner
}

func (inst *innerChecking) Reset() permissions.Checking {
	return inst.init()
}

func (inst *innerChecking) SetAuthIsRequired(required bool) permissions.Checking {
	inst.isAuthRequired = required
	return inst
}

func (inst *innerChecking) CheckDTO(o *rbac.BaseDTO) permissions.Checking {
	if o == nil {
		return inst
	}
	if o.Owner != inst.subject.Owner {
		inst.handleError(fmt.Errorf("no permission"))
	}
	return inst
}

func (inst *innerChecking) CheckEntity(o *rbacdb.BaseEntity) permissions.Checking {
	if o == nil {
		return inst
	}
	if o.Owner != inst.subject.Owner {
		inst.handleError(fmt.Errorf("no permission"))
	}
	return inst
}

func (inst *innerChecking) Check() error {
	if inst.isAuthRequired {
		sub := &inst.subject
		if sub.Owner < 1 {
			inst.handleError(fmt.Errorf("auth is required"))
		}
		if !sub.Authenticated {
			inst.handleError(fmt.Errorf("auth is required"))
		}
	}
	return inst.err
}

////////////////////////////////////////////////////////////////////////////////
// EOF
