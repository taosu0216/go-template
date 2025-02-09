package interfaces

import (
	"context"
	"errors"
	"github.com/redis/go-redis/v9"
	"go-template/data/cons"
	"go-template/data/db/ent/user"
	"go-template/service/types"
	"time"
)

func (dr *DataRepo) SendChangePasswordEmailInData(ctx context.Context, email, verifyCode string) (*types.SendChangePasswdEmailResp, error) {
	start := time.Now()
	defer dr.ToolsCtx.LogDuration("SendChangePasswordEmailInData", start)

	err := dr.Cache.Set(ctx, email, verifyCode, time.Second*120).Err()
	if err != nil {
		dr.ErrorfInData("SendChangePasswordEmailInData [Set to Cache] err is: %v", err)
		return &types.SendChangePasswdEmailResp{
			BasicResp: types.BasicResp{
				Code: cons.InternalServerError,
				Msg:  cons.InternalErrorMsg,
				Err:  cons.InternalErr,
			},
		}, err
	}

	return &types.SendChangePasswdEmailResp{
		BasicResp: types.BasicResp{
			Code: cons.BasicSuccessCode,
			Msg:  cons.BasicSuccessMsg,
			Err:  nil,
		},
	}, nil
}

func (dr *DataRepo) ChangePasswordInData(ctx context.Context, req *types.ChangePasswordReq) (*types.ChangePasswordResp, error) {
	// 获取验证码
	info, err := dr.Cache.Get(ctx, req.Email).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			// key不存在的情况，说明验证码已过期或未发送
			return &types.ChangePasswordResp{
				BasicResp: types.BasicResp{
					Code: cons.VerifyCodeExpiredCode, // 需要在cons中定义适当的错误码
					Msg:  cons.VerifyCodeExpiredMsg,
					Err:  nil,
				},
			}, nil // 这种情况下返回nil作为error，因为这是一个业务逻辑错误而不是系统错误
		}
		// 其他错误情况（比如Redis连接错误等）
		dr.ErrorfInData("ChangePasswordInData [Get VerifyCode from Cache] err is: %v", err)
		return &types.ChangePasswordResp{
			BasicResp: types.BasicResp{
				Code: cons.InternalServerError,
				Msg:  cons.InternalErrorMsg,
				Err:  cons.InternalErr,
			},
		}, err
	}
	if info == "" || info != req.VerifyCode {
		dr.WarnfInData("ChangePasswordInData [Get from Cache] info is: %v", info)
		return &types.ChangePasswordResp{
			BasicResp: types.BasicResp{
				Code: cons.VerifyCodeFailCode,
				Msg:  cons.VerifyCodeFailMsg,
				Err:  nil,
			},
		}, nil
	}

	// 看数据库里是否存在这个用户
	isExist, err := dr.DB.User.Query().Where(user.EmailEQ(req.Email)).Exist(ctx)
	if err != nil {
		dr.ErrorfInData("ChangePasswordInData [isExist] err is: %v", err)
		return &types.ChangePasswordResp{
			BasicResp: types.BasicResp{
				Code: cons.InternalServerError,
				Msg:  cons.InternalErrorMsg,
				Err:  cons.InternalErr,
			},
		}, err
	}
	if !isExist {
		return &types.ChangePasswordResp{
			BasicResp: types.BasicResp{
				Code: cons.UserNotFoundCode,
				Msg:  cons.UserNotFoundMsg,
				Err:  nil,
			},
		}, nil
	}

	// 更新密码
	err = dr.DB.User.Update().Where(user.EmailEQ(req.Email)).SetPasswd(req.PassWd).Exec(ctx)
	if err != nil {
		dr.ErrorfInData("ChangePasswordInData [UpdatePasswd] err is: %v", err)
		return &types.ChangePasswordResp{
			BasicResp: types.BasicResp{
				Code: cons.InternalServerError,
				Msg:  cons.InternalErrorMsg,
				Err:  cons.InternalErr,
			},
		}, err
	}

	return &types.ChangePasswordResp{
		BasicResp: types.BasicResp{
			Code: cons.BasicSuccessCode,
			Msg:  cons.BasicSuccessMsg,
			Err:  nil,
		},
	}, nil
}
