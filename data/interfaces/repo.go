package interfaces

import (
	"go-template/data"
	"go-template/logger"
	"strconv"
	"time"
)

type DataRepo struct {
	*data.ToolsCtx

	*logger.ZapLogger
}

func NewDataRepo(logger *logger.ZapLogger) *DataRepo {
	return &DataRepo{data.Tools, logger}
}

func (dr *DataRepo) GetSrvAddr() string { return dr.Cfg.RouterCfg.Addr }
func (dr *DataRepo) GetSrvPort() string { return dr.Cfg.RouterCfg.Port }
func (dr *DataRepo) GetSrvExitWait() time.Duration {
	exitWait, err := strconv.Atoi(dr.Cfg.RouterCfg.ExitWait)
	if err != nil {
		return time.Second * 1
	}
	return time.Millisecond * time.Duration(exitWait)
}
func (dr *DataRepo) GetJwtKey() string { return dr.Cfg.JwtCfg.JwtSecret }

func (dr *DataRepo) GetMailInfo() (string, string, string, string) {
	return dr.Cfg.MailCfg.Addr, dr.Cfg.MailCfg.Host, dr.Cfg.MailCfg.Password, dr.Cfg.MailCfg.From
}
