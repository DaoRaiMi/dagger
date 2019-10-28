package server

import (
	accountRepo "github.com/daoraimi/dagger/repository/account"
	"github.com/daoraimi/dagger/server/account"
	"github.com/gin-gonic/gin"
)

type Server struct {
	AccountServer *account.Server
}

func New() *Server {
	return &Server{
		AccountServer: &account.Server{&accountRepo.Repo{}},
	}
}

func (s *Server) Run() error {
	gin.SetMode(gin.ReleaseMode)
	rootRouter := gin.New()
	rootRouter.Use(gin.Recovery())
	rootRouter.Use(CustomLogger())

	rootRouter.POST("/dagger/login", s.AccountServer.LoginHandler)

	userRouter := rootRouter.Group("/dagger/users")
	userRouter.Use(RequireToken())
	userRouter.POST("", s.AccountServer.AddUserHandler)
	userRouter.POST("logout", s.AccountServer.LogoutHandler)

	return rootRouter.Run()
}

func (s *Server) Stop() {}
