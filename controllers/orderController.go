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

func (server *Server) CreateOrder(ctx *gin.Context) {
	var order *models.Order
	var err error
	var res = NewRequiredResponse()

	if err := ctx.ShouldBindJSON(&order); err != nil {
		res.code = http.StatusBadRequest
		res.message = "request validation errors"
		resError(ctx, res, extractBindError(err))
		return
	}

	order, err = order.CreateOrderAndItems(server.db)
	if err != nil {
		res.code = http.StatusBadRequest
		res.message = err.Error()
		resError(ctx, res, nil)
		return
	}

	res.code = http.StatusCreated
	res.message = "the order created successfully"
	resSuccess(ctx, res, order)
}

func (server *Server) UpdateOrderById(ctx *gin.Context) {
	var order *models.Order
	var payloadOrder models.OrderUpdatePayload
	var orderId = ctx.Param("orderId")

	var res = NewRequiredResponse()
	var parseOrderId, err = strconv.ParseUint(orderId, 10, 32)
	if err != nil {
		res.code = http.StatusBadRequest
		res.message = "order id must be unsigned integer"
		resError(ctx, res, nil)
		return
	}

	if err := ctx.ShouldBindJSON(&payloadOrder); err != nil {
		res.code = http.StatusBadRequest
		res.message = "request validation errors"
		resError(ctx, res, extractBindError(err))
		return
	}

	order, err = order.UpdateOrderAndItems(server.db, uint(parseOrderId), payloadOrder)
	if err != nil {
		res.code = http.StatusNotFound
		res.message = err.Error()
		resError(ctx, res, nil)
		return
	}

	res.message = "the order updated successfully"
	resSuccess(ctx, res, order)
}

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
