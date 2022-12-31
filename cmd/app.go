package main

import (
	serverInit "digitalrepublic/server"
)

var (
	server serverInit.Server
)

func main() {
	server = serverInit.New()
	server.Start()
}
