package main

import (
	"ewallet-ums/cmd"
	"ewallet-ums/helpers"
)

func main() {
	// Load Config
	helpers.SetupConfig()

	// Load Logger
	helpers.SetupLogger()

	// Load Database
	helpers.SetupMySQL()

	// Run GRPC Server
	go cmd.ServeGRPC()

	// Run HTTP Server
	cmd.ServeHTTP()
}
