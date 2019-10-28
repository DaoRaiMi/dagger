package service

import (
	"context"

	"github.com/daoraimi/dagger/api"
)

type AccountService interface {
	// 添加帐号
	AddUser(ctx context.Context, req *api.AddUserRequest) (*api.AddUserResponse, error)
	// 删除帐号
	DeleteUser(ctx context.Context, req *api.DeleteUserRequest) (*api.DeleteUserResponse, error)
	// 更新帐号
	UpdateUser(ctx context.Context, req *api.UpdateUserRequest) (*api.UpdateUserResponse, error)
	// 登录
	Login(ctx context.Context, req *api.LoginRequest) (*api.LoginResponse, error)
	// 注销
	Logout(ctx context.Context, req *api.LogoutRequest) (*api.LogoutResponse, error)
	// 验证用户权限
	ValidateUserPerm(ctx context.Context, req *api.ValidateUserPermRequest) (*api.ValidateUserPermResponse, error)
}
