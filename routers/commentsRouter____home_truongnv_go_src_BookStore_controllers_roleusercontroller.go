package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["BookStore/controllers/roleusercontroller:RoleUserController"] = append(beego.GlobalControllerRouter["BookStore/controllers/roleusercontroller:RoleUserController"],
        beego.ControllerComments{
            Method: "GetRoleUsers",
            Router: "/all",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["BookStore/controllers/roleusercontroller:RoleUserController"] = append(beego.GlobalControllerRouter["BookStore/controllers/roleusercontroller:RoleUserController"],
        beego.ControllerComments{
            Method: "AddMultiRoleForUser",
            Router: "/multi",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["BookStore/controllers/roleusercontroller:RoleUserController"] = append(beego.GlobalControllerRouter["BookStore/controllers/roleusercontroller:RoleUserController"],
        beego.ControllerComments{
            Method: "GetRoleUsersForUser",
            Router: "/user/:userId",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
