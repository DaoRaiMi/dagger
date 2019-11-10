package environment

import (
	"context"

	"github.com/daoraimi/dagger/api"
)

type Repo struct{}

//
func (r *Repo) AddEnvironment(ctx context.Context, req *api.AddEnvironmentRequest) (*api.AddEnvironmentResponse, error) {
	return &api.AddEnvironmentResponse{}, nil
}

//
func (r *Repo) DeleteEnvironment(ctx context.Context, req *api.DeleteEnvironmentRequest) (*api.DeleteEnvironmentResponse, error) {
	return &api.DeleteEnvironmentResponse{}, nil
}

//
func (r *Repo) ModifyEnvironment(ctx context.Context, req *api.ModifyEnvironmentRequest) (*api.ModifyEnvironmentResponse, error) {
	return &api.ModifyEnvironmentResponse{}, nil
}

//
func (r *Repo) ListEnvironment(ctx context.Context, req *api.ListEnvironmentRequest) (*api.ListEnvironmentResponse, error) {
	return &api.ListEnvironmentResponse{}, nil
}
