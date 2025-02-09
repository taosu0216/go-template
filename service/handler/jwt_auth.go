package handler

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/golang-jwt/jwt/v4"
	"go-template/service/svc"
	"strings"
)

func JWTAuthMiddleware(r *svc.TemplateHandler) app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		authHeaderBytes := c.GetHeader("Authorization")
		authHeader := string(authHeaderBytes)
		if authHeader == "" {
			c.JSON(401, utils.H{
				"code":    401,
				"message": "请求头中 auth 缺失",
			})
			c.Abort()
			return
		}

		// 按照 "Bearer <token>" 的格式从请求头中获取 token
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			c.JSON(401, utils.H{
				"code":    401,
				"message": "请求头中 auth 格式有误",
			})
			c.Abort()
			return
		}

		tokenString := parts[1]

		// 解析并验证 token
		claims := &MyClaims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(r.TemplateService.GetJwtKey()), nil
		})
		if err != nil || !token.Valid {
			c.JSON(401, utils.H{
				"code":    401,
				"message": "无效的 Token",
			})
			c.Abort()
			return
		}

		// 将用户名信息存入上下文
		c.Set("username", claims.UserName)

		// 继续处理请求
		c.Next(ctx)
	}
}

// 定义 Claims 结构体
type MyClaims struct {
	UserName string `json:"username"`
	jwt.RegisteredClaims
}
