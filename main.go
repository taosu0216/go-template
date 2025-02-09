package main

import (
	"go-template/data"
	"go-template/data/interfaces"
	"go-template/logger"
	"go-template/server/biz"
	"go-template/service/handler"
	"go-template/service/svc"
)

var (
	dataRepo *interfaces.DataRepo
	srvRepo  *svc.TemplateHandler
)

func init() {
	loggerObj := logger.InitLog()
	data.Init(loggerObj)

	dataRepo = interfaces.NewDataRepo(loggerObj)
	bizRepo := biz.NewTemplateService(dataRepo, loggerObj)
	srvRepo = svc.NewTemplateHandler(bizRepo, loggerObj)
}

func main() {
	h := handler.InitService(srvRepo)
	srvRepo.TemplateService.Info("server start.....")
	h.Spin()
}
