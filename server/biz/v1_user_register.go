package biz

import (
	"context"
	"encoding/json"
	"go-template/data/cons"
	"go-template/server/utils"
	"go-template/service/types"
)

func (uc *TemplateService) SendRegisterEmailInBiz(ctx context.Context, req *types.SendRegisterEmailReq) (*types.SendRegisterEmailResp, error) {
	code, err := uc.repo.IsUserExistInDataByEmail(ctx, req.Email)
	if err != nil {
		return &types.SendRegisterEmailResp{
			BasicResp: types.BasicResp{
				Code: cons.InternalServerError,
				Msg:  cons.InternalErrorMsg,
				Err:  cons.InternalErr,
			},
		}, err
	}
	if code == cons.UserStatusRegistered {
		return &types.SendRegisterEmailResp{
			BasicResp: types.BasicResp{
				Code: cons.UserIsExistCode,
				Msg:  cons.UserIsExistMsg,
				Err:  nil,
			},
		}, nil
	} else if code == cons.UserStatusWaitToRegister {
		return &types.SendRegisterEmailResp{
			BasicResp: types.BasicResp{
				Code: cons.UserWaitToRegisterCode,
				Msg:  cons.UserWaitToRegisterMsg,
				Err:  nil,
			},
		}, nil
	}

	verifyCode := utils.GenerateRandomString()

	// 创建要发送的消息
	message := &cons.SendEmailModel{
		Email:      req.Email,
		VerifyCode: verifyCode,
		Subject:    "注册验证码",
	}
	messageBytes, err := json.Marshal(message)
	if err != nil {
		uc.ErrorfInBiz("SendRegisterEmailInBiz [Marshal] err is: %v", err)
		return nil, err
	}

	// 发送消息到 Kafka
	go func() {
		err = uc.repo.SendKafkaMessage(cons.EmailRegisterNotificationTopic, messageBytes)
		if err != nil {
			uc.ErrorfInBiz("SendRegisterEmailInBiz [WriteMessages] err is: %v", err)
		}
	}()

	return uc.repo.SendRegisterEmailInData(ctx, req.Email, verifyCode)
}

func (uc *TemplateService) RegisterUserInBiz(ctx context.Context, req *types.RegisterUserReq) (*types.RegisterUserResp, error) {
	return uc.repo.RegisterUserInData(ctx, req)
}
