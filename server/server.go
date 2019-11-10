package server

import (
	accountRepo "github.com/daoraimi/dagger/repository/account"
	envRepo "github.com/daoraimi/dagger/repository/environment"
	"github.com/daoraimi/dagger/server/account"
	"github.com/daoraimi/dagger/server/environment"
	"github.com/gin-gonic/gin"
)

type Server struct {
	AccountServer     *account.Server
	EnvironmentServer *environment.Server
}

func New() *Server {
	return &Server{
		AccountServer:     &account.Server{&accountRepo.Repo{}},
		EnvironmentServer: &environment.Server{&envRepo.Repo{}},
	}
}

func (s *Server) Run() error {
	gin.SetMode(gin.ReleaseMode)
	rootRouter := gin.New()
	rootRouter.Use(gin.Recovery())
	rootRouter.Use(CustomLogger())

	rootRouter.POST("/dagger/login", s.AccountServer.LoginHandler)

	userRouter := rootRouter.Group("/dagger/users")
	{
		userRouter.Use(RequireToken())
		userRouter.POST("", s.AccountServer.AddUserHandler)
		userRouter.GET("", s.AccountServer.UserListHandler)
		userRouter.PATCH("/:user_id", s.AccountServer.UpdateUserHandler)
		userRouter.POST("logout", s.AccountServer.LogoutHandler)
	}

	envRouter := rootRouter.Group("/dagger/environments")
	{
		envRouter.Use(RequireToken())
		envRouter.POST("", s.EnvironmentServer.AddEnvironmentHandler)
		envRouter.GET("", s.EnvironmentServer.ListEnvironmentHandler)
		envRouter.PATCH("/:environment_id", s.EnvironmentServer.ModifyEnvironmentHandler)
		envRouter.DELETE("/:environment_id", s.EnvironmentServer.DeleteEnvironmentHandler)
	}

	return rootRouter.Run()
}

func (s *Server) Stop() {}
