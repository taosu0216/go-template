package user

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"go-template/service/logic/userLogic"
	"go-template/service/svc"
)

func GenerateUserLoginLogic(r *svc.TemplateHandler) app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		l := userLogic.NewUserLoginLogic(r)
		l.Login(ctx, c)
	}
}
