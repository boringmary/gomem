package engine

import (
	"github.com/boringmary/gomem/gen/mservices/engine"
	"github.com/caarlos0/env/v6"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

type Engine struct {
}

type config struct {
}

type Dependences struct {
	Registrator *grpc.Server
	GrpcConn  *grpc.ClientConn
}


func NewEngine(dep Dependences) error {
	cfg := config{}
	if err := env.Parse(&cfg); err != nil {
		return errors.WithStack(err)
	}

	en := &Engine{}
	engine.RegisterEngineServer(dep.Registrator, en)

	return nil
}