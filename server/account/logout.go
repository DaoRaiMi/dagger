package account

import (
	"github.com/daoraimi/dagger/api"
	"github.com/gin-gonic/gin"
)

func (s Server) LogoutHandler(c *gin.Context) {
	resp, err := s.Logout(c, &api.LogoutRequest{})
	if err != nil {
		api.RespError(c, err)
	}

	api.RespOk(c, resp)
}
