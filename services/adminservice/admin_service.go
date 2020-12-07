package adminservice

import (
	"BookStore/models"
	"BookStore/requestbody"
	"BookStore/restapi/responses"
	"errors"
	"github.com/prometheus/common/log"
)

func GetAll() ([]models.Admin, error){
	admin := models.Admin{}
	items, err := models.GetAll(&admin)
	if err != nil {
		log.Info(err.Error())
		return make([]models.Admin, 0), err
	}
	return items.([]models.Admin), nil
}
func GetPaginate(pos, count int32)([]models.Admin, int32, error)  {
	result := make([]models.Admin, 0)
	data, totalCount, err := models.GetPaginate(&models.Admin{}, pos, count)
	if err != nil {
		log.Info(err.Error())
		return result, totalCount,  responses.ErrSystem
	}
	result = data.([]models.Admin)
	return result, totalCount, nil
}

func RegisterAdminAccount(body requestbody.AccountRequest) error {
	conds := "email = ?"
	data, err := models.GetWithCondition(&models.Admin{}, conds, body.Email)
	if err != nil {
		log.Info(err.Error())
		return err
	}
	if len(data.([]models.Admin)) <= 0 {
		admin := models.Admin{
			Email: body.Email,
			Password: body.Password,
			FullName: body.FullName,
			Address: body.Address,
			Phone: body.Phone,
		}
		err := models.Create(&admin)
		if err != nil {
			log.Info(err.Error())
			return err
		}
		return nil
	}
	return errors.New("email is already taken")
}

func LoginAdmin(body requestbody.AccountLoginRequest) (int, error) {
	uid, err := (&models.Admin{}).Login(body.Email, body.Password)
	if err != nil {
		log.Info(err.Error())
		return 0, responses.ErrSystem
	}
	if uid <= 0 {
		return -1, responses.NotExisted
	}
	return uid, nil
}

func UpdateAdminAccount(body requestbody.AccountPutRequest) error {
	admin := models.Admin{ID: body.Id}
	if exist, err := models.Exists(&admin); err != nil {
		log.Info(err.Error())
		return err
	} else if exist {
		newAdmin := models.Admin{
			FullName: body.FullName,
			Address: body.Address,
			Phone: body.Phone,
		}
		err = models.Update(&newAdmin)
		if err != nil {
			log.Info(err.Error())
			return err
		}
		return nil
	}
	return responses.NotExisted
}

//func DeleteBook(id uint) error {
//	book := models.Book{ID: id}
//	if exist, err := models.Exists(book); err != nil {
//		log.Info(err.Error())
//		return err
//	} else if exist {
//		err = models.Remove(book)
//		if err != nil {
//			log.Info(err.Error())
//			return err
//		}
//		return nil
//	}
//	return responses.NotExisted
//}
