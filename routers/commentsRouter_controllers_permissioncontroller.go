package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["BookStore/controllers/permissioncontroller:PermissionController"] = append(beego.GlobalControllerRouter["BookStore/controllers/permissioncontroller:PermissionController"],
        beego.ControllerComments{
            Method: "Get",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["BookStore/controllers/permissioncontroller:PermissionController"] = append(beego.GlobalControllerRouter["BookStore/controllers/permissioncontroller:PermissionController"],
        beego.ControllerComments{
            Method: "CreatePermission",
            Router: "/create",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["BookStore/controllers/permissioncontroller:PermissionController"] = append(beego.GlobalControllerRouter["BookStore/controllers/permissioncontroller:PermissionController"],
        beego.ControllerComments{
            Method: "RemovePermission",
            Router: "/delete/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["BookStore/controllers/permissioncontroller:PermissionController"] = append(beego.GlobalControllerRouter["BookStore/controllers/permissioncontroller:PermissionController"],
        beego.ControllerComments{
            Method: "GetPermissionsForRole",
            Router: "/role/:roleId",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["BookStore/controllers/permissioncontroller:PermissionController"] = append(beego.GlobalControllerRouter["BookStore/controllers/permissioncontroller:PermissionController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/update",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
