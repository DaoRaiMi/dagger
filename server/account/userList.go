package account

import (
	"github.com/daoraimi/dagger/api"
	"github.com/gin-gonic/gin"
)

func (s Server) UserListHandler(c *gin.Context) {
	var err error
	var req api.UserListRequest

	_ = c.ShouldBindQuery(&req)

	if err = req.Validate(); err != nil {
		api.RespError(c, err)
		return
	}

	resp, err := s.UserList(c, &req)
	if err != nil {
		api.RespError(c, err)
		return
	}

	api.RespOk(c, resp)
}
