package main

import (
	"os"

	"github.com/kuzmindeniss/prost/internal/bff/helpers"
	"github.com/kuzmindeniss/prost/internal/bff/router"
	"github.com/kuzmindeniss/prost/internal/db"
)

func init() {
	helpers.LoadEnv()
	db.ConnectToDb()
}

func main() {
	r := router.SetupRouter()
	r.Run(":" + os.Getenv("BFF_PORT"))
}
