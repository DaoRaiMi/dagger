package account

import (
	"context"
	"time"

	"github.com/daoraimi/dagger/api"
	"github.com/daoraimi/dagger/box/log"
	"github.com/daoraimi/dagger/box/orm"
	"github.com/daoraimi/dagger/box/redis"
	"github.com/daoraimi/dagger/config"
	"github.com/daoraimi/dagger/model"
	"github.com/daoraimi/dagger/share"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type Repo struct{}

func (a *Repo) AddUser(ctx context.Context, req *api.AddUserRequest) (*api.AddUserResponse, error) {
	// 检查用户是否存在
	user, err := model.DaggerUser{}.FindByUsername(req.Username)
	if err != nil {
		return nil, err
	}
	if user != nil {
		return nil, api.Error{Code: api.InvalidArgument, Msg: "抱歉，该帐号已经存在"}
	}

	user, err = model.DaggerUser{}.FindByPhone(req.Phone)
	if err != nil {
		return nil, err
	}
	if user != nil {
		return nil, api.Error{Code: api.InvalidArgument, Msg: "抱歉，该手机号已被使用"}
	}

	user, err = model.DaggerUser{}.FindByEmail(req.Email)
	if err != nil {
		return nil, err
	}
	if user != nil {
		return nil, api.Error{Code: api.InvalidArgument, Msg: "抱歉，该邮件地址已被使用"}
	}

	// generate random password for new user
	randomBytePassword := GenerateRandomPassword(share.RandomPasswordLength)
	randomPassword, err := EncryptUserPassword(randomBytePassword)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	newUser := &model.DaggerUser{
		Username: req.Username,
		Nickname: req.Nickname,
		Password: randomPassword,
		RoleId:   req.RoleId,
		Phone:    req.Phone,
		Email:    req.Email,
	}

	if err = orm.R().Create(newUser).Error; err != nil {
		return nil, errors.WithStack(err)
	}

	// todo Send password to user
	log.Info(string(randomBytePassword), zap.String("RandomPassword", "yes"))
	return &api.AddUserResponse{UserId: newUser.ID}, nil
}

func (a *Repo) DeleteUser(ctx context.Context, req *api.DeleteUserRequest) (*api.DeleteUserResponse, error) {
	return &api.DeleteUserResponse{}, nil
}

func (a *Repo) UpdateUser(ctx context.Context, req *api.UpdateUserRequest) (*api.UpdateUserResponse, error) {
	return &api.UpdateUserResponse{}, nil
}

func (a *Repo) Login(ctx context.Context, req *api.LoginRequest) (*api.LoginResponse, error) {
	// 检查登录失败次数
	currentFailedLoginCount, err := redis.R().Get(share.GetKeyFailedLoginCount(req.Username)).Int()
	if err != nil && err != redis.Nil {
		return nil, errors.WithStack(err)
	}
	if currentFailedLoginCount > share.MaxFailedLoginCount {
		log.Warn("用户失败登录次数过多，已锁定", zap.String("username", req.Username))
		return nil, api.Error{Code: api.PermissionDenied, Msg: "登录失败次数过多，帐号被锁定"}
	}

	// 验证用户名和密码
	user, err := model.DaggerUser{}.FindByUsername(req.Username)
	if err != nil {
		return nil, err
	}
	if user == nil {
		// todo block login ?
		return nil, api.Error{Code: api.InvalidArgument, Msg: "帐号或密码错误"}
	}

	if !ValidateCredential([]byte(req.Password), []byte(user.Password)) {
		_, err := redis.R().Incr(share.GetKeyFailedLoginCount(req.Username)).Result()
		if err != nil {
			return nil, errors.WithStack(err)
		}
		_, err = redis.R().Expire(share.GetKeyFailedLoginCount(req.Username), time.Duration(share.KeyFailedLoginCountExpiration*time.Second)).Result()
		if err != nil {
			return nil, errors.WithStack(err)
		}
		return nil, api.Error{Code: api.InvalidArgument, Msg: "帐号或密码不正确"}
	}

	// 生成token
	token, err := GenerateUserToken(user.ID)
	if err != nil {
		return nil, err
	}

	// 更新帐号最近一次登录时间
	_ = model.DaggerUser{}.TouchLoginTime(user.ID)

	var lastLoginTime string
	if user.LastLoginTime != nil {
		lastLoginTime = user.LastLoginTime.Format(time.RFC3339)
	}

	return &api.LoginResponse{
		UserID:         user.ID,
		Username:       user.Username,
		Phone:          user.Phone,
		Email:          user.Email,
		RoleId:         user.RoleId,
		Token:          token,
		PermissionList: []uint64{},
		LastLoginTime:  lastLoginTime,
	}, nil
}

func (a *Repo) Logout(ctx context.Context, req *api.LogoutRequest) (*api.LogoutResponse, error) {
	userID := ctx.Value(share.ContextKeyUserID).(uint64)
	signature := ctx.Value(share.ContextKeyTokenSignature).(string)

	_, err := redis.R().Set(share.GetKeyTokenBlacklist(userID, signature), 1, config.GetDuration("token.expireDuration")).Result()
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &api.LogoutResponse{}, nil
}

func (a *Repo) ValidateUserPerm(ctx context.Context, req *api.ValidateUserPermRequest) (*api.ValidateUserPermResponse, error) {
	return &api.ValidateUserPermResponse{}, nil
}
