package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["BookStore/controllers/customercontroller:CustomerController"] = append(beego.GlobalControllerRouter["BookStore/controllers/customercontroller:CustomerController"],
        beego.ControllerComments{
            Method: "Get",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["BookStore/controllers/customercontroller:CustomerController"] = append(beego.GlobalControllerRouter["BookStore/controllers/customercontroller:CustomerController"],
        beego.ControllerComments{
            Method: "GetInfo",
            Router: "/detail/:userId",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["BookStore/controllers/customercontroller:CustomerController"] = append(beego.GlobalControllerRouter["BookStore/controllers/customercontroller:CustomerController"],
        beego.ControllerComments{
            Method: "LoginAdmin",
            Router: "/login",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["BookStore/controllers/customercontroller:CustomerController"] = append(beego.GlobalControllerRouter["BookStore/controllers/customercontroller:CustomerController"],
        beego.ControllerComments{
            Method: "RegisterAdmin",
            Router: "/register",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
