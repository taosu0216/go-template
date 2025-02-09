package userLogic

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"go-template/data/cons"
	"go-template/service/svc"
	"go-template/service/types"
	"time"
)

type ChangePasswdLogic struct {
	*svc.TemplateHandler
}

func NewChangePasswdLogic(r *svc.TemplateHandler) *ChangePasswdLogic {
	return &ChangePasswdLogic{r}
}

func (u *ChangePasswdLogic) ChangePasswd(ctx context.Context, c *app.RequestContext) {
	now := time.Now()
	defer u.LogDuration("ChangePasswd", now)

	var req types.ChangePasswordReq
	if err := c.Bind(&req); err != nil {
		c.JSON(400, &types.ChangePasswordResp{
			BasicResp: types.BasicResp{
				Code: cons.BadRequestCode,
				Msg:  cons.BadRequestMsg,
				Err:  nil,
			},
		})
		return
	}
	resp, err := u.TemplateService.ChangePasswordInBiz(ctx, &req)
	if err != nil {
		u.ErrorfInService("ChangePasswordInBiz err is: %v", err)
	}
	c.JSON(200, resp)
}

type ChangePasswdSendEmailLogic struct {
	*svc.TemplateHandler
}

func NewChangePasswdSendEmailLogic(r *svc.TemplateHandler) *ChangePasswdSendEmailLogic {
	return &ChangePasswdSendEmailLogic{r}
}

func (u *ChangePasswdSendEmailLogic) SendChangePasswdEmail(ctx context.Context, c *app.RequestContext) {
	now := time.Now()
	defer u.LogDuration("SendChangePasswdEmail", now)

	var req types.SendChangePasswdEmailReq
	if err := c.Bind(&req); err != nil {
		c.JSON(400, &types.SendChangePasswdEmailResp{
			BasicResp: types.BasicResp{
				Code: cons.BadRequestCode,
				Msg:  cons.BadRequestMsg,
				Err:  nil,
			},
		})
		return
	}
	resp, err := u.TemplateService.SendChangePasswordEmailInBiz(ctx, &req)
	if err != nil {
		u.ErrorfInService("SendRegisterEmailInBiz err is: %v", err)
	}
	c.JSON(200, resp)
}
