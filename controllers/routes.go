package controllers

func (server *Server) AddRoutes() {
	server.router.GET("/orders", server.GetAllOrders)
	server.router.POST("/orders", server.CreateOrder)
	server.router.PUT("/orders/:orderId", server.UpdateOrderById)
	server.router.DELETE("/orders/:orderId", server.DeleteOrderById)
}
