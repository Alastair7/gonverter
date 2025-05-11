package main

import (
	"github.com/Alastair7/gonverter/internal/server"
)

func main() {
	server := server.New(":8080")
	server.RunServer()
}
