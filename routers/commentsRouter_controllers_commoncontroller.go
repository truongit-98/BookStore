package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["BookStore/controllers/commoncontroller:CommonController"] = append(beego.GlobalControllerRouter["BookStore/controllers/commoncontroller:CommonController"],
        beego.ControllerComments{
            Method: "GetDistrict",
            Router: "/administrative/districts/:cityId",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["BookStore/controllers/commoncontroller:CommonController"] = append(beego.GlobalControllerRouter["BookStore/controllers/commoncontroller:CommonController"],
        beego.ControllerComments{
            Method: "GetWards",
            Router: "/administrative/wards/:districtId",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
