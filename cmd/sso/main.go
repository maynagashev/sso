package main

import (
	"fmt"
	"github.com/maynagashev/sso/internal/config"
)

func main() {
	cfg := config.MustLoad()

	fmt.Printf("%+v", cfg)

	// TODO: init logger

	// TODO: init app

	// TODO: init grpc server

}
