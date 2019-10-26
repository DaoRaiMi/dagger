package dconst

const PasswordElements = "abcdefghjkmnpqrstuvwxy3456789ABCDEFGHJKLMNPQRSTUVWXY#$%&*()_+=<>?"

const (
	RandomPasswordLength          = 8                   // 初始化随机密码的长度
	MaxFailedLoginCount           = 5                   // 最大登陆失败次数
	UserTokenFormat               = "%s_%d_Token"       // 用户token
	TokenInfoFormat               = "%s_UserInfo_%s"    // token对应的用户信息
	FailedLoginCountFormat        = "%s_FailedLogin_%s" // 用户登录失败的次数
	KeyFailedLoginCountExpiration = 600                 // 登录失败次数超限后，帐号锁定时长
)
