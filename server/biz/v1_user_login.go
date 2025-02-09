package biz

import (
	"context"
	"go-template/data/cons"
	"go-template/service/types"
	"time"

	"github.com/golang-jwt/jwt"
	jwtv4 "github.com/golang-jwt/jwt/v4"
)

var jwtKey = ""

type MyClaims struct {
	UserName string `json:"username"`
	jwtv4.StandardClaims
}

func (uc *TemplateService) UserLoginInBiz(ctx context.Context, req *types.UserLoginReq) (*types.UserLoginResp, error) {
	jwtKey = uc.repo.GetJwtKey()
	ok, err := uc.repo.IsPasswordRightInData(ctx, req.UserName, req.PassWd)
	if err != nil {
		return &types.UserLoginResp{
			BasicResp: types.BasicResp{
				Code: cons.InternalServerError,
				Msg:  cons.InternalErrorMsg,
				Err:  cons.InternalErr,
			},
		}, err
	}
	if !ok {
		return &types.UserLoginResp{
			BasicResp: types.BasicResp{
				Code: cons.UserPasswordNotRightCode,
				Msg:  cons.UserPasswordNotRightMsg,
				Err:  nil,
			},
		}, nil
	}

	myClaims := MyClaims{
		req.UserName,
		jwtv4.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 8).Unix(), //设置JWT过期时间,此处设置为2小时
			Issuer:    "taosu",                              //设置签发人
		},
	}
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, myClaims)
	token, err := claims.SignedString([]byte(jwtKey))
	if err != nil {
		uc.ErrorfInBiz("UserLoginInBiz [SignedString] err is: %v", err)
		return &types.UserLoginResp{
			BasicResp: types.BasicResp{
				Code: cons.InternalServerError,
				Msg:  cons.InternalErrorMsg,
				Err:  cons.InternalErr,
			},
		}, err
	}
	return &types.UserLoginResp{
		BasicResp: types.BasicResp{
			Code: cons.BasicSuccessCode,
			Msg:  cons.BasicSuccessMsg,
			Err:  nil,
		},
		Token: token,
	}, nil

}
