package middleware

import (
	"BookStore/common/structs"
	"BookStore/consts"
	"BookStore/restapi/responses"
	"BookStore/services/redisservice"
	"BookStore/services/responseservice"
	"errors"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/dgrijalva/jwt-go"
	"github.com/prometheus/common/log"
	"net/http"
	"strings"
)
func CheckSession() beego.FilterFunc {
	return func(ctx *context.Context) {
		if !strings.Contains(ctx.Request.URL.String(), "/v1/api/account/admins/") && !strings.Contains(ctx.Request.URL.String(), "/v1/api/account/admins/login") && !strings.Contains(ctx.Request.URL.String(), "/v1/api/account/admins/register") {
			tokenAuth, err := ExtractTokenMetadata(ctx.Request)
			if err != nil {
				log.Info(err.Error())
				ctx.Output.JSON(responseservice.GetCommonErrorResponse(responses.ErrUnknown), true, true)
				ctx.ResponseWriter.WriteHeader(http.StatusOK)
				return
			}
			user, err := redisservice.FetchAuth(tokenAuth.AccessUuid)
			if err != nil {
				log.Info(err.Error())
				ctx.Output.JSON(responseservice.GetCommonErrorResponse(responses.UnAuthorized), true, true)
				ctx.ResponseWriter.WriteHeader(http.StatusOK)
				return
			}
			if *user != tokenAuth.User {
				ctx.Output.JSON(responseservice.GetCommonErrorResponse(responses.UnAuthorized), true, true)
				ctx.ResponseWriter.WriteHeader(http.StatusOK)
				return
			}
			ctx.Input.SetParam("user", *user)
			return
		}
		return
	}
}

func ExtractTokenMetadata(r *http.Request) (*structs.AccessDetails, error) {
	token, err := verifyToken(r)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		accessUuid, ok := claims["access_uuid"].(string)
		if !ok {
			return nil, errors.New("token is not valid")
		}
		user, ok := claims["user"].(string)
		if !ok {
			return nil, errors.New("token is not valid")
		}
		return &structs.AccessDetails{
			AccessUuid: accessUuid,
			User:   user,
		}, nil
	}
	return nil, err
}
func verifyToken(r *http.Request) (*jwt.Token, error) {
	tokenString := r.Header.Get("token")
	log.Info(tokenString)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(consts.ACCESS_SECRET), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}
