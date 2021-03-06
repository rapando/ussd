package main

import (
	"github.com/joho/godotenv"
	"github.com/rapando/ussd/controller"
)

var b controller.Base

func main() {
	_ = godotenv.Load()
	b.Init()

	b.Run()
}
