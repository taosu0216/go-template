package user

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"go-template/service/logic/userLogic"
	"go-template/service/svc"
)

func GenerateUserRegisterLogic(r *svc.TemplateHandler) app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		l := userLogic.NewUserRegisterLogic(r)
		l.Register(ctx, c)
	}
}

func GenerateUserRegisterSendEmailLogic(r *svc.TemplateHandler) app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		l := userLogic.NewUserRegisterSendEmailLogic(r)
		l.SendRegisterEmail(ctx, c)
	}
}
