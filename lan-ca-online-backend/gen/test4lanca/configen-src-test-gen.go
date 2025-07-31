package test4lanca
import (
    pf399d19d1 "github.com/bitwormhole/lan-ca/backend/src/test/golang/unit"
     "github.com/starter-go/application"
)

// type pf399d19d1.DemoUnit in package:github.com/bitwormhole/lan-ca/backend/src/test/golang/unit
//
// id:com-f399d19d1f61bb86-unit-DemoUnit
// class:class-0dc072ed44b3563882bff4e657a52e62-Units
// alias:
// scope:singleton
//
type pf399d19d1f_unit_DemoUnit struct {
}

func (inst* pf399d19d1f_unit_DemoUnit) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-f399d19d1f61bb86-unit-DemoUnit"
	r.Classes = "class-0dc072ed44b3563882bff4e657a52e62-Units"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pf399d19d1f_unit_DemoUnit) new() any {
    return &pf399d19d1.DemoUnit{}
}

func (inst* pf399d19d1f_unit_DemoUnit) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pf399d19d1.DemoUnit)
	nop(ie, com)

	


    return nil
}


