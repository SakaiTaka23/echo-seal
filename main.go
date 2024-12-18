package main

import (
	"echo-seal/common"
	yescryptv1 "echo-seal/gen/yescrypt/v1"
	"echo-seal/yescrypt"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	"github.com/BurntSushi/toml"
	"google.golang.org/grpc"
)

func main() {
	var cfg AppConfig

	_, err := toml.DecodeFile("./config.toml", &cfg)
	if err != nil {
		panic(err)
	}

	salt := common.NewSaltGenerator(cfg.Common.SaltLength)
	y := yescrypt.New(salt, cfg.Yescrypt.YescryptNLog2, cfg.Yescrypt.YescryptR)

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.Common.Port))
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	yescryptv1.RegisterYescryptServiceServer(s, yescrypt.NewService(*y))

	go func() {
		log.Printf("started server on port %v", cfg.Common.Port)

		if err = s.Serve(listener); err != nil {
			panic(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	s.GracefulStop()
}
