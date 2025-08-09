package apiv1

import (
	"strconv"

	"github.com/bitwormhole/lan-ca/backend/app/classes/certificates"
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

// CertificateController ...
type CertificateController struct {

	//starter:component
	_as func(libgin.Controller) //starter:as(".")

	Sender       libgin.Responder           //starter:inject("#")
	PermCheckers permissions.CheckerService //starter:inject('#')

	CertService certificates.Service //starter:inject("#")
}

func (inst *CertificateController) _impl() libgin.Controller {
	return inst
}

// Registration ...
func (inst *CertificateController) Registration() *libgin.ControllerRegistration {
	return &libgin.ControllerRegistration{Route: inst.route}
}

func (inst *CertificateController) route(rp libgin.RouterProxy) error {

	rp = rp.For("certificates")

	rp.POST("", inst.handlePostInsert)
	rp.POST("insert", inst.handlePostInsert)
	rp.PUT(":id", inst.handle)
	rp.DELETE(":id", inst.handle)

	rp.GET("", inst.handleGetList)
	rp.GET(":id", inst.handleGetOne)
	rp.GET("example", inst.handleGetMock)

	return nil
}

func (inst *CertificateController) handle(c *gin.Context) {
	req := &myCertificateRequest{
		context:       c,
		controller:    inst,
		wantRequestID: false,
	}
	req.execute(req.doNOP)
}

func (inst *CertificateController) handleGetOne(c *gin.Context) {
	req := &myCertificateRequest{
		context:       c,
		controller:    inst,
		wantRequestID: true,
	}
	req.execute(req.doGetOne)
}

func (inst *CertificateController) handleGetList(c *gin.Context) {
	req := &myCertificateRequest{
		context:         c,
		controller:      inst,
		wantRequestID:   false,
		wantRequestPage: true,
		wantPermChecker: true,
	}
	req.hChecker.AcceptRoles(rbac.RoleOwner)
	req.execute(req.doGetList)
}

func (inst *CertificateController) handleGetMock(c *gin.Context) {
	req := &myCertificateRequest{
		context:       c,
		controller:    inst,
		wantRequestID: false,
	}
	req.execute(req.doGetMock)
}

func (inst *CertificateController) handlePostInsert(c *gin.Context) {
	req := &myCertificateRequest{
		context:         c,
		controller:      inst,
		wantRequestID:   false,
		wantRequestBody: true,
		wantPermChecker: true,
	}
	req.hChecker.AcceptRoles(rbac.RoleUser)
	req.execute(req.doPostInsert)
}

////////////////////////////////////////////////////////////////////////////////

type myCertificateRequest struct {
	context    *gin.Context
	controller *CertificateController

	wantRequestID   bool
	wantRequestBody bool
	wantRequestPage bool
	wantPermChecker bool

	hChecker   permissions.CheckerHolder
	pagination rbac.Pagination
	id         dxo.CertificateID
	body1      vo.Certificates
	bodyTmp    vo.Certificates
	body2      vo.Certificates
}

func (inst *myCertificateRequest) open() error {

	c := inst.context

	if inst.wantRequestID {
		idstr := c.Param("id")
		idnum, err := strconv.ParseInt(idstr, 10, 64)
		if err != nil {
			return err
		}
		inst.id = dxo.CertificateID(idnum)
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

func (inst *myCertificateRequest) send(err error) {
	data := &inst.body2
	code := inst.body2.Status
	resp := new(libgin.Response)
	resp.Context = inst.context
	resp.Error = err
	resp.Data = data
	resp.Status = code
	inst.controller.Sender.Send(resp)
}

func (inst *myCertificateRequest) execute(fn func() error) {
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

func (inst *myCertificateRequest) doNOP() error {
	return nil
}

func (inst *myCertificateRequest) doGetMock() error {
	item := &dto.Certificate{}
	inst.bodyTmp.Items = []*dto.Certificate{item}
	return nil
}

func (inst *myCertificateRequest) doGetOne() error {
	id := inst.id
	ctx := inst.context
	ser := inst.controller.CertService
	o1, err := ser.Find(ctx, id)
	if err != nil {
		return err
	}
	inst.body2.Items = []*dto.Certificate{o1}
	return nil
}

func (inst *myCertificateRequest) doGetList() error {

	ctx := inst.context
	ser := inst.controller.CertService
	q := &certificates.Query{}
	want := &entity.Certificate{}

	checker := inst.hChecker.GetChecker()
	uid := checker.CurrentUserID()
	want.Owner = uid

	q.Pagination = inst.pagination
	q.User = uid
	q.Want = want

	list, err := ser.List(ctx, q)
	checker.HandleError(err)

	for _, item := range list {
		checker.CheckOwnerForDTO(&item.BaseDTO)
	}

	if checker.Error() == nil {
		inst.bodyTmp.Items = list
		inst.bodyTmp.Pagination = &q.Pagination
	}

	return nil
}

func (inst *myCertificateRequest) doPostInsert() error {

	ctx := inst.context
	ser := inst.controller.CertService
	checker := inst.hChecker.GetChecker()
	uid := checker.CurrentUserID()

	item1 := inst.body1.Items[0]
	item1.Owner = uid
	item1.Updater = uid
	item1.Creator = uid

	item2, err := ser.Insert(ctx, item1)
	checker.HandleError(err)
	if err != nil {
		return checker.Error()
	}
	checker.CheckOwnerForDTO(&item2.BaseDTO)

	if checker.HasError() {
		return checker.Error()
	}
	inst.bodyTmp.Items = []*dto.Certificate{item2}
	return nil
}
