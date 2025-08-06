package apiv1

import (
	"strconv"

	"github.com/bitwormhole/lan-ca/backend/app/classes/domains"
	"github.com/bitwormhole/lan-ca/backend/app/classes/permissions"
	"github.com/bitwormhole/lan-ca/backend/app/data/dxo"
	"github.com/bitwormhole/lan-ca/backend/app/data/entity"
	"github.com/bitwormhole/lan-ca/backend/app/web/dto"
	"github.com/bitwormhole/lan-ca/backend/app/web/vo"
	"github.com/gin-gonic/gin"
	"github.com/starter-go/libgin"
	"github.com/starter-go/rbac"
	"github.com/starter-go/security-gin-gorm/components/web/controllers"
)

// DomainsController ...
type DomainsController struct {

	//starter:component
	_as func(libgin.Controller) //starter:as(".")

	Sender       libgin.Responder           //starter:inject("#")
	PermCheckers permissions.CheckerService //starter:inject('#')

	DomainService domains.Service //starter:inject("#")
}

func (inst *DomainsController) _impl() libgin.Controller {
	return inst
}

// Registration ...
func (inst *DomainsController) Registration() *libgin.ControllerRegistration {
	return &libgin.ControllerRegistration{Route: inst.route}
}

func (inst *DomainsController) route(rp libgin.RouterProxy) error {

	rp = rp.For("domains")

	rp.POST("", inst.handle)
	rp.PUT(":id", inst.handle)
	rp.DELETE(":id", inst.handle)

	rp.GET("", inst.handleGetList)
	rp.GET(":id", inst.handleGetOne)
	rp.GET("example", inst.handleGetMock)

	return nil
}

func (inst *DomainsController) handle(c *gin.Context) {
	req := &myDomainsRequest{
		context:       c,
		controller:    inst,
		wantRequestID: false,
	}
	req.execute(req.doNOP)
}

func (inst *DomainsController) handleGetOne(c *gin.Context) {
	req := &myDomainsRequest{
		context:       c,
		controller:    inst,
		wantRequestID: true,
	}
	req.execute(req.doGetOne)
}

func (inst *DomainsController) handleGetList(c *gin.Context) {
	req := &myDomainsRequest{
		context:         c,
		controller:      inst,
		wantRequestID:   false,
		wantRequestPage: true,
		wantPermChecker: true,
	}
	req.hChecker.AcceptRoles(rbac.RoleOwner)
	req.execute(req.doGetList)
}

func (inst *DomainsController) handleGetMock(c *gin.Context) {
	req := &myDomainsRequest{
		context:       c,
		controller:    inst,
		wantRequestID: false,
	}
	req.execute(req.doGetMock)
}

////////////////////////////////////////////////////////////////////////////////

type myDomainsRequest struct {
	context    *gin.Context
	controller *DomainsController

	wantRequestID   bool
	wantRequestBody bool
	wantRequestPage bool
	wantPermChecker bool

	hChecker   permissions.CheckerHolder
	pagination rbac.Pagination
	id         dxo.DomainID
	body1      vo.Domains
	bodyTmp    vo.Domains
	body2      vo.Domains
}

func (inst *myDomainsRequest) open() error {

	c := inst.context

	if inst.wantRequestID {
		idstr := c.Param("id")
		idnum, err := strconv.ParseInt(idstr, 10, 64)
		if err != nil {
			return err
		}
		inst.id = dxo.DomainID(idnum)
	}

	if inst.wantRequestBody {
		err := c.BindJSON(&inst.body1)
		if err != nil {
			return err
		}
	}

	if inst.wantRequestPage {
		inst.pagination = controllers.GetPagination(c)
	}

	if inst.wantPermChecker {
		pc_ser := inst.controller.PermCheckers
		err := inst.hChecker.Init(c, pc_ser)
		if err != nil {
			return err
		}
	}

	return nil
}

func (inst *myDomainsRequest) send(err error) {
	data := &inst.body2
	code := inst.body2.Status
	resp := new(libgin.Response)
	resp.Context = inst.context
	resp.Error = err
	resp.Data = data
	resp.Status = code
	inst.controller.Sender.Send(resp)
}

func (inst *myDomainsRequest) execute(fn func() error) {
	err := inst.open()
	if err == nil {
		err = fn()
	}
	if inst.wantPermChecker {
		err = inst.hChecker.Done(err)
	}
	if err == nil {
		inst.body2.Items = inst.bodyTmp.Items
		inst.body2.Pagination = inst.bodyTmp.Pagination
	}
	inst.send(err)
}

func (inst *myDomainsRequest) doNOP() error {
	return nil
}

func (inst *myDomainsRequest) doGetMock() error {
	item := &dto.Domain{}
	inst.bodyTmp.Items = []*dto.Domain{item}
	return nil
}

func (inst *myDomainsRequest) doGetOne() error {
	// id := inst.id
	// o1, err := inst.controller.Dao.Find(nil, id)
	// if err != nil {
	// 	return err
	// }

	o2 := &dto.Domain{
		ID: 0, //  o1.ID,

		Name: "a.b.c",
		IPv4: "0.0.0.0",
	}
	inst.body2.Items = []*dto.Domain{o2}
	return nil
}

func (inst *myDomainsRequest) doGetList() error {

	ctx := inst.context
	ser := inst.controller.DomainService
	q := &domains.Query{}
	want := &entity.Domain{}

	checker := inst.hChecker.GetChecker()
	uid := checker.CurrentUserID()
	want.Owner = uid

	q.Pagination = inst.pagination
	q.User = uid
	q.Want = want

	list, err := ser.List(ctx, q)
	checker.HandleError(err)

	if checker.Error() == nil {
		inst.bodyTmp.Items = list
		inst.bodyTmp.Pagination = &q.Pagination
	}

	return nil
}
