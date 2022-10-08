package main

import (
	"fmt"
	"go-orders-api/controllers"
	"go-orders-api/utils"

	"github.com/gin-gonic/gin"
)

const PORT = ":8000"

func main() {
	var db = utils.NewDB()
	var inDB = controllers.NewInDB(db)

	var router = gin.Default()
	router.GET("/orders", inDB.GetAllOrders)
	router.POST("/orders", inDB.CreateOrder)
	router.PUT("/orders/:orderId", inDB.UpdateOrderById)
	router.DELETE("/orders/:orderId", inDB.DeleteOrderById)

	fmt.Printf("Server running at http://localhost%s\n", PORT)
	router.Run(PORT)
}
