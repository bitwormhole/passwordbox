package idb

import (
	"github.com/bitwormhole/passwordbox/app/data/entity"
	"github.com/starter-go/libgorm"
)

type MyDataGroup struct {

	//starter:component

	_as func(libgorm.GroupRegistry) //starter:as(".")

	Alias   string //starter:inject("${datagroup.passwordbox.alias}")
	Prefix  string //starter:inject("${datagroup.passwordbox.table-name-prefix}")
	Source  string //starter:inject("${datagroup.passwordbox.datasource}")
	URI     string //starter:inject("${datagroup.passwordbox.uri}")
	Enabled bool   //starter:inject("${datagroup.passwordbox.enabled}")
}

func (inst *MyDataGroup) _impl() (libgorm.GroupRegistry, libgorm.Group) {
	return inst, inst
}

func (inst *MyDataGroup) Groups() []*libgorm.GroupRegistration {

	r1 := new(libgorm.GroupRegistration)
	r1.Alias = inst.Alias
	r1.Enabled = inst.Enabled
	r1.Prefix = inst.Prefix
	r1.Source = inst.Source
	r1.URI = inst.URI
	r1.Group = inst

	return []*libgorm.GroupRegistration{r1}
}

func (inst *MyDataGroup) Prototypes() []any {
	prefix := inst.Prefix
	all := entity.ListAll(prefix)
	return all
}
