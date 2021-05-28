package main

import (
	"fmt"
	"github.com/boringmary/gomem/mservices/card"
	"github.com/boringmary/gomem/mservices/engine"
	"github.com/boringmary/gomem/mservices/server"
)

func main() {
	srv, err := server.NewServer()
	if err != nil {
		panic(err)
	}

	grpcConn, err := srv.GrpcDial()
	if err != nil {
		panic(err)
	}

	err = engine.NewEngine(engine.Dependences{
		Registrator: srv.RpcServer,
		GrpcConn:  grpcConn,
	})
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	err = card.NewCard(card.Dependences{
		Registrator: srv.RpcServer,
		GrpcConn:  grpcConn,
	})
	if err != nil {
		panic(err)
	}

	if err := srv.Serve(); err != nil {
		panic(err)
	}
}