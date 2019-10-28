package server

import (
	"time"

	"github.com/daoraimi/dagger/api"
	"github.com/daoraimi/dagger/box/log"
	"github.com/daoraimi/dagger/box/redis"
	"github.com/daoraimi/dagger/config"
	"github.com/daoraimi/dagger/share"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

// token middleware
func RequireToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		if token == "" {
			api.RespError(c, api.Error{api.Unauthorized, "Invalid token"})
			return
		}

		var claim api.TokenClaim
		validToken, err := jwt.ParseWithClaims(token, &claim, func(tk *jwt.Token) (i interface{}, e error) {
			return []byte(config.GetString("token.secret")), nil
		})
		if err != nil {
			api.RespError(c, api.Error{api.Unauthorized, "Invalid token"})
			return
		}

		if claim.UserID == 0 {
			api.RespError(c, api.Error{api.Unauthorized, "Invalid token"})
			return
		}

		// check token blacklist
		exist, err := redis.R().Exists(share.GetKeyTokenBlacklist(claim.UserID, validToken.Signature)).Result()
		if err != nil {
			api.RespError(c, errors.WithStack(err))
		}
		if exist != 0 {
			api.RespError(c, api.Error{api.Unauthorized, "Invalid Token"})
		}

		c.Set(share.ContextKeyUserID, claim.UserID)
		c.Set(share.ContextKeyTokenSignature, validToken.Signature)

		c.Next()
	}
}

// log middleware
func CustomLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()
		scheme := c.Request.URL.Scheme
		host := c.Request.Host
		httpProto := c.Request.Proto
		method := c.Request.Method
		uri := c.Request.URL.RequestURI()
		userAgent := c.Request.UserAgent()

		c.Next()

		stopTime := time.Now()
		latency := stopTime.Sub(startTime)
		clientIP := c.ClientIP()
		statusCode := c.Writer.Status()
		bodySize := c.Writer.Size()

		log.Info("",
			zap.String("Client", clientIP),
			zap.String("Scheme", scheme),
			zap.String("Host", host),
			zap.String("Method", method),
			zap.String("Uri", uri),
			zap.String("Proto", httpProto),
			zap.Int("Status", statusCode),
			zap.Int("BodySize", bodySize),
			zap.Duration("Latency", latency),
			zap.String("UserAgent", userAgent),
		)
	}
}
