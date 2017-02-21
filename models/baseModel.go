package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
)

const (
	MAX_CONN = 30
	MAX_IDLE = 30
)

var (
	o orm.Ormer
)

func init() {
	orm.RegisterDriver("msyql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", sprintfDsn("inchat_user"), MAX_IDLE, MAX_CONN)
}

func sprintfDsn(dbName string) string {
	return fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=%s",
		beego.AppConfig.String("mysql_user"),
		beego.AppConfig.String("mysql_password"),
		beego.AppConfig.String("mysql_host"),
		beego.AppConfig.String("mysql_port"),
		dbName,
		beego.AppConfig.String("mysql_charset"))
}