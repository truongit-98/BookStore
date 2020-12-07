package bookcontroller

import (
	"BookStore/restapi/responses"
	"BookStore/services/bookservice"
	"github.com/astaxie/beego"
	"BookStore/services/responseservice"

)

type BookController struct{
	beego.Controller
}



// @Title GetInfoBook
// @Description get info book of id
// @Param	id	path 	int32	true	"ProductID"
// @Success 200 {object} responses.ResponseCommonSingle
// @router /info/:id [get]
func (this *BookController) GetInfo(){
	defer this.ServeJSON()

	id, err := this.GetInt32(":id")
	if err != nil {
		this.Data["json"] = responseservice.GetCommonErrorResponse(responses.BadRequest)
		return 
	}
	data, err := bookservice.GetBookInfo(uint(id))
	if err != nil {
		this.Data["json"] = responseservice.GetCommonErrorResponse(err)
		return 
	}
	this.Data["json"] = responseservice.GetCommonSucceedResponseWithData(data)

}
// @Title GetFeatureadProducts
// @Description get featured producs
// @Success 200 {object} responses.ResponseCommonArray
// @router /featured [get]
func (this *BookController) GetFeatured(){
	defer this.ServeJSON()
	data, err := bookservice.GetFeaturedProducts()
	if err != nil {
		this.Data["json"] = responseservice.GetCommonErrorResponseArray(err)
		return 
	}
	this.Data["json"] = responseservice.GetCommonSucceedResponseWithData(data)

}

// @Title GetNewProducts
// @Description get new products
// @Success 200 {object} responses.ResponseCommonArray
// @router /new [get]
func (this *BookController) GetNews(){
	defer this.ServeJSON()
	data, err := bookservice.GetNewProducts()
	if err != nil {
		this.Data["json"] = responseservice.GetCommonErrorResponseArray(err)
		return 
	}
	this.Data["json"] = responseservice.GetCommonSucceedResponseWithData(data)

}

// @Title GetBestSellerProducts
// @Description get best seller products
// @Success 200 {object} responses.ResponseCommonArray
// @router /seller [get]
func (this *BookController) GetBestSellers(){
	defer this.ServeJSON()
	data, err := bookservice.GetBestSellers()
	if err != nil {
		this.Data["json"] = responseservice.GetCommonErrorResponseArray(err)
		return 
	}
	this.Data["json"] = responseservice.GetCommonSucceedResponseWithData(data)

}

// @Title GetProductsWithFilter
// @Description get products with filter
// @Param	category	query 	int32	false 	"Category"
// @Param	topic		query	int32	false	"Topic"
// @Param	priceMin		query	float64	false	"PriceMin"
// @Param	priceMax		query	float64	false	"PriceMax"
// @Param	manufacturer	query	int32	false	"Manufacturers"
// @Param	sort	query	string	false	"Sorted"
// @Param	offset	query	int32	false	"Offset"
// @Param	limit	query	int32	false	"Limit"
// @Success 200 {object} responses.ResponseCommonArray
// @router /filter [get]
func (this *BookController) GetWithFilter(){
	defer this.ServeJSON()
	categoryId, err := this.GetInt32("category", 0)
	if err != nil {
		this.Data["json"] = responseservice.GetCommonErrorResponseArray(responses.BadRequest)
		return
	}
	topicId, err := this.GetInt32("topic", 0)
	if err != nil {
		this.Data["json"] = responseservice.GetCommonErrorResponseArray(responses.BadRequest)
		return
	}
	priceMin, err := this.GetFloat("priceMin", 0)
	if err != nil {
		this.Data["json"] = responseservice.GetCommonErrorResponseArray(responses.BadRequest)
		return
	}
	priceMax, err := this.GetFloat("priceMax", 0)
	if err != nil {
		this.Data["json"] = responseservice.GetCommonErrorResponseArray(responses.BadRequest)
		return
	}
	manufacturerId, err := this.GetInt32("manufacturer", 0)
	if err != nil {
		this.Data["json"] = responseservice.GetCommonErrorResponseArray(responses.BadRequest)
		return
	}
	sort := this.GetString("sort")
	offset, err := this.GetInt32("offset", 0)
	if err != nil {
		this.Data["json"] = responseservice.GetCommonErrorResponseArray(responses.BadRequest)
		return
	}
	limit, err := this.GetInt32("limit", 0)
	if err != nil {
		this.Data["json"] = responseservice.GetCommonErrorResponseArray(responses.BadRequest)
		return
	}
	data, err := bookservice.GetWithFilter(categoryId, topicId, manufacturerId, priceMin, priceMax, sort, offset, limit)
	if err != nil {
		this.Data["json"] = responseservice.GetCommonErrorResponseArray(err)
		return
	}
	this.Data["json"] = responseservice.GetCommonSucceedResponseWithData(data)

}