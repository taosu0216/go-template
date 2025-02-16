package biz

import (
	"context"
	"fmt"
	"go-template/logger"
	"go-template/service/types"
	"time"
)

type TemplateRepo interface {
	GetSrvAddr() string
	GetSrvPort() string
	GetSrvExitWait() time.Duration
	GetMailInfo() (string, string, string, string)
	GetJwtKey() string
	GetVideoStoragePath() string

	SendRegisterEmailInData(ctx context.Context, email, verifyCode string) (*types.SendRegisterEmailResp, error)
	RegisterUserInData(ctx context.Context, req *types.RegisterUserReq) (*types.RegisterUserResp, error)

	SendChangePasswordEmailInData(ctx context.Context, email, verifyCode string) (*types.SendChangePasswdEmailResp, error)
	ChangePasswordInData(ctx context.Context, req *types.ChangePasswordReq) (*types.ChangePasswordResp, error)
	IsUserExistInDataByEmail(ctx context.Context, email string) (string, error)
	IsPasswordRightInData(ctx context.Context, userName, password string) (bool, error)

	SendKafkaMessage(topic string, message []byte) error
}

type TemplateService struct {
	repo TemplateRepo
	*logger.ZapLogger
}

func NewTemplateService(repo TemplateRepo, l *logger.ZapLogger) *TemplateService {
	return &TemplateService{repo, l}
}

func (uc *TemplateService) LogDuration(name string, start time.Time) {
	duration := time.Since(start).Seconds()
	str := fmt.Sprintf("func: [%s] | duration=%f s\n", name, duration)
	uc.LogDurationInBiz(str)
}

func (uc *TemplateService) GetSrvAddr() string            { return uc.repo.GetSrvAddr() }
func (uc *TemplateService) GetSrvPort() string            { return uc.repo.GetSrvPort() }
func (uc *TemplateService) GetSrvExitWait() time.Duration { return uc.repo.GetSrvExitWait() }
func (uc *TemplateService) GetJwtKey() string             { return uc.repo.GetJwtKey() }
func (uc *TemplateService) GetVideoStoragePath() string   { return uc.repo.GetVideoStoragePath() }
