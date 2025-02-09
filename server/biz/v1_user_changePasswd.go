package biz

import (
	"context"
	"go-template/data/cons"
	"go-template/server/utils"
	"go-template/service/types"
)

func (uc *TemplateService) ChangePasswordInBiz(ctx context.Context, req *types.ChangePasswordReq) (*types.ChangePasswordResp, error) {
	return uc.repo.ChangePasswordInData(ctx, req)
}

func (uc *TemplateService) SendChangePasswordEmailInBiz(ctx context.Context, req *types.SendChangePasswdEmailReq) (*types.SendChangePasswdEmailResp, error) {
	isExist, err := uc.repo.IsUserExistInDataByEmail(ctx, req.Email)
	if err != nil {
		return &types.SendChangePasswdEmailResp{
			BasicResp: types.BasicResp{
				Code: cons.InternalServerError,
				Msg:  cons.InternalErrorMsg,
				Err:  cons.InternalErr,
			},
		}, err
	}
	if !isExist {
		return &types.SendChangePasswdEmailResp{
			BasicResp: types.BasicResp{
				Code: cons.UserNotFoundCode,
				Msg:  cons.UserNotFoundMsg,
				Err:  nil,
			},
		}, nil
	}

	verifyCode := utils.GenerateRandomString()
	addr, host, passswd, from := uc.repo.GetMailInfo()
	err = utils.SendVerifyCodeToEmail(verifyCode, from, req.Email, addr, passswd, host, "修改密码验证码")
	if err != nil {
		return &types.SendChangePasswdEmailResp{
			BasicResp: types.BasicResp{
				Code: cons.InternalServerError,
				Msg:  cons.InternalErrorMsg,
				Err:  cons.InternalErr,
			},
		}, err
	}
	return uc.repo.SendChangePasswordEmailInData(ctx, req.Email, verifyCode)
}
