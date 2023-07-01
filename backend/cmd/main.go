package main

import "bebrah/app/server"

func main() {
	r := server.SetupRouter()
	r.Run()
}
