package main4lanca
import (
    p39f1aacbc "github.com/bitwormhole/lan-ca/backend/app/classes/certificates"
    pce737ddfd "github.com/bitwormhole/lan-ca/backend/app/classes/domains"
    pdb766c45a "github.com/bitwormhole/lan-ca/backend/app/classes/permissions"
    p7eb5e7931 "github.com/bitwormhole/lan-ca/backend/app/classes/scenes"
    pe4760b7b1 "github.com/bitwormhole/lan-ca/backend/app/data/dao"
    p7d92875fe "github.com/bitwormhole/lan-ca/backend/app/data/database"
    p10d00aedf "github.com/bitwormhole/lan-ca/backend/app/data/dxo"
    pfaa92fe24 "github.com/bitwormhole/lan-ca/backend/app/implements/icertificates"
    p07037744c "github.com/bitwormhole/lan-ca/backend/app/implements/idomains"
    p22b8c2eb7 "github.com/bitwormhole/lan-ca/backend/app/implements/iexample"
    pf3f9eff28 "github.com/bitwormhole/lan-ca/backend/app/implements/ipermissions"
    p782b4d905 "github.com/bitwormhole/lan-ca/backend/app/implements/iscenes"
    p77d297856 "github.com/bitwormhole/lan-ca/backend/app/web/controllers/apiv1"
    pd1a916a20 "github.com/starter-go/libgin"
    p512a30914 "github.com/starter-go/libgorm"
    p24287f458 "github.com/starter-go/rbac"
    p9621e8b71 "github.com/starter-go/security/random"
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



// type pfaa92fe24.CertificateDaoImpl in package:github.com/bitwormhole/lan-ca/backend/app/implements/icertificates
//
// id:com-faa92fe24168d630-icertificates-CertificateDaoImpl
// class:
// alias:alias-39f1aacbcc95c25190d9f6295f48486e-DAO
// scope:singleton
//
type pfaa92fe241_icertificates_CertificateDaoImpl struct {
}

func (inst* pfaa92fe241_icertificates_CertificateDaoImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-faa92fe24168d630-icertificates-CertificateDaoImpl"
	r.Classes = ""
	r.Aliases = "alias-39f1aacbcc95c25190d9f6295f48486e-DAO"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pfaa92fe241_icertificates_CertificateDaoImpl) new() any {
    return &pfaa92fe24.CertificateDaoImpl{}
}

func (inst* pfaa92fe241_icertificates_CertificateDaoImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pfaa92fe24.CertificateDaoImpl)
	nop(ie, com)

	
    com.Agent = inst.getAgent(ie)
    com.UUIDGen = inst.getUUIDGen(ie)


    return nil
}


func (inst*pfaa92fe241_icertificates_CertificateDaoImpl) getAgent(ie application.InjectionExt)p10d00aedf.DatabaseAgent{
    return ie.GetComponent("#alias-10d00aedf543c0d34cbdf4945116adfb-DatabaseAgent").(p10d00aedf.DatabaseAgent)
}


func (inst*pfaa92fe241_icertificates_CertificateDaoImpl) getUUIDGen(ie application.InjectionExt)p9621e8b71.UUIDService{
    return ie.GetComponent("#alias-9621e8b71013b0fc25942a1749ed3652-UUIDService").(p9621e8b71.UUIDService)
}



// type pfaa92fe24.CertificateServiceImpl in package:github.com/bitwormhole/lan-ca/backend/app/implements/icertificates
//
// id:com-faa92fe24168d630-icertificates-CertificateServiceImpl
// class:
// alias:alias-39f1aacbcc95c25190d9f6295f48486e-Service
// scope:singleton
//
type pfaa92fe241_icertificates_CertificateServiceImpl struct {
}

func (inst* pfaa92fe241_icertificates_CertificateServiceImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-faa92fe24168d630-icertificates-CertificateServiceImpl"
	r.Classes = ""
	r.Aliases = "alias-39f1aacbcc95c25190d9f6295f48486e-Service"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pfaa92fe241_icertificates_CertificateServiceImpl) new() any {
    return &pfaa92fe24.CertificateServiceImpl{}
}

func (inst* pfaa92fe241_icertificates_CertificateServiceImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pfaa92fe24.CertificateServiceImpl)
	nop(ie, com)

	
    com.Dao = inst.getDao(ie)


    return nil
}


func (inst*pfaa92fe241_icertificates_CertificateServiceImpl) getDao(ie application.InjectionExt)p39f1aacbc.DAO{
    return ie.GetComponent("#alias-39f1aacbcc95c25190d9f6295f48486e-DAO").(p39f1aacbc.DAO)
}



// type p07037744c.DomainDaoImpl in package:github.com/bitwormhole/lan-ca/backend/app/implements/idomains
//
// id:com-07037744c4ba7983-idomains-DomainDaoImpl
// class:
// alias:alias-ce737ddfdc57a545d8084b62ba0d0d14-DAO
// scope:singleton
//
type p07037744c4_idomains_DomainDaoImpl struct {
}

func (inst* p07037744c4_idomains_DomainDaoImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-07037744c4ba7983-idomains-DomainDaoImpl"
	r.Classes = ""
	r.Aliases = "alias-ce737ddfdc57a545d8084b62ba0d0d14-DAO"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p07037744c4_idomains_DomainDaoImpl) new() any {
    return &p07037744c.DomainDaoImpl{}
}

func (inst* p07037744c4_idomains_DomainDaoImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p07037744c.DomainDaoImpl)
	nop(ie, com)

	
    com.Agent = inst.getAgent(ie)
    com.UUIDGen = inst.getUUIDGen(ie)


    return nil
}


func (inst*p07037744c4_idomains_DomainDaoImpl) getAgent(ie application.InjectionExt)p10d00aedf.DatabaseAgent{
    return ie.GetComponent("#alias-10d00aedf543c0d34cbdf4945116adfb-DatabaseAgent").(p10d00aedf.DatabaseAgent)
}


func (inst*p07037744c4_idomains_DomainDaoImpl) getUUIDGen(ie application.InjectionExt)p9621e8b71.UUIDService{
    return ie.GetComponent("#alias-9621e8b71013b0fc25942a1749ed3652-UUIDService").(p9621e8b71.UUIDService)
}



// type p07037744c.DomainServiceImpl in package:github.com/bitwormhole/lan-ca/backend/app/implements/idomains
//
// id:com-07037744c4ba7983-idomains-DomainServiceImpl
// class:
// alias:alias-ce737ddfdc57a545d8084b62ba0d0d14-Service
// scope:singleton
//
type p07037744c4_idomains_DomainServiceImpl struct {
}

func (inst* p07037744c4_idomains_DomainServiceImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-07037744c4ba7983-idomains-DomainServiceImpl"
	r.Classes = ""
	r.Aliases = "alias-ce737ddfdc57a545d8084b62ba0d0d14-Service"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p07037744c4_idomains_DomainServiceImpl) new() any {
    return &p07037744c.DomainServiceImpl{}
}

func (inst* p07037744c4_idomains_DomainServiceImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p07037744c.DomainServiceImpl)
	nop(ie, com)

	
    com.Dao = inst.getDao(ie)
    com.Subjects = inst.getSubjects(ie)


    return nil
}


func (inst*p07037744c4_idomains_DomainServiceImpl) getDao(ie application.InjectionExt)pce737ddfd.DAO{
    return ie.GetComponent("#alias-ce737ddfdc57a545d8084b62ba0d0d14-DAO").(pce737ddfd.DAO)
}


func (inst*p07037744c4_idomains_DomainServiceImpl) getSubjects(ie application.InjectionExt)p24287f458.SubjectService{
    return ie.GetComponent("#alias-24287f4589fe5add27fb48a88d706565-SubjectService").(p24287f458.SubjectService)
}



// type p22b8c2eb7.DaoImpl in package:github.com/bitwormhole/lan-ca/backend/app/implements/iexample
//
// id:com-22b8c2eb74f73c94-iexample-DaoImpl
// class:
// alias:alias-e4760b7b15bc9ea96b83a06f37708622-ExampleDAO
// scope:singleton
//
type p22b8c2eb74_iexample_DaoImpl struct {
}

func (inst* p22b8c2eb74_iexample_DaoImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-22b8c2eb74f73c94-iexample-DaoImpl"
	r.Classes = ""
	r.Aliases = "alias-e4760b7b15bc9ea96b83a06f37708622-ExampleDAO"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p22b8c2eb74_iexample_DaoImpl) new() any {
    return &p22b8c2eb7.DaoImpl{}
}

func (inst* p22b8c2eb74_iexample_DaoImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p22b8c2eb7.DaoImpl)
	nop(ie, com)

	
    com.Agent = inst.getAgent(ie)


    return nil
}


func (inst*p22b8c2eb74_iexample_DaoImpl) getAgent(ie application.InjectionExt)p10d00aedf.DatabaseAgent{
    return ie.GetComponent("#alias-10d00aedf543c0d34cbdf4945116adfb-DatabaseAgent").(p10d00aedf.DatabaseAgent)
}



// type pf3f9eff28.CheckerServiceImpl in package:github.com/bitwormhole/lan-ca/backend/app/implements/ipermissions
//
// id:com-f3f9eff282c69889-ipermissions-CheckerServiceImpl
// class:
// alias:alias-db766c45ad9bffb478a55f076577ab20-CheckerService
// scope:singleton
//
type pf3f9eff282_ipermissions_CheckerServiceImpl struct {
}

func (inst* pf3f9eff282_ipermissions_CheckerServiceImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-f3f9eff282c69889-ipermissions-CheckerServiceImpl"
	r.Classes = ""
	r.Aliases = "alias-db766c45ad9bffb478a55f076577ab20-CheckerService"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pf3f9eff282_ipermissions_CheckerServiceImpl) new() any {
    return &pf3f9eff28.CheckerServiceImpl{}
}

func (inst* pf3f9eff282_ipermissions_CheckerServiceImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pf3f9eff28.CheckerServiceImpl)
	nop(ie, com)

	
    com.Subjects = inst.getSubjects(ie)


    return nil
}


func (inst*pf3f9eff282_ipermissions_CheckerServiceImpl) getSubjects(ie application.InjectionExt)p24287f458.SubjectService{
    return ie.GetComponent("#alias-24287f4589fe5add27fb48a88d706565-SubjectService").(p24287f458.SubjectService)
}



// type p782b4d905.SceneDaoImpl in package:github.com/bitwormhole/lan-ca/backend/app/implements/iscenes
//
// id:com-782b4d9055a66595-iscenes-SceneDaoImpl
// class:
// alias:alias-7eb5e79313db1ba31143b12200682323-DAO
// scope:singleton
//
type p782b4d9055_iscenes_SceneDaoImpl struct {
}

func (inst* p782b4d9055_iscenes_SceneDaoImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-782b4d9055a66595-iscenes-SceneDaoImpl"
	r.Classes = ""
	r.Aliases = "alias-7eb5e79313db1ba31143b12200682323-DAO"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p782b4d9055_iscenes_SceneDaoImpl) new() any {
    return &p782b4d905.SceneDaoImpl{}
}

func (inst* p782b4d9055_iscenes_SceneDaoImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p782b4d905.SceneDaoImpl)
	nop(ie, com)

	
    com.Agent = inst.getAgent(ie)
    com.UUIDGen = inst.getUUIDGen(ie)


    return nil
}


func (inst*p782b4d9055_iscenes_SceneDaoImpl) getAgent(ie application.InjectionExt)p10d00aedf.DatabaseAgent{
    return ie.GetComponent("#alias-10d00aedf543c0d34cbdf4945116adfb-DatabaseAgent").(p10d00aedf.DatabaseAgent)
}


func (inst*p782b4d9055_iscenes_SceneDaoImpl) getUUIDGen(ie application.InjectionExt)p9621e8b71.UUIDService{
    return ie.GetComponent("#alias-9621e8b71013b0fc25942a1749ed3652-UUIDService").(p9621e8b71.UUIDService)
}



// type p782b4d905.SceneServiceImpl in package:github.com/bitwormhole/lan-ca/backend/app/implements/iscenes
//
// id:com-782b4d9055a66595-iscenes-SceneServiceImpl
// class:
// alias:alias-7eb5e79313db1ba31143b12200682323-Service
// scope:singleton
//
type p782b4d9055_iscenes_SceneServiceImpl struct {
}

func (inst* p782b4d9055_iscenes_SceneServiceImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-782b4d9055a66595-iscenes-SceneServiceImpl"
	r.Classes = ""
	r.Aliases = "alias-7eb5e79313db1ba31143b12200682323-Service"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p782b4d9055_iscenes_SceneServiceImpl) new() any {
    return &p782b4d905.SceneServiceImpl{}
}

func (inst* p782b4d9055_iscenes_SceneServiceImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p782b4d905.SceneServiceImpl)
	nop(ie, com)

	
    com.Dao = inst.getDao(ie)


    return nil
}


func (inst*p782b4d9055_iscenes_SceneServiceImpl) getDao(ie application.InjectionExt)p7eb5e7931.DAO{
    return ie.GetComponent("#alias-7eb5e79313db1ba31143b12200682323-DAO").(p7eb5e7931.DAO)
}



// type p77d297856.CertificateController in package:github.com/bitwormhole/lan-ca/backend/app/web/controllers/apiv1
//
// id:com-77d2978569998805-apiv1-CertificateController
// class:class-d1a916a203352fd5d33eabc36896b42e-Controller
// alias:
// scope:singleton
//
type p77d2978569_apiv1_CertificateController struct {
}

func (inst* p77d2978569_apiv1_CertificateController) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-77d2978569998805-apiv1-CertificateController"
	r.Classes = "class-d1a916a203352fd5d33eabc36896b42e-Controller"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p77d2978569_apiv1_CertificateController) new() any {
    return &p77d297856.CertificateController{}
}

func (inst* p77d2978569_apiv1_CertificateController) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p77d297856.CertificateController)
	nop(ie, com)

	
    com.Sender = inst.getSender(ie)
    com.PermCheckers = inst.getPermCheckers(ie)
    com.CertService = inst.getCertService(ie)


    return nil
}


func (inst*p77d2978569_apiv1_CertificateController) getSender(ie application.InjectionExt)pd1a916a20.Responder{
    return ie.GetComponent("#alias-d1a916a203352fd5d33eabc36896b42e-Responder").(pd1a916a20.Responder)
}


func (inst*p77d2978569_apiv1_CertificateController) getPermCheckers(ie application.InjectionExt)pdb766c45a.CheckerService{
    return ie.GetComponent("#alias-db766c45ad9bffb478a55f076577ab20-CheckerService").(pdb766c45a.CheckerService)
}


func (inst*p77d2978569_apiv1_CertificateController) getCertService(ie application.InjectionExt)p39f1aacbc.Service{
    return ie.GetComponent("#alias-39f1aacbcc95c25190d9f6295f48486e-Service").(p39f1aacbc.Service)
}



// type p77d297856.DomainsController in package:github.com/bitwormhole/lan-ca/backend/app/web/controllers/apiv1
//
// id:com-77d2978569998805-apiv1-DomainsController
// class:class-d1a916a203352fd5d33eabc36896b42e-Controller
// alias:
// scope:singleton
//
type p77d2978569_apiv1_DomainsController struct {
}

func (inst* p77d2978569_apiv1_DomainsController) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-77d2978569998805-apiv1-DomainsController"
	r.Classes = "class-d1a916a203352fd5d33eabc36896b42e-Controller"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p77d2978569_apiv1_DomainsController) new() any {
    return &p77d297856.DomainsController{}
}

func (inst* p77d2978569_apiv1_DomainsController) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p77d297856.DomainsController)
	nop(ie, com)

	
    com.Sender = inst.getSender(ie)
    com.PermCheckers = inst.getPermCheckers(ie)
    com.DomainService = inst.getDomainService(ie)


    return nil
}


func (inst*p77d2978569_apiv1_DomainsController) getSender(ie application.InjectionExt)pd1a916a20.Responder{
    return ie.GetComponent("#alias-d1a916a203352fd5d33eabc36896b42e-Responder").(pd1a916a20.Responder)
}


func (inst*p77d2978569_apiv1_DomainsController) getPermCheckers(ie application.InjectionExt)pdb766c45a.CheckerService{
    return ie.GetComponent("#alias-db766c45ad9bffb478a55f076577ab20-CheckerService").(pdb766c45a.CheckerService)
}


func (inst*p77d2978569_apiv1_DomainsController) getDomainService(ie application.InjectionExt)pce737ddfd.Service{
    return ie.GetComponent("#alias-ce737ddfdc57a545d8084b62ba0d0d14-Service").(pce737ddfd.Service)
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



// type p77d297856.SceneController in package:github.com/bitwormhole/lan-ca/backend/app/web/controllers/apiv1
//
// id:com-77d2978569998805-apiv1-SceneController
// class:class-d1a916a203352fd5d33eabc36896b42e-Controller
// alias:
// scope:singleton
//
type p77d2978569_apiv1_SceneController struct {
}

func (inst* p77d2978569_apiv1_SceneController) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-77d2978569998805-apiv1-SceneController"
	r.Classes = "class-d1a916a203352fd5d33eabc36896b42e-Controller"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p77d2978569_apiv1_SceneController) new() any {
    return &p77d297856.SceneController{}
}

func (inst* p77d2978569_apiv1_SceneController) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p77d297856.SceneController)
	nop(ie, com)

	
    com.Sender = inst.getSender(ie)
    com.PermCheckers = inst.getPermCheckers(ie)
    com.SceneService = inst.getSceneService(ie)


    return nil
}


func (inst*p77d2978569_apiv1_SceneController) getSender(ie application.InjectionExt)pd1a916a20.Responder{
    return ie.GetComponent("#alias-d1a916a203352fd5d33eabc36896b42e-Responder").(pd1a916a20.Responder)
}


func (inst*p77d2978569_apiv1_SceneController) getPermCheckers(ie application.InjectionExt)pdb766c45a.CheckerService{
    return ie.GetComponent("#alias-db766c45ad9bffb478a55f076577ab20-CheckerService").(pdb766c45a.CheckerService)
}


func (inst*p77d2978569_apiv1_SceneController) getSceneService(ie application.InjectionExt)p7eb5e7931.Service{
    return ie.GetComponent("#alias-7eb5e79313db1ba31143b12200682323-Service").(p7eb5e7931.Service)
}


