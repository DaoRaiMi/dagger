package account

import (
	"github.com/daoraimi/dagger/api"
	"github.com/gin-gonic/gin"
)

func (s Server) UpdateUserHandler(c *gin.Context) {
	var err error
	var req api.UpdateUserRequest
	var uriUserID api.UriUserID

	_ = c.ShouldBindUri(&uriUserID)
	if uriUserID.UserID == 0 {
		api.RespError(c, api.Error{api.InvalidArgument, "url中user_id不能为空"})
		return
	}

	_ = c.ShouldBindJSON(&req)

	if err = req.Validate(); err != nil {
		api.RespError(c, err)
		return
	}

	resp, err := s.UpdateUser(c, &req)
	if err != nil {
		api.RespError(c, err)
		return
	}

	api.RespOk(c, resp)
}
