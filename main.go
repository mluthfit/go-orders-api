package main

import (
	"go-orders-api/controllers"
	"go-orders-api/utils"

	"github.com/gin-gonic/gin"
)

const PORT = ":8000"

func main() {
	var db = utils.NewDB()
	var router = gin.Default()
	var server = controllers.NewServer(router, db)

	server.AddRoutes()
	server.Run(PORT)
}
