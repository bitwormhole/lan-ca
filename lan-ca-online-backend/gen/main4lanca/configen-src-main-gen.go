package main4lanca
import (
    pe4760b7b1 "github.com/bitwormhole/lan-ca/backend/app/data/dao"
    p7d92875fe "github.com/bitwormhole/lan-ca/backend/app/data/database"
    p10d00aedf "github.com/bitwormhole/lan-ca/backend/app/data/dxo"
    p1f8806f62 "github.com/bitwormhole/lan-ca/backend/app/implements/example"
    p77d297856 "github.com/bitwormhole/lan-ca/backend/app/web/controllers/apiv1"
    pd1a916a20 "github.com/starter-go/libgin"
    p512a30914 "github.com/starter-go/libgorm"
     "github.com/starter-go/application"
)

// type p7d92875fe.ThisGroup in package:github.com/bitwormhole/lan-ca/backend/app/data/database
//
// id:com-7d92875fe9f1f5f0-database-ThisGroup
// class:class-512a309140d0ad99eb1c95c8dc0d02f9-GroupRegistry
// alias:alias-10d00aedf543c0d34cbdf4945116adfb-DatabaseAgent
// scope:singleton
//
type p7d92875fe9_database_ThisGroup struct {
}

func (inst* p7d92875fe9_database_ThisGroup) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-7d92875fe9f1f5f0-database-ThisGroup"
	r.Classes = "class-512a309140d0ad99eb1c95c8dc0d02f9-GroupRegistry"
	r.Aliases = "alias-10d00aedf543c0d34cbdf4945116adfb-DatabaseAgent"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p7d92875fe9_database_ThisGroup) new() any {
    return &p7d92875fe.ThisGroup{}
}

func (inst* p7d92875fe9_database_ThisGroup) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p7d92875fe.ThisGroup)
	nop(ie, com)

	
    com.Enabled = inst.getEnabled(ie)
    com.Alias = inst.getAlias(ie)
    com.URI = inst.getURI(ie)
    com.Prefix = inst.getPrefix(ie)
    com.Source = inst.getSource(ie)
    com.SourceManager = inst.getSourceManager(ie)


    return nil
}


func (inst*p7d92875fe9_database_ThisGroup) getEnabled(ie application.InjectionExt)bool{
    return ie.GetBool("${datagroup.lanca.enabled}")
}


func (inst*p7d92875fe9_database_ThisGroup) getAlias(ie application.InjectionExt)string{
    return ie.GetString("${datagroup.lanca.alias}")
}


func (inst*p7d92875fe9_database_ThisGroup) getURI(ie application.InjectionExt)string{
    return ie.GetString("${datagroup.lanca.uri}")
}


func (inst*p7d92875fe9_database_ThisGroup) getPrefix(ie application.InjectionExt)string{
    return ie.GetString("${datagroup.lanca.table-name-prefix}")
}


func (inst*p7d92875fe9_database_ThisGroup) getSource(ie application.InjectionExt)string{
    return ie.GetString("${datagroup.lanca.datasource}")
}


func (inst*p7d92875fe9_database_ThisGroup) getSourceManager(ie application.InjectionExt)p512a30914.DataSourceManager{
    return ie.GetComponent("#alias-512a309140d0ad99eb1c95c8dc0d02f9-DataSourceManager").(p512a30914.DataSourceManager)
}



// type p1f8806f62.DaoImpl in package:github.com/bitwormhole/lan-ca/backend/app/implements/example
//
// id:com-1f8806f6240e0b67-example-DaoImpl
// class:
// alias:alias-e4760b7b15bc9ea96b83a06f37708622-ExampleDAO
// scope:singleton
//
type p1f8806f624_example_DaoImpl struct {
}

func (inst* p1f8806f624_example_DaoImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-1f8806f6240e0b67-example-DaoImpl"
	r.Classes = ""
	r.Aliases = "alias-e4760b7b15bc9ea96b83a06f37708622-ExampleDAO"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p1f8806f624_example_DaoImpl) new() any {
    return &p1f8806f62.DaoImpl{}
}

func (inst* p1f8806f624_example_DaoImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p1f8806f62.DaoImpl)
	nop(ie, com)

	
    com.Agent = inst.getAgent(ie)


    return nil
}


func (inst*p1f8806f624_example_DaoImpl) getAgent(ie application.InjectionExt)p10d00aedf.DatabaseAgent{
    return ie.GetComponent("#alias-10d00aedf543c0d34cbdf4945116adfb-DatabaseAgent").(p10d00aedf.DatabaseAgent)
}



// type p77d297856.ExampleController in package:github.com/bitwormhole/lan-ca/backend/app/web/controllers/apiv1
//
// id:com-77d2978569998805-apiv1-ExampleController
// class:class-d1a916a203352fd5d33eabc36896b42e-Controller
// alias:
// scope:singleton
//
type p77d2978569_apiv1_ExampleController struct {
}

func (inst* p77d2978569_apiv1_ExampleController) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-77d2978569998805-apiv1-ExampleController"
	r.Classes = "class-d1a916a203352fd5d33eabc36896b42e-Controller"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p77d2978569_apiv1_ExampleController) new() any {
    return &p77d297856.ExampleController{}
}

func (inst* p77d2978569_apiv1_ExampleController) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p77d297856.ExampleController)
	nop(ie, com)

	
    com.Sender = inst.getSender(ie)
    com.Dao = inst.getDao(ie)


    return nil
}


func (inst*p77d2978569_apiv1_ExampleController) getSender(ie application.InjectionExt)pd1a916a20.Responder{
    return ie.GetComponent("#alias-d1a916a203352fd5d33eabc36896b42e-Responder").(pd1a916a20.Responder)
}


func (inst*p77d2978569_apiv1_ExampleController) getDao(ie application.InjectionExt)pe4760b7b1.ExampleDAO{
    return ie.GetComponent("#alias-e4760b7b15bc9ea96b83a06f37708622-ExampleDAO").(pe4760b7b1.ExampleDAO)
}


