package account

import (
	"github.com/daoraimi/dagger/api"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func (s *Server) LoginHandler(c *gin.Context) {
	var err error
	var req api.LoginRequest

	if err = c.MustBindWith(&req, binding.JSON); err != nil {
		api.RespError(c, err)
		return
	}

	if err = req.Validate(); err != nil {
		api.RespError(c, err)
		return
	}

	// 开始登录
	resp, err := s.Login(c, &req)
	if err != nil {
		api.RespError(c, err)
		return
	}

	api.RespOk(c, resp)
}
