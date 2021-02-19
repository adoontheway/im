package service

import (
	"fmt"
	"log"

	"im/model"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var DbEngin *xorm.Engine

func init() {
	drivername := "mysql"
	DsName := "root:12345@(127.0.0.1:3306)/chat?charset=utf8"
	var err error
	DbEngin, err = xorm.NewEngine(drivername, DsName)
	if err != nil {
		log.Fatal(err.Error())
	}
	// 是否显示sql
	DbEngin.ShowSQL(true)
	// 设置数据库的最大连接数
	DbEngin.SetMaxOpenConns(2)
	// 自动建表
	DbEngin.Sync2(new(model.User), new(model.Contact), new(model.Community))
	fmt.Println("init database ok")
}
