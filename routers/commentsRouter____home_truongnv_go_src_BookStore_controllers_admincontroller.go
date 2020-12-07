package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["BookStore/controllers/admincontroller:AdminController"] = append(beego.GlobalControllerRouter["BookStore/controllers/admincontroller:AdminController"],
        beego.ControllerComments{
            Method: "Get",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["BookStore/controllers/admincontroller:AdminController"] = append(beego.GlobalControllerRouter["BookStore/controllers/admincontroller:AdminController"],
        beego.ControllerComments{
            Method: "LoginAdmin",
            Router: "/login",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["BookStore/controllers/admincontroller:AdminController"] = append(beego.GlobalControllerRouter["BookStore/controllers/admincontroller:AdminController"],
        beego.ControllerComments{
            Method: "RegisterAdmin",
            Router: "/register",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
