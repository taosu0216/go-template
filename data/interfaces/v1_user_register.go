package interfaces

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"go-template/data/cons"
	"go-template/data/db/ent/user"
	"go-template/service/types"
	"time"
)

func (dr *DataRepo) SendRegisterEmailInData(ctx context.Context, email, verifyCode string) (*types.SendRegisterEmailResp, error) {
	start := time.Now()
	defer dr.ToolsCtx.LogDuration("SendRegisterEmailInData", start)

	err := dr.Cache.Set(ctx, email, verifyCode, time.Second*120).Err()
	if err != nil {
		dr.ErrorfInData("SendRegisterEmailInData [Set to Cache] err is: %v", err)
		return &types.SendRegisterEmailResp{
			BasicResp: types.BasicResp{
				Code: cons.InternalServerError,
				Msg:  cons.InternalErrorMsg,
				Err:  cons.InternalErr,
			},
		}, err
	}

	return &types.SendRegisterEmailResp{
		BasicResp: types.BasicResp{
			Code: cons.BasicSuccessCode,
			Msg:  cons.BasicSuccessMsg,
			Err:  nil,
		},
	}, nil
}

func (dr *DataRepo) RegisterUserInData(ctx context.Context, req *types.RegisterUserReq) (*types.RegisterUserResp, error) {
	start := time.Now()
	defer dr.ToolsCtx.LogDuration("RegisterUserInData", start)

	info, err := dr.Cache.Get(ctx, req.Email).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			// key不存在的情况，说明验证码已过期或未发送
			return &types.RegisterUserResp{
				BasicResp: types.BasicResp{
					Code: cons.VerifyCodeExpiredCode, // 需要在cons中定义适当的错误码
					Msg:  cons.VerifyCodeExpiredMsg,
					Err:  nil,
				},
			}, nil // 这种情况下返回nil作为error，因为这是一个业务逻辑错误而不是系统错误
		}
		// 其他错误情况（比如Redis连接错误等）
		dr.ErrorfInData("RegisterUserInData [Get VerifyCode from Cache] err is: %v", err)
		return &types.RegisterUserResp{
			BasicResp: types.BasicResp{
				Code: cons.InternalServerError,
				Msg:  cons.InternalErrorMsg,
				Err:  cons.InternalErr,
			},
		}, err
	}
	if info == "" || info != req.VerifyCode {
		dr.WarnfInData("RegisterUserInData [Get from Cache] info is: %v", info)
		return &types.RegisterUserResp{
			BasicResp: types.BasicResp{
				Code: cons.VerifyCodeFailCode,
				Msg:  cons.VerifyCodeFailMsg,
				Err:  nil,
			},
		}, nil
	}

	isExist, err := dr.DB.User.Query().Where(
		user.Or(
			user.NameEQ(req.UserName),
			user.EmailEQ(req.Email),
		),
	).Exist(ctx)
	if err != nil {
		dr.ErrorfInData("RegisterUserInData [isExist] err is: %v", err)
		return &types.RegisterUserResp{
			BasicResp: types.BasicResp{
				Code: cons.InternalServerError,
				Msg:  cons.InternalErrorMsg,
				Err:  cons.InternalErr,
			},
		}, err
	}
	if isExist {
		dr.InfofInData("RegisterUserInData [isExist] isExist")
		return &types.RegisterUserResp{
			BasicResp: types.BasicResp{
				Code: cons.UserIsExistCode,
				Msg:  cons.UserIsExistMsg,
				Err:  nil,
			},
		}, nil
	}

	id, _ := uuid.NewUUID()
	_, err = dr.DB.User.Create().
		SetID(id.String()).
		SetName(req.UserName).
		SetPasswd(req.PassWd).
		SetEmail(req.Email).
		SetCreateTime(time.Now().Format("2006-01-02 15:04:05")).
		Save(ctx)
	if err != nil {
		dr.ErrorfInData("RegisterUserInData [CreateUser] [%s] err is: %v", req.UserName, err)
		return &types.RegisterUserResp{
			BasicResp: types.BasicResp{
				Code: cons.InternalServerError,
				Msg:  cons.InternalErrorMsg,
				Err:  cons.InternalErr,
			},
		}, err
	}
	go func() {
		userKey := fmt.Sprintf("user_%s", req.Email)
		err = dr.Cache.HSet(ctx, cons.UserMapKey, userKey, cons.UserStatusRegistered).Err()
		if err != nil {
			dr.ErrorfInData("RegisterUserInData [Cache HSet] err: %v", err)
		}
	}()
	return &types.RegisterUserResp{
		BasicResp: types.BasicResp{
			Code: cons.BasicSuccessCode,
			Msg:  cons.BasicSuccessMsg,
			Err:  nil,
		},
	}, nil
}

func (dr *DataRepo) IsUserExistInDataByEmail(ctx context.Context, email string) (string, error) {
	userKey := fmt.Sprintf("user_%s", email)
	value, err := dr.Cache.HGet(ctx, cons.UserMapKey, userKey).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return cons.UserStatusNotRegister, nil
		}
		return cons.UserStatusNotRegister, err
	}
	return value, nil

	//return dr.DB.User.Query().Where(user.EmailEQ(email)).Exist(ctx)
}
