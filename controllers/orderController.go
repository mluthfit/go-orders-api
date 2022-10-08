package controllers

import (
	"go-orders-api/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (server *Server) GetAllOrders(ctx *gin.Context) {
	var order models.Order
	var res = NewRequiredResponse()

	var orders, err = order.GetAllOrdersAndItems(server.db)
	if err != nil {
		res.code = http.StatusInternalServerError
		res.message = err.Error()
		resError(ctx, res, nil)
		return
	}

	res.message = "the data retrieved successfully"
	resSuccess(ctx, res, orders)
}

// func (server *Server) CreateOrder(ctx *gin.Context) {
// 	var order models.Order

// 	if err := ctx.ShouldBindJSON(&order); err != nil {
// 		resError(ctx, http.StatusBadRequest, extractBindError(err))
// 		return
// 	}

// 	if err := server.db.Create(&order).Error; err != nil {
// 		resError(ctx, http.StatusBadRequest, err.Error())
// 		return
// 	}

// 	resSuccess(ctx, http.StatusCreated, order, "the order created successfully")
// }

// func (server *Server) UpdateOrderById(ctx *gin.Context) {
// 	var order models.Order
// 	var payloadOrder models.OrderUpdatePayload
// 	var orderId = ctx.Param("orderId")

// 	if err := ctx.ShouldBindJSON(&payloadOrder); err != nil {
// 		resError(ctx, http.StatusBadRequest, extractBindError(err))
// 		return
// 	}

// 	if err := server.db.First(&order, orderId).Error; err != nil {
// 		resError(ctx, http.StatusNotFound, err.Error())
// 		return
// 	}

// 	order.CustomerName = payloadOrder.CustomerName
// 	order.OrderedAt = payloadOrder.OrderedAt

// 	server.db.Save(&order)

// 	// for _, item := range payloadOrder.Items {
// 	// 	server.db.Model(&models.Item{}).Where("item_id = ?", item.LineItemID).
// 	// 		Where("order_id = ?", order.OrderID).
// 	// 		Updates(item)
// 	// }

// 	resSuccess(ctx, http.StatusOK, order, "the order updated successfully")
// }

func (server *Server) DeleteOrderById(ctx *gin.Context) {
	var order *models.Order
	var orderId = ctx.Param("orderId")

	var res = NewRequiredResponse()
	var parseOrderId, err = strconv.ParseUint(orderId, 10, 32)
	if err != nil {
		res.code = http.StatusBadRequest
		res.message = "order id must be unsigned integer"
		resError(ctx, res, nil)
		return
	}

	order, err = order.DeleteOrderAndItems(server.db, uint(parseOrderId))
	if err != nil {
		res.code = http.StatusNotFound
		res.message = err.Error()
		resError(ctx, res, nil)
		return
	}

	res.message = "the order deleted successfully"
	resSuccess(ctx, res, order)
}
