package biz

import (
	"context"
	"go-template/data/cons"
	"go-template/server/utils"
	"go-template/service/types"
)

func (uc *TemplateService) SendRegisterEmailInBiz(ctx context.Context, req *types.SendRegisterEmailReq) (*types.SendRegisterEmailResp, error) {
	isExist, err := uc.repo.IsUserExistInDataByEmail(ctx, req.Email)
	if err != nil {
		return &types.SendRegisterEmailResp{
			BasicResp: types.BasicResp{
				Code: cons.InternalServerError,
				Msg:  cons.InternalErrorMsg,
				Err:  cons.InternalErr,
			},
		}, err
	}
	if isExist {
		return &types.SendRegisterEmailResp{
			BasicResp: types.BasicResp{
				Code: cons.UserIsExistCode,
				Msg:  cons.UserIsExistMsg,
				Err:  nil,
			},
		}, nil
	}

	verifyCode := utils.GenerateRandomString()
	addr, host, passswd, from := uc.repo.GetMailInfo()
	err = utils.SendVerifyCodeToEmail(verifyCode, from, req.Email, addr, passswd, host, "注册验证码")
	if err != nil {
		return &types.SendRegisterEmailResp{
			BasicResp: types.BasicResp{
				Code: cons.InternalServerError,
				Msg:  cons.InternalErrorMsg,
				Err:  cons.InternalErr,
			},
		}, err
	}
	return uc.repo.SendRegisterEmailInData(ctx, req.Email, verifyCode)
}

func (uc *TemplateService) RegisterUserInBiz(ctx context.Context, req *types.RegisterUserReq) (*types.RegisterUserResp, error) {
	return uc.repo.RegisterUserInData(ctx, req)
}
