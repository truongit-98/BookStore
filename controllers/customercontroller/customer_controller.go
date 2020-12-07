package customercontroller

import (
	"BookStore/requestbody"
	"BookStore/restapi/responses"
	"BookStore/services/customerservice"
	"BookStore/services/responseservice"
	"BookStore/util"
	"encoding/json"
	"github.com/astaxie/beego"
)

type CustomerController struct {
	beego.Controller
}

// @Title RegisterAdmin
// @Description register admin
// @Param	token	header true	string	"token"
// @Param	admin	body 	requestbody.AccountRequest	true	"Admin"
// @Success 200 {object} responses.ResponseCommonSingle
// @router /register [post]
func (this *CustomerController) RegisterAdmin() {
	defer this.ServeJSON()
	var err error
	body := requestbody.AccountRequest{}
	err = json.Unmarshal(this.Ctx.Input.RequestBody, &body)
	if err != nil {
		this.Data["json"] = responseservice.GetCommonErrorResponse(responses.BadRequest)
		return
	}
	err = customerservice.RegisterUserAccount(body)
	if err != nil {
		if _, ok := responses.MapDescription[err]; ok {
			this.Data["json"] = responseservice.GetCommonErrorResponse(err)
		} else {
			this.Data["json"] = responses.ResponseCommonSingle{
				Data: map[string]string{
					"message": err.Error(),
				},
				Error: responses.NewErr(responses.UnSuccess),
			}
		}
		return
	}
	this.Data["json"] = responseservice.GetCommonSucceedResponse()
}

// @Title RegisterAdmin
// @Description register admin
// @Param	token	header true	string	"token"
// @Param	admin	body 	requestbody.AccountLoginRequest	true	"Admin"
// @Success 200 {object} responses.ResponseCommonSingle
// @router /login [post]
func (this *CustomerController) LoginAdmin() {
	defer this.ServeJSON()
	var err error
	body := requestbody.AccountLoginRequest{}
	err = json.Unmarshal(this.Ctx.Input.RequestBody, &body)
	if err != nil {
		this.Data["json"] = responseservice.GetCommonErrorResponse(responses.BadRequest)
		return
	}
	userId, err := customerservice.LoginUser(body)
	if err != nil {
		if _, ok := responses.MapDescription[err]; ok {
			this.Data["json"] = responseservice.GetCommonErrorResponse(err)
		} else {
			this.Data["json"] = responses.ResponseCommonSingle{
				Data: map[string]string{
					"message": err.Error(),
				},
				Error: responses.NewErr(responses.UnSuccess),
			}
		}
		return
	}
	token, err := util.CreateToken(userId, util.ACCOUNT_USER)
	if err != nil {
		this.Data["json"] = responseservice.GetCommonErrorResponse(responses.ErrUnknown)
		return
	}
	// err = redisservice.CreateAuth(userId, token)
	// if err != nil {
	// 	this.Data["json"] = responseservice.GetCommonErrorResponse(responses.ErrUnknown)
	// 	return
	// }
	tokens := map[string]string{
		"access_token":  token.AccessToken,
		"refresh_token": token.RefreshToken,
	}
	this.Data["json"] = responseservice.GetCommonSucceedResponseWithData(tokens)
}

// @Title GetAdmins
// @Description get roles
// @Param	token	header 	string true	"token"
// @Param	pos	query	int64	false	"Position"
// @Param	count	query	int64	false	"Count"
// @Success 200 {object} responses.ResponseCommonArray
// @router / [get]
func (this *CustomerController) Get() {
	defer this.ServeJSON()
	pos, err := this.GetInt32("pos", 0)
	if err != nil {
		this.Data["json"] = responseservice.GetCommonErrorResponseArray(responses.BadRequest)
		return
	}
	count, err := this.GetInt32("count",0)
	if err != nil {
		this.Data["json"] = responseservice.GetCommonErrorResponseArray(responses.BadRequest)
		return
	}
	if pos == 0 && count == 0 {
		this.GetAll()
	} else {
		this.GetPaginate(pos, count)
	}
}

func (this *CustomerController) GetAll() {
	roles, err := customerservice.GetAll()
	if err != nil {
		this.Data["json"] = responseservice.GetCommonErrorResponseArray(err)
		return
	}
	this.Data["json"] = responseservice.GetCommonSucceedResponseWithData(roles)
}

func (this *CustomerController) GetPaginate(pos, count int32) {
	var err error
	roles, totalCount, err := customerservice.GetPaginate(pos, count)
	if err != nil {
		this.Data["json"] = responseservice.GetCommonErrorResponseArray(err)
		return
	}
	this.Data["json"] = responseservice.GetCommonSucceedResponseArray(roles, totalCount)
}
