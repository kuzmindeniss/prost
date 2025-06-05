package main

import (
	"os"

	"github.com/kuzmindeniss/prost/internal/bff/helpers"
	"github.com/kuzmindeniss/prost/internal/bff/router"
)

func init() {
	helpers.LoadEnv()
	helpers.ConnectToDb()
}

func main() {
	r := router.SetupRouter()
	r.Run(":" + os.Getenv("PORT"))
}
