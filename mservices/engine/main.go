package engine

import (

	"github.com/caarlos0/env/v6"
	"github.com/pkg/errors"
	"google.golang.org/grpc"

	"github.com/boringmary/gomem/gen/gomem/mservices/engine"
	)

type Engine struct {
}

type config struct {
}

type Dependences struct {
	Registrar *grpc.Server
	GrpcConn  *grpc.ClientConn
}

func NewEngine(dep Dependences) (engine.EngineClient, error) {
	cfg := config{}
	if err := env.Parse(&cfg); err != nil {
		return nil, errors.WithStack(err)
	}

	en := &Engine{}

	engine.RegisterEngineServer(dep.Registrar, en)

	return engine.NewEngineClient(dep.GrpcConn), nil
}
