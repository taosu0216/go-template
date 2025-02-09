package userLogic

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"go-template/data/cons"
	"go-template/service/svc"
	"go-template/service/types"
	"time"
)

type UserRegisterLogic struct {
	*svc.TemplateHandler
}

func NewUserRegisterLogic(r *svc.TemplateHandler) *UserRegisterLogic {
	return &UserRegisterLogic{r}
}

func (u *UserRegisterLogic) Register(ctx context.Context, c *app.RequestContext) {
	now := time.Now()
	defer u.LogDuration("RegisterUserInService", now)

	var req types.RegisterUserReq
	if err := c.Bind(&req); err != nil {
		c.JSON(400, &types.RegisterUserResp{
			BasicResp: types.BasicResp{
				Code: cons.BadRequestCode,
				Msg:  cons.BadRequestMsg,
				Err:  nil,
			},
		})
		return
	}
	resp, err := u.TemplateService.RegisterUserInBiz(ctx, &req)
	if err != nil {
		u.ErrorfInService("RegisterUserInBiz err is: %v", err)
	}
	c.JSON(200, resp)
}

type UserRegisterSendEmailLogic struct {
	*svc.TemplateHandler
}

func NewUserRegisterSendEmailLogic(r *svc.TemplateHandler) *UserRegisterSendEmailLogic {
	return &UserRegisterSendEmailLogic{r}
}

func (u *UserRegisterSendEmailLogic) SendRegisterEmail(ctx context.Context, c *app.RequestContext) {
	now := time.Now()
	defer u.LogDuration("SendRegisterEmailInService", now)

	var req types.SendRegisterEmailReq
	if err := c.Bind(&req); err != nil {
		c.JSON(400, &types.SendRegisterEmailResp{
			BasicResp: types.BasicResp{
				Code: cons.BadRequestCode,
				Msg:  cons.BadRequestMsg,
				Err:  nil,
			},
		})
		return
	}
	resp, err := u.TemplateService.SendRegisterEmailInBiz(ctx, &req)
	if err != nil {
		u.ErrorfInService("SendRegisterEmailInBiz err is: %v", err)
	}
	c.JSON(200, resp)
}
