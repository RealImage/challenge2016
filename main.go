package main

import (
	"github.com/nikhilsiwach28/Cinema-Distribution-System/config"
	"github.com/nikhilsiwach28/Cinema-Distribution-System/handler"
)

func main() {
	config := config.NewEnvServerConfig()
	handler.StartHttpServer(config)
}
