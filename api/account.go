package api

import (
	"regexp"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

var PhoneRegexp, EmailRegexp *regexp.Regexp

func init() {
	// 校验手机号码正则
	PhoneRegexp = regexp.MustCompile(`^1[3456789]\d{9}$`)
	// 校验邮箱正则
	EmailRegexp = regexp.MustCompile(`^([A-Za-z0-9_\-\.])+\@([A-Za-z0-9_\-\.])+\.([A-Za-z]{2,4})$`)
}

// 添加帐号
type AddUserRequest struct {
	Username string `json:"username"`
	Nickname string `json:"nickname"`
	RoleId   uint64 `json:"role_id"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
}

func (r *AddUserRequest) Validate() error {
	r.Username = strings.TrimSpace(r.Username)
	r.Nickname = strings.TrimSpace(r.Nickname)
	r.Phone = strings.TrimSpace(r.Phone)
	r.Email = strings.TrimSpace(r.Email)
	if r.Username == "" {
		return Error{InvalidArgument, "抱歉，帐号名不能为空"}
	}
	if r.Phone == "" || !PhoneRegexp.MatchString(r.Phone) {
		return Error{InvalidArgument, "抱歉，手机地址不能为空并且格式要正确"}
	}
	if r.Email == "" || !EmailRegexp.MatchString(r.Email) {
		return Error{InvalidArgument, "抱歉，邮件地址不能为空并且格式要正确"}
	}

	return nil
}

type AddUserResponse struct {
	UserId uint64 `json:"user_id"`
}

//
// 删除帐号
type DeleteUserRequest struct{}
type DeleteUserResponse struct{}

//
// 更新帐号
type UpdateUserRequest struct{}
type UpdateUserResponse struct{}

//
// 登录
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (r *LoginRequest) Validate() error {
	r.Username = strings.TrimSpace(r.Username)
	if r.Username == "" || r.Password == "" {
		return Error{InvalidArgument, "抱歉，帐号和密码不能为空"}
	}
	return nil
}

type TokenClaim struct {
	UserID uint64 `json:"user_id"`
	jwt.StandardClaims
}

type LoginResponse struct {
	UserID         uint64   `json:"user_id"`
	Username       string   `json:"username"`
	Nickname       string   `json:"nickname"`
	Phone          string   `json:"phone"`
	Email          string   `json:"email"`
	RoleId         uint64   `json:"role_id"`
	Token          string   `json:"token"`
	PermissionList []uint64 `json:"permission_list"`
	LastLoginTime  string   `json:"last_login_time"`
}

// 注销
type LogoutRequest struct {
	UserID uint64 `json:"user_id"`
}
type LogoutResponse struct{}

// 获取用户Token信息
type GetUserTokenInfoRequest struct {
	Token string
}

type GetUserTokenInfoResponse struct {
	UserID uint64
}

// 验证用户权限
type ValidateUserPermRequest struct {
	UserID   uint64
	PermList []uint32
}

type ValidateUserPermResponse struct {
	UserID uint64
}
