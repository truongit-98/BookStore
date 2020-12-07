package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["BookStore/controllers/authorcontrolcontroller:AuthorControlController"] = append(beego.GlobalControllerRouter["BookStore/controllers/authorcontrolcontroller:AuthorControlController"],
        beego.ControllerComments{
            Method: "Get",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["BookStore/controllers/authorcontrolcontroller:AuthorControlController"] = append(beego.GlobalControllerRouter["BookStore/controllers/authorcontrolcontroller:AuthorControlController"],
        beego.ControllerComments{
            Method: "CreateNewControl",
            Router: "/create",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["BookStore/controllers/authorcontrolcontroller:AuthorControlController"] = append(beego.GlobalControllerRouter["BookStore/controllers/authorcontrolcontroller:AuthorControlController"],
        beego.ControllerComments{
            Method: "RemoveControl",
            Router: "/delete/:controlId",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["BookStore/controllers/authorcontrolcontroller:AuthorControlController"] = append(beego.GlobalControllerRouter["BookStore/controllers/authorcontrolcontroller:AuthorControlController"],
        beego.ControllerComments{
            Method: "UpdateControl",
            Router: "/update",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
