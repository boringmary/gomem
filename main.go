package main

import (
	"fmt"
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

	engineCLi, err := engine.NewEngine(engine.Dependences{
		Registrar: srv.RpcServer,
		GrpcConn:  grpcConn,
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(engineCLi)


	if err := srv.Serve(); err != nil {
		panic(err)
	}
}