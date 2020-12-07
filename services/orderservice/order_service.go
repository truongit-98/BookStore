package orderservice

import (
	"BookStore/models"
	"BookStore/requestbody"
	"BookStore/restapi/data"
	"BookStore/restapi/responses"
	"BookStore/services/commonservice"
	"BookStore/services/customerservice"
	"BookStore/util"
	"encoding/json"
	"errors"
	"github.com/prometheus/common/log"
	"gorm.io/gorm"
	"strconv"
	"strings"
)

const (
	STATUS_PENDING = "STATUS_PENDING"
	STATUS_PROCESSING  = "STATUS_PROCESSING"
	STATUS_SUCCESS = "STATUS_SUCCESS"
)

var mapStatusOrder = map[string]string{
	STATUS_PENDING: "Đang chờ xử lý",
	STATUS_PROCESSING: "Đang xử lý",
	STATUS_SUCCESS: "Thành công",
}

func handleAddress(cityId, districtId, wardId string) (string, error) {
	citiesByte, err := commonservice.LoadCities()
	if err != nil {
		return "", err
	}
	var citiesMap map[string]data.Administrative
	err = json.Unmarshal(citiesByte, &citiesMap)
	if err != nil {
		log.Info(err)
		return "", responses.ErrSystem
	}
	if _, exist := citiesMap[cityId]; !exist {
		return "", responses.NotExisted
	}
	districtsByte, err := commonservice.LoadDistrictForCity(cityId)
	if err != nil {
		return "", err
	}
	var districtsMap map[string]data.Administrative
	err = json.Unmarshal(districtsByte, &districtsMap)
	if err != nil {
		log.Info(err)
		return "", responses.ErrSystem
	}
	if _, exist := districtsMap[districtId]; !exist {
		return "", responses.NotExisted
	}
	wardsByte, err := commonservice.LoadCities()
	if err != nil {
		return "", err
	}
	var wardsMap map[string]data.Administrative
	err = json.Unmarshal(wardsByte, &wardsMap)
	if err != nil {
		log.Info(err)
		return "", responses.ErrSystem
	}
	if ward, exist := citiesMap[cityId]; !exist {
		return "", responses.NotExisted
	} else {
		return ward.PathWithType, nil
	}
}

func CheckoutOrder(req requestbody.OrderInformation, user string) error {
	userId, err := strconv.Atoi(user)
	if err != nil {
		log.Info(err.Error())
		return responses.BadRequest
	}
	userInfo, err := customerservice.GetCustomerInfo(userId)
	if err != nil {
		return err
	}
	var receiveInfo models.ReceiveInfo
	if req.ReceiveInfo != nil {
		addressDetail := req.ReceiveInfo.AddressDetail
		tempAddress, err := handleAddress(req.ReceiveInfo.CityId, req.ReceiveInfo.DistrictId, req.ReceiveInfo.WardId)
		if err != nil {
			return err
		}
		address := strings.Join([]string{addressDetail, tempAddress}, ", " )

		receiveInfo.FullName = req.ReceiveInfo.FullName
		receiveInfo.Email = req.ReceiveInfo.Email
		receiveInfo.Phone = req.ReceiveInfo.Phone
		receiveInfo.Address = address
	} else {
		receiveInfo.FullName = userInfo.FullName
		receiveInfo.Email = userInfo.Email
		receiveInfo.Phone = userInfo.Phone
		receiveInfo.Address = userInfo.Address
	}
	now := util.Now()
	bytes, _ := json.Marshal(receiveInfo)
	receiveInfoStr := string(bytes)
	// begin transaction
	err = models.Transaction(func(tx *gorm.DB) error {
		defer func() {
			if r := recover(); r != nil {
				tx.Rollback()
			}
		}()
		//check coupon
		couponId, err := checkCoupon(now, req.Coupon)
		if couponId <= 0{
			return errors.New("coupon is not valid")
		}
		orderIns := models.Order{
			Status: mapStatusOrder[STATUS_PENDING],
			OrderDate: now,
			ReceiveInfo: &receiveInfoStr,
			CustomerID: uint(userId),
			PaymentID: uint(req.PaymentMethod),
		}
		if err = tx.Create(&orderIns).Error; err != nil {
			return err
		}
		if orderIns.ID == 0{
			return errors.New("create order error")
		}

		// create order detail
		for _, order := range req.Orders{
			orderDetail := models.OrderDetail{
				OrderID: orderIns.ID,
				BookID: uint(order.ProductID),
				Quantity: order.Amount,
				Price: order.Price,
			}
			if err = tx.Create(&orderDetail).Error; err != nil {
				return err
			}
		}

		//create order voucher
		orderVoucher := models.OrderVouchers{
			OrderID: orderIns.ID,
			VoucherID: couponId,
		}
		if err = tx.Create(&orderVoucher).Error; err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		log.Info(err.Error())
		return responses.ErrSystem
	}
	return nil
}

func checkCoupon(checkTime int64, code string) (id uint, err error) {
	coupon, err := (&models.Voucher{}).GetByCode(code)
	if err != nil  {
		log.Info(err)
		if errors.Is(err, gorm.ErrRecordNotFound){
			return 0, errors.New("coupon is not existed")
		}
		return 0, responses.ErrSystem
	}
	return coupon.(*models.Voucher).ID, err
}
