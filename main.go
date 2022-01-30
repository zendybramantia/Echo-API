package main

import (
	"echo-api/db"
	"echo-api/routes"
)

func main() {
	db.Init()

	e := routes.Init()

	e.Logger.Fatal(e.Start(":1002"))
}
