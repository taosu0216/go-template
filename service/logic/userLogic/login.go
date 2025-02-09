package userLogic

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"go-template/data/cons"
	"go-template/service/svc"
	"go-template/service/types"
)

type UserLoginLogic struct {
	*svc.TemplateHandler
}

func NewUserLoginLogic(r *svc.TemplateHandler) *UserLoginLogic {
	return &UserLoginLogic{r}
}

func (u *UserLoginLogic) Login(ctx context.Context, c *app.RequestContext) {

	var req types.UserLoginReq
	if err := c.Bind(&req); err != nil {
		c.JSON(400, &types.UserLoginResp{
			BasicResp: types.BasicResp{
				Code: cons.BadRequestCode,
				Msg:  cons.BadRequestMsg,
				Err:  nil,
			},
		})
		return
	}
	resp, err := u.TemplateService.UserLoginInBiz(ctx, &req)
	if err != nil {
		u.ErrorfInService("UserLoginIn err is: %v", err)
	}
	c.JSON(200, resp)
}
