package auth

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
)

func GeneratePingLogic() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		usernameValue, exists := c.Get("username")
		if !exists {
			c.JSON(500, utils.H{
				"code":    500,
				"message": "获取用户名失败",
			})
			return
		}
		fmt.Println("ping ", usernameValue)

		c.JSON(200, utils.H{
			"code":    200,
			"message": usernameValue,
		})
	}
}
