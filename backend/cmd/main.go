package main

import (
	"bebrah/app/db"
	"bebrah/app/server"
)

func main() {
	db.InitDb(".")
	r := server.SetupRouter()
	r.Run()
}
