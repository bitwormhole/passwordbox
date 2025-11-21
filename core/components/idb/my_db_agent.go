package idb

import (
	"fmt"

	"github.com/bitwormhole/passwordbox/app/data/database"
	"github.com/starter-go/libgorm"
	"gorm.io/gorm"
)

////////////////////////////////////////////////////////////////////////////////

type MyDatabaseAgent struct {

	//starter:component

	_as func(database.Agent) //starter:as("#")

	DSM libgorm.DataSourceManager //starter:inject("#")

	inner *libgorm.DataSourceAgent
}

func (inst *MyDatabaseAgent) _impl() database.Agent {
	return inst
}

func (inst *MyDatabaseAgent) loadInner() (*libgorm.DataSourceAgent, error) {

	agent := new(libgorm.DataSourceAgent)
	man := inst.DSM
	name := "main"

	agent.Init(man, name)

	var db *gorm.DB = nil
	db = agent.DB(db)
	if db == nil {
		return nil, fmt.Errorf("idb.MyDatabaseAgent: no gorm.DB")
	}

	return agent, nil
}

func (inst *MyDatabaseAgent) getInner() (*libgorm.DataSourceAgent, error) {
	in := inst.inner
	if in == nil {
		in, err := inst.loadInner()
		if err != nil {
			return nil, err
		}
		inst.inner = in
	}
	return in, nil
}

func (inst *MyDatabaseAgent) DB(db *gorm.DB) *gorm.DB {
	in, err := inst.getInner()
	if err != nil {
		panic(err)
	}
	return in.DB(db)
}
