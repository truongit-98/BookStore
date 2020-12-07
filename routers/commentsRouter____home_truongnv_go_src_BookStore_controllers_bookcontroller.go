package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["BookStore/controllers/bookcontroller:BookController"] = append(beego.GlobalControllerRouter["BookStore/controllers/bookcontroller:BookController"],
        beego.ControllerComments{
            Method: "GetFeatured",
            Router: "/featured",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["BookStore/controllers/bookcontroller:BookController"] = append(beego.GlobalControllerRouter["BookStore/controllers/bookcontroller:BookController"],
        beego.ControllerComments{
            Method: "GetWithFilter",
            Router: "/filter",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["BookStore/controllers/bookcontroller:BookController"] = append(beego.GlobalControllerRouter["BookStore/controllers/bookcontroller:BookController"],
        beego.ControllerComments{
            Method: "GetInfo",
            Router: "/info/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["BookStore/controllers/bookcontroller:BookController"] = append(beego.GlobalControllerRouter["BookStore/controllers/bookcontroller:BookController"],
        beego.ControllerComments{
            Method: "GetNews",
            Router: "/new",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["BookStore/controllers/bookcontroller:BookController"] = append(beego.GlobalControllerRouter["BookStore/controllers/bookcontroller:BookController"],
        beego.ControllerComments{
            Method: "GetBestSellers",
            Router: "/seller",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
