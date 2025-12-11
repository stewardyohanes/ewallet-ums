package cmd

import (
	tokenvalidation "ewallet-ums/cmd/proto"
	"ewallet-ums/helpers"
	"fmt"
	"log"
	"net"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func ServeGRPC() {
	dependency := dependencyInject()

	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", helpers.GetEnv("GRPC_PORT", "7000")))
	if err != nil {
		log.Fatal("Error to listen GRPC server", err)
	}


	grpcServer := grpc.NewServer()

	tokenvalidation.RegisterTokenValidationServer(grpcServer, dependency.TokenValidationAPI)


	logrus.Info("Listening GRPC server on port", fmt.Sprintf(":%s", helpers.GetEnv("GRPC_PORT", "7000")))
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatal("Error starting GRPC server", err)
	}
}
