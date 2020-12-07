package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["BookStore/controllers/rolepermissioncontroller:RolePermissionController"] = append(beego.GlobalControllerRouter["BookStore/controllers/rolepermissioncontroller:RolePermissionController"],
        beego.ControllerComments{
            Method: "GetRolePermissions",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["BookStore/controllers/rolepermissioncontroller:RolePermissionController"] = append(beego.GlobalControllerRouter["BookStore/controllers/rolepermissioncontroller:RolePermissionController"],
        beego.ControllerComments{
            Method: "AddMultiPermsToRole",
            Router: "/create/multi",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
