package main

import (
	"belajar-go-echo/config"
	"belajar-go-echo/routes"
)

func main() {
	config.InitDB()
	e := routes.New()
	e.Logger.Fatal(e.Start(":8000"))
}
