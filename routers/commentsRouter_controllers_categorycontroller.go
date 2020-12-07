package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["BookStore/controllers/categorycontroller:CategoryController"] = append(beego.GlobalControllerRouter["BookStore/controllers/categorycontroller:CategoryController"],
        beego.ControllerComments{
            Method: "Get",
            Router: "/with-book-count",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
