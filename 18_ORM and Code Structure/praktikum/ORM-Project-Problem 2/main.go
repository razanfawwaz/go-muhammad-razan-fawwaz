package main

import (
	"part3-orm/config"
	"part3-orm/routes"
)

func main() {
	config.InitDB()
	e := routes.New()
	e.Logger.Fatal(e.Start(":8000"))
}
