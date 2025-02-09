package user

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"go-template/service/logic/userLogic"
	"go-template/service/svc"
)

func GenerateUserChangePasswdLogic(r *svc.TemplateHandler) app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		l := userLogic.NewChangePasswdLogic(r)
		l.ChangePasswd(ctx, c)
	}
}

func GenerateUserChangePasswdSendEmailLogic(r *svc.TemplateHandler) app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		l := userLogic.NewChangePasswdSendEmailLogic(r)
		l.SendChangePasswdEmail(ctx, c)
	}
}
