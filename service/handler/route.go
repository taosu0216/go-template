package handler

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/hertz-contrib/cors"
	v1 "go-template/service/handler/api/v1"
	"go-template/service/handler/api/v1/auth"
	"go-template/service/handler/api/v1/user"
	"go-template/service/svc"
	"time"
)

func InitService(r *svc.TemplateHandler) *server.Hertz {
	h := server.Default(
		server.WithHostPorts(r.TemplateService.GetSrvAddr()+":"+r.TemplateService.GetSrvPort()),
		server.WithExitWaitTime(r.TemplateService.GetSrvExitWait()),
		server.WithHandleMethodNotAllowed(true),
		server.WithMaxRequestBodySize(1000*1024*1024),
	)

	h.Use(cors.New(cors.Config{
		//准许跨域请求网站,多个使用,分开,限制使用*
		AllowOrigins: []string{"*"},
		//准许使用的请求方式
		AllowMethods: []string{"POST", "GET"},
		//准许使用的请求表头
		AllowHeaders: []string{"Origin", "Authorization", "Content-Type"},
		//显示的请求表头
		ExposeHeaders: []string{"Content-Type"},
		//凭证共享,确定共享
		AllowCredentials: true,
		//容许跨域的原点网站,可以直接return true就万事大吉了
		AllowOriginFunc: func(origin string) bool { return true },
		//超时时间设定
		MaxAge: 24 * time.Hour,
	}))
	h.NoHijackConnPool = true

	v1Group := h.Group("/v1")
	v1Group.GET("/ping", v1.GeneratePingLogic())

	userGroup := v1Group.Group("/user")
	{
		userGroup.POST("/register/sendEmail", user.GenerateUserRegisterSendEmailLogic(r))
		userGroup.POST("/register", user.GenerateUserRegisterLogic(r))
		userGroup.POST("/changePasswd", user.GenerateUserChangePasswdLogic(r))
		userGroup.POST("/changePasswd/sendEmail", user.GenerateUserChangePasswdSendEmailLogic(r))
		userGroup.POST("/login", user.GenerateUserLoginLogic(r))
	}

	authGroup := v1Group.Group("/auth")
	authGroup.Use(JWTAuthMiddleware(r))
	{
		authGroup.GET("/ping", auth.GeneratePingLogic())
	}

	return h
}
