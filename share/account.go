package share

import "fmt"

const PasswordElements = "abcdefghjkmnpqrstuvwxy3456789ABCDEFGHJKLMNPQRSTUVWXY#$%&*()_+=<>?"

const (
	RandomPasswordLength          = 8                         // 初始化随机密码的长度
	MaxFailedLoginCount           = 5                         // 最大登陆失败次数
	UserTokenBlacklist            = "%s_TokenBlacklist_%d_%s" // 用户token黑名单 使用token.Signature区分
	FailedLoginCountFormat        = "%s_FailedLogin_%s"       // 用户登录失败的次数
	KeyFailedLoginCountExpiration = 600                       // 登录失败次数超限后，帐号锁定时长
	ContextKeyUserID              = "user_id"                 // 帐号ID在context中的key
	ContextKeyTokenSignature      = "token_signature"         // token.Signature在context中的key
)

func GetKeyTokenBlacklist(userID uint64, signature string) string {
	return fmt.Sprintf(UserTokenBlacklist, KeyPrefix, userID, signature)
}

func GetKeyFailedLoginCount(username string) string {
	return fmt.Sprintf(FailedLoginCountFormat, KeyPrefix, username)
}
