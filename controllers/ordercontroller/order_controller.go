package ordercontroller

import (
	"BookStore/requestbody"
	"BookStore/restapi/responses"
	"BookStore/services/orderservice"
	"BookStore/services/responseservice"
	"encoding/json"
	"github.com/astaxie/beego"
)

type OrderController struct {
	beego.Controller
}

// @Title OrderCheckout
// @Description checkout order
// @Param	token	header true	string	"token"
// @Param	body	body 	requestbody.OrderInformation	true	"OrderInfo"
// @Success 200 {object} responses.ResponseCommonSingle
// @router /checkout [post]
func (this *OrderController) CheckoutOrder() {
	defer this.ServeJSON()
	user := this.Ctx.Request.Header.Get("user")
	if user != "" {
		body := requestbody.OrderInformation{}
		err := json.Unmarshal(this.Ctx.Input.RequestBody, &body)
		if err != nil {
			this.Data["json"] = responseservice.GetCommonErrorResponse(responses.BadRequest)
			return
		}
		err = orderservice.CheckoutOrder(body, user)
		if err != nil {
			this.Data["json"] = responseservice.GetCommonErrorResponse(err)
			return
		}
		this.Data["json"] = responseservice.GetCommonSucceedResponse()
		return
	}
	this.Data["json"] = responseservice.GetCommonErrorResponse(responses.BadRequest)

}