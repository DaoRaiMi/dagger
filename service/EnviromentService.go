package service

import (
	"context"

	"github.com/daoraimi/dagger/api"
)

type EnvironmentService interface {
	// add env
	AddEnvironment(ctx context.Context, req *api.AddEnvironmentRequest) (*api.AddEnvironmentResponse, error)

	// modify env
	ModifyEnvironment(ctx context.Context, req *api.ModifyEnvironmentRequest) (*api.ModifyEnvironmentResponse, error)

	// delete env
	DeleteEnvironment(ctx context.Context, req *api.DeleteEnvironmentRequest) (*api.DeleteEnvironmentResponse, error)

	// list env
	ListEnvironment(ctx context.Context, req *api.ListEnvironmentRequest) (*api.ListEnvironmentResponse, error)
}
