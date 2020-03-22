package main

import (
	_ "antiCov-server/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

func main() {
	beego.Run()
}

func init() {

	//文件日志初始化
	err := logs.SetLogger(
		logs.AdapterFile,
		`{
		"filename":"log/server.log",
		"level":7,
		"maxlines":0,
		"maxsize":0,
		"daily":true,
		"maxdays":10
					}`,
	)
	if err != nil {
		logs.Error("文件日志初始化失败", err)
		os.Exit(0)
	}

	//数据库连接初始化
	err = initDatabase()
	if err != nil {
		logs.Error("数据库初始化失败")
	}
}

func initDatabase() error {
	//注册Mysql驱动
	err := orm.RegisterDriver("mysql", orm.DRMySQL)
	if err != nil {
		return err
	}

	//连接数据库
	err = orm.RegisterDataBase(
		"default",
		"mysql",
		"root:eduyibo2016@tcp(120.78.143.216)/anti_cov?charset=utf8mb4&loc=Asia%2FShanghai",
	)
	if err != nil {
		return err
	}

	//同步数据库结构
	err = orm.RunSyncdb("default", false, true)
	if err != nil {
		return err
	}
	return nil
}
