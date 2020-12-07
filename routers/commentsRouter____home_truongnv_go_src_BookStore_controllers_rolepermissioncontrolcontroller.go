package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["BookStore/controllers/rolepermissioncontrolcontroller:RolePermsControlController"] = append(beego.GlobalControllerRouter["BookStore/controllers/rolepermissioncontrolcontroller:RolePermsControlController"],
        beego.ControllerComments{
            Method: "GetAllRolePermControls",
            Router: "/all",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
