package apiv1

import (
	"strconv"

	"github.com/bitwormhole/lan-ca/backend/app/classes/permissions"
	"github.com/bitwormhole/lan-ca/backend/app/classes/scenes"
	"github.com/bitwormhole/lan-ca/backend/app/data/dxo"
	"github.com/bitwormhole/lan-ca/backend/app/data/entity"
	"github.com/bitwormhole/lan-ca/backend/app/web/dto"
	"github.com/bitwormhole/lan-ca/backend/app/web/vo"
	"github.com/gin-gonic/gin"
	"github.com/starter-go/libgin"
	"github.com/starter-go/rbac"
	"github.com/starter-go/security-gin-gorm/components/web/controllers"
)

// SceneController ...
type SceneController struct {

	//starter:component
	_as func(libgin.Controller) //starter:as(".")

	Sender       libgin.Responder           //starter:inject("#")
	PermCheckers permissions.CheckerService //starter:inject('#')

	SceneService scenes.Service //starter:inject("#")
}

func (inst *SceneController) _impl() libgin.Controller {
	return inst
}

// Registration ...
func (inst *SceneController) Registration() *libgin.ControllerRegistration {
	return &libgin.ControllerRegistration{Route: inst.route}
}

func (inst *SceneController) route(rp libgin.RouterProxy) error {

	rp = rp.For("scenes")

	rp.POST("", inst.handlePost)
	rp.PUT(":id", inst.handle)
	rp.DELETE(":id", inst.handle)

	rp.GET("", inst.handleGetList)
	rp.GET(":id", inst.handleGetOne)
	rp.GET("example", inst.handleGetMock)

	return nil
}

func (inst *SceneController) handle(c *gin.Context) {
	req := &mySceneRequest{
		context:       c,
		controller:    inst,
		wantRequestID: false,
	}
	req.execute(req.doNOP)
}

func (inst *SceneController) handlePost(c *gin.Context) {
	req := &mySceneRequest{
		context:         c,
		controller:      inst,
		wantRequestID:   false,
		wantRequestBody: true,
		wantPermChecker: true,
	}
	req.hChecker.AcceptRoles(rbac.RoleOwner)
	req.execute(req.doPostInsert)
}

func (inst *SceneController) handleGetOne(c *gin.Context) {
	req := &mySceneRequest{
		context:         c,
		controller:      inst,
		wantRequestID:   true,
		wantPermChecker: true,
	}
	req.hChecker.AcceptRoles(rbac.RoleOwner)
	req.execute(req.doGetOne)
}

func (inst *SceneController) handleGetList(c *gin.Context) {
	req := &mySceneRequest{
		context:         c,
		controller:      inst,
		wantRequestID:   false,
		wantRequestPage: true,
		wantPermChecker: true,
	}
	req.hChecker.AcceptRoles(rbac.RoleOwner)
	req.execute(req.doGetList)
}

func (inst *SceneController) handleGetMock(c *gin.Context) {
	req := &mySceneRequest{
		context:       c,
		controller:    inst,
		wantRequestID: false,
	}
	req.execute(req.doGetMock)
}

////////////////////////////////////////////////////////////////////////////////

type mySceneRequest struct {
	context    *gin.Context
	controller *SceneController

	wantRequestID   bool
	wantRequestBody bool
	wantRequestPage bool
	wantPermChecker bool

	hChecker   permissions.CheckerHolder
	pagination rbac.Pagination
	id         dxo.SceneID
	body1      vo.Scenes
	bodyTmp    vo.Scenes
	body2      vo.Scenes
}

func (inst *mySceneRequest) open() error {

	c := inst.context

	if inst.wantRequestID {
		idstr := c.Param("id")
		idnum, err := strconv.ParseInt(idstr, 10, 64)
		if err != nil {
			return err
		}
		inst.id = dxo.SceneID(idnum)
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

func (inst *mySceneRequest) send(err error) {
	data := &inst.body2
	code := inst.body2.Status
	resp := new(libgin.Response)
	resp.Context = inst.context
	resp.Error = err
	resp.Data = data
	resp.Status = code
	inst.controller.Sender.Send(resp)
}

func (inst *mySceneRequest) execute(fn func() error) {
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

func (inst *mySceneRequest) doNOP() error {
	return nil
}

func (inst *mySceneRequest) doGetMock() error {
	item := &dto.Scene{}
	inst.bodyTmp.Items = []*dto.Scene{item}
	return nil
}

func (inst *mySceneRequest) doGetOne() error {

	ctx := inst.context
	ser := inst.controller.SceneService
	id := inst.id
	checker := inst.hChecker.GetChecker()

	item, err := ser.Find(ctx, id)
	if err != nil {
		checker.HandleError(err)
		return err
	}

	checker.CheckOwnerForDTO(&item.BaseDTO)
	inst.bodyTmp.Items = []*dto.Scene{item}
	inst.bodyTmp.Pagination = nil
	return nil
}

func (inst *mySceneRequest) doGetList() error {

	ctx := inst.context
	ser := inst.controller.SceneService
	q := &scenes.Query{}
	want := &entity.Scene{}

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

func (inst *mySceneRequest) doPostInsert() error {

	ctx := inst.context
	ser := inst.controller.SceneService
	checker := inst.hChecker.GetChecker()
	uid := checker.CurrentUserID()

	list1 := inst.body1.Items
	item1 := list1[0]
	item1.Owner = uid
	item1.Creator = uid
	item1.Updater = uid

	item2, err := ser.Insert(ctx, item1)
	if err != nil {
		checker.HandleError(err)
		return err
	}

	checker.CheckOwnerForDTO(&item2.BaseDTO)
	inst.bodyTmp.Items = []*dto.Scene{item2}
	inst.bodyTmp.Pagination = nil
	return nil
}
