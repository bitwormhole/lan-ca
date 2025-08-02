package main4lanca

import "github.com/starter-go/application"

func nop(a ... any) {    
}

func registerComponents(cr application.ComponentRegistry) error {
    ac:=&autoRegistrar{}
    ac.init(cr)
    return ac.addAll()
}

type comFactory interface {
    register(cr application.ComponentRegistry) error
}

type autoRegistrar struct {
    cr application.ComponentRegistry
}

func (inst *autoRegistrar) init(cr application.ComponentRegistry) {
	inst.cr = cr
}

func (inst *autoRegistrar) register(factory comFactory) error {
	return factory.register(inst.cr)
}

func (inst*autoRegistrar) addAll() error {

    
    inst.register(&p07037744c4_idomains_DomainDaoImpl{})
    inst.register(&p07037744c4_idomains_DomainServiceImpl{})
    inst.register(&p22b8c2eb74_iexample_DaoImpl{})
    inst.register(&p77d2978569_apiv1_CertificateController{})
    inst.register(&p77d2978569_apiv1_DomainsController{})
    inst.register(&p77d2978569_apiv1_ExampleController{})
    inst.register(&p77d2978569_apiv1_SceneController{})
    inst.register(&p782b4d9055_iscenes_SceneDaoImpl{})
    inst.register(&p782b4d9055_iscenes_SceneServiceImpl{})
    inst.register(&p7d92875fe9_database_ThisGroup{})
    inst.register(&pf3f9eff282_ipermissions_CheckerServiceImpl{})
    inst.register(&pfaa92fe241_icertificates_CertificateDaoImpl{})
    inst.register(&pfaa92fe241_icertificates_CertificateServiceImpl{})


    return nil
}
