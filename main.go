package main

import (
	_ "Polux_API/routers"
	"github.com/astaxie/beego/plugins/cors"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
)

func init() {
	beego.Debug("Filters init...")
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		 AllowAllOrigins:  true,
		 AllowMethods: 	[]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		 AllowHeaders: 	[]string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Content-Type"},
		 ExposeHeaders:	[]string{"Content-Length", "Access-Control-Allow-Origin"},
		 AllowCredentials: true,
	}))
	orm.RegisterDataBase("default", "postgres", "postgres://polux_admin:PolLu10Adm$2016@10.20.0.62:5432/udistrital?sslmode=disable&search_path=polux")
}

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	beego.Run()
}
