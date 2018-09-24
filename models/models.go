package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var engine *xorm.Engine

func DBConnect() (db *xorm.Engine) {
	if engine != nil {
		return engine
	}
	driver := "mysql"
	user := "user"
	pass := "user"
	name := "todo"
	host := "tcp(db:3306)"
	var err error
	engine, err = xorm.NewEngine(driver, user+":"+pass+"@"+host+"/"+name)
	if err != nil {
		panic(err.Error())
	}
	return engine
}
