package main

import (

	"BookStore/models"


	_ "BookStore/routers"
	"github.com/astaxie/beego/plugins/cors"
	"log"

	"github.com/astaxie/beego"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "x-access-token", "content-type", "Content-Type", "sessionkey", "token"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Content-Type"},
		AllowCredentials: true,
	}))
	_, err := models.InitDB()
	if err != nil {
		log.Println(err.Error())
	}

	//beego.InsertFilter("*", beego.BeforeExec, middleware.CheckSession())
	beego.Run()
}
