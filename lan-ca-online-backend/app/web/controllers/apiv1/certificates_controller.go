package apiv1

import (
	"strconv"

	"github.com/bitwormhole/lan-ca/backend/app/classes/certificates"
	"github.com/bitwormhole/lan-ca/backend/app/data/dxo"
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

	Sender libgin.Responder //starter:inject("#")

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

	rp.POST("", inst.handle)
	rp.PUT(":id", inst.handle)
	rp.DELETE(":id", inst.handle)

	rp.GET("", inst.handleGetList)
	rp.GET(":id", inst.handleGetOne)

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
	}
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

////////////////////////////////////////////////////////////////////////////////

type myCertificateRequest struct {
	context    *gin.Context
	controller *CertificateController

	wantRequestID   bool
	wantRequestBody bool
	wantRequestPage bool

	pagination rbac.Pagination
	id         dxo.CertificateID
	body1      vo.Certificates
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
	inst.send(err)
}

func (inst *myCertificateRequest) doNOP() error {
	return nil
}

func (inst *myCertificateRequest) doGetMock() error {

	o2 := &dto.Certificate{
		ID: 0, //  o1.ID,

	}
	inst.body2.Items = []*dto.Certificate{o2}
	return nil
}

func (inst *myCertificateRequest) doGetOne() error {
	id := inst.id
	ser := inst.controller.CertService
	o1, err := ser.Find(nil, id)
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
	q.Pagination = inst.pagination
	list, err := ser.List(ctx, q)
	if err != nil {
		return err
	}
	inst.body2.Items = list
	inst.body2.Pagination = &q.Pagination
	return nil
}
