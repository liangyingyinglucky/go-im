package service

import (
	"xorm.io/xorm"
	"log"
	"fmt"
	"errors"
	"../model"
)

var DbEngin *xorm.Engine

func  init()  {
	drivename :="mysql"
	DsName := "root:Lyy@@@123@(127.0.0.1:3306)/test?charset=utf8"
	err := errors.New("")
	DbEngin,err = xorm.NewEngine(drivename,DsName)
	if nil!=err && ""!=err.Error() {
		log.Fatal(err.Error())
	}
	//是否显示SQL语句
	DbEngin.ShowSQL(true)
	//数据库最大打开的连接数
	DbEngin.SetMaxOpenConns(2)

	//检查数据库表，没有则创建
	DbEngin.Sync2(new(model.User),
		new(model.Contact),
		new(model.Community))
	
	fmt.Println("init data base ok")
}

