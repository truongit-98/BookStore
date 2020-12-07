package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["BookStore/controllers/ordercontroller:OrderController"] = append(beego.GlobalControllerRouter["BookStore/controllers/ordercontroller:OrderController"],
        beego.ControllerComments{
            Method: "CheckoutOrder",
            Router: "/checkout",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
