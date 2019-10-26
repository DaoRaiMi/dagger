package account

import (
	"github.com/daoraimi/dagger/api"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func (s Server) AddUserHandler(c *gin.Context) {
	var err error
	var req api.AddUserRequest

	if err = c.MustBindWith(&req, binding.JSON); err != nil {
		api.RespError(c, err)
		return
	}

	if err = req.Validate(); err != nil {
		api.RespError(c, err)
		return
	}

	resp, err := s.AddUser(c, &req)
	if err != nil {
		api.RespError(c, err)
		return
	}

	api.RespOk(c, resp)
}
