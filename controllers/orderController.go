package controllers

import (
	"go-orders-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (server *Server) GetAllOrders(ctx *gin.Context) {
	var orders []models.Order

	if err := server.db.Preload("Items").Find(&orders).Error; err != nil {
		resError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	resSuccess(ctx, http.StatusOK, orders, "the data retrieved successfully")
}

func (server *Server) CreateOrder(ctx *gin.Context) {
	var order models.Order

	if err := ctx.ShouldBindJSON(&order); err != nil {
		resError(ctx, http.StatusBadRequest, extractBindError(err))
		return
	}

	if err := server.db.Create(&order).Error; err != nil {
		resError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	resSuccess(ctx, http.StatusCreated, order, "the order created successfully")
}

func (server *Server) UpdateOrderById(ctx *gin.Context) {
	var order models.Order
	var payloadOrder models.OrderUpdatePayload
	var orderId = ctx.Param("orderId")

	if err := ctx.ShouldBindJSON(&payloadOrder); err != nil {
		resError(ctx, http.StatusBadRequest, extractBindError(err))
		return
	}

	if err := server.db.First(&order, orderId).Error; err != nil {
		resError(ctx, http.StatusNotFound, err.Error())
		return
	}

	order.CustomerName = payloadOrder.CustomerName
	order.OrderedAt = payloadOrder.OrderedAt

	server.db.Save(&order)

	// for _, item := range payloadOrder.Items {
	// 	server.db.Model(&models.Item{}).Where("item_id = ?", item.LineItemID).
	// 		Where("order_id = ?", order.OrderID).
	// 		Updates(item)
	// }

	resSuccess(ctx, http.StatusOK, order, "the order updated successfully")
}

func (server *Server) DeleteOrderById(ctx *gin.Context) {
	var order models.Order
	var orderId = ctx.Param("orderId")

	if err := server.db.First(&order, orderId).Error; err != nil {
		resError(ctx, http.StatusNotFound, err.Error())
		return
	}

	server.db.Select("Items").Delete(&order)
	resSuccess(ctx, http.StatusOK, order, "the order deleted successfully")
}
