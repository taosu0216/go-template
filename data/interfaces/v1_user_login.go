package interfaces

import (
	"context"
	"go-template/data/db/ent/user"
	"time"
)

func (dr *DataRepo) IsPasswordRightInData(ctx context.Context, userName, password string) (bool, error) {
	now := time.Now()
	defer dr.ToolsCtx.LogDuration("IsPasswordRightInData", now)

	// 从数据库中获取密码
	info, err := dr.DB.User.Query().Where(user.NameEQ(userName)).First(ctx)
	if err != nil {
		dr.ErrorfInData("IsPasswordRightInData [Query User] err is: %v", err)
		return false, err
	}

	if info.Passwd != password {
		return false, nil
	}

	return true, nil
}
