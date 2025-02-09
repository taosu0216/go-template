package svc

import (
	"fmt"
	"go-template/logger"
	"go-template/server/biz"
	"time"
)

type TemplateHandler struct {
	TemplateService *biz.TemplateService
	*logger.ZapLogger
}

func NewTemplateHandler(templateService *biz.TemplateService, l *logger.ZapLogger) *TemplateHandler {
	return &TemplateHandler{templateService, l}
}

func (uc *TemplateHandler) LogDuration(name string, start time.Time) {
	duration := time.Since(start).Milliseconds()
	str := fmt.Sprintf("func: [%s] | duration=%d ms", name, duration)
	uc.LogDurationInService(str)
}
