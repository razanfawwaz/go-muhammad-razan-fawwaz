package main

import (
	"part3-orm/config"
	"part3-orm/middlewares"
	"part3-orm/routes"
)

func main() {
	config.InitDB()
	e := routes.New()
	// middleware logger from middlewares folder
	middlewares.LogMiddlewares(e)
	e.Logger.Fatal(e.Start(":8000"))
}
