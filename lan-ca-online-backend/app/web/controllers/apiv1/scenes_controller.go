package apiv1

import (
	"strconv"

	"github.com/bitwormhole/lan-ca/backend/app/classes/scenes"
	"github.com/bitwormhole/lan-ca/backend/app/data/dxo"
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

	Sender libgin.Responder //starter:inject("#")

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

	rp.POST("", inst.handle)
	rp.PUT(":id", inst.handle)
	rp.DELETE(":id", inst.handle)

	rp.GET("", inst.handleGetList)
	rp.GET(":id", inst.handleGetOne)

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

func (inst *SceneController) handleGetOne(c *gin.Context) {
	req := &mySceneRequest{
		context:       c,
		controller:    inst,
		wantRequestID: true,
	}
	req.execute(req.doGetOne)
}

func (inst *SceneController) handleGetList(c *gin.Context) {
	req := &mySceneRequest{
		context:         c,
		controller:      inst,
		wantRequestID:   false,
		wantRequestPage: true,
	}
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

	pagination rbac.Pagination
	id         dxo.SceneID
	body1      vo.Scenes
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
	inst.send(err)
}

func (inst *mySceneRequest) doNOP() error {
	return nil
}

func (inst *mySceneRequest) doGetMock() error {

	o2 := &dto.Scene{
		ID: 0, //  o1.ID,

		// Name: "a.b.c",
		// IPv4: "0.0.0.0",
	}
	inst.body2.Items = []*dto.Scene{o2}
	return nil
}

func (inst *mySceneRequest) doGetOne() error {
	// id := inst.id
	// o1, err := inst.controller.Dao.Find(nil, id)
	// if err != nil {
	// 	return err
	// }

	o2 := &dto.Scene{
		ID: 0, //  o1.ID,

		// Name: "a.b.c",
		// IPv4: "0.0.0.0",
	}
	inst.body2.Items = []*dto.Scene{o2}
	return nil
}

func (inst *mySceneRequest) doGetList() error {
	ctx := inst.context
	ser := inst.controller.SceneService
	q := &scenes.Query{}
	q.Pagination = inst.pagination
	list, err := ser.List(ctx, q)
	if err != nil {
		return err
	}
	inst.body2.Items = list
	inst.body2.Pagination = &q.Pagination
	return nil
}
