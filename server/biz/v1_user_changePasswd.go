package biz

import (
	"context"
	"encoding/json"
	"go-template/data/cons"
	"go-template/server/utils"
	"go-template/service/types"
)

func (uc *TemplateService) ChangePasswordInBiz(ctx context.Context, req *types.ChangePasswordReq) (*types.ChangePasswordResp, error) {
	return uc.repo.ChangePasswordInData(ctx, req)
}

func (uc *TemplateService) SendChangePasswordEmailInBiz(ctx context.Context, req *types.SendChangePasswdEmailReq) (*types.SendChangePasswdEmailResp, error) {
	code, err := uc.repo.IsUserExistInDataByEmail(ctx, req.Email)
	if err != nil {
		return &types.SendChangePasswdEmailResp{
			BasicResp: types.BasicResp{
				Code: cons.InternalServerError,
				Msg:  cons.InternalErrorMsg,
				Err:  cons.InternalErr,
			},
		}, err
	}
	if code == cons.UserStatusNotRegister || code == cons.UserStatusWaitToRegister {
		return &types.SendChangePasswdEmailResp{
			BasicResp: types.BasicResp{
				Code: cons.UserNotFoundCode,
				Msg:  cons.UserNotFoundMsg,
				Err:  nil,
			},
		}, nil
	}

	verifyCode := utils.GenerateRandomString()
	message := &cons.SendEmailModel{
		Email:      req.Email,
		VerifyCode: verifyCode,
		Subject:    "修改密码-验证码",
	}
	messageBytes, err := json.Marshal(message)
	if err != nil {
		uc.ErrorfInBiz("SendChangePasswordEmailInBiz [Marshal] err is: %v", err)
		return nil, err
	}

	// 发送消息到 Kafka
	go func() {
		err = uc.repo.SendKafkaMessage(cons.EmailPasswordResetNotificationTopic, messageBytes)
		if err != nil {
			uc.ErrorfInBiz("SendChangePasswordEmailInBiz [WriteMessages] err is: %v", err)
		}
	}()

	return uc.repo.SendChangePasswordEmailInData(ctx, req.Email, verifyCode)
}
