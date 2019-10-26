package api

import (
	"github.com/daoraimi/dagger/box/log"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

const (
	Ok               = 200
	InvalidArgument  = 400
	Unauthorized     = 401
	PermissionDenied = 403
	NotFound         = 404
	Internal         = 500
	Unavailable      = 503
)

// Custom error
type Error struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func (e Error) Error() string {
	return e.Msg
}

func RespError(c *gin.Context, err error) {
	switch err.(type) {
	case Error:
		status := err.(Error)
		c.JSON(status.Code, gin.H{"msg": status.Msg})
		c.Abort()
	default:
		log.Error("Server Internal Error", zap.Error(err))
		c.JSON(Internal, gin.H{"msg": "Server Internal Error"})
		c.Abort()
	}
}

func RespOk(c *gin.Context, obj interface{}) {
	c.JSON(Ok, obj)
}
