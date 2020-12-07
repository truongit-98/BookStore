package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["BookStore/controllers/rolecontroller:RoleController"] = append(beego.GlobalControllerRouter["BookStore/controllers/rolecontroller:RoleController"],
        beego.ControllerComments{
            Method: "Get",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["BookStore/controllers/rolecontroller:RoleController"] = append(beego.GlobalControllerRouter["BookStore/controllers/rolecontroller:RoleController"],
        beego.ControllerComments{
            Method: "CreateRole",
            Router: "/create",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["BookStore/controllers/rolecontroller:RoleController"] = append(beego.GlobalControllerRouter["BookStore/controllers/rolecontroller:RoleController"],
        beego.ControllerComments{
            Method: "Removerole",
            Router: "/delete/:roleId",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["BookStore/controllers/rolecontroller:RoleController"] = append(beego.GlobalControllerRouter["BookStore/controllers/rolecontroller:RoleController"],
        beego.ControllerComments{
            Method: "GetRolesById",
            Router: "/detail/:roleId",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["BookStore/controllers/rolecontroller:RoleController"] = append(beego.GlobalControllerRouter["BookStore/controllers/rolecontroller:RoleController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/update",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["BookStore/controllers/rolecontroller:RoleController"] = append(beego.GlobalControllerRouter["BookStore/controllers/rolecontroller:RoleController"],
        beego.ControllerComments{
            Method: "GetRolesForUser",
            Router: "/user/:userId",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
