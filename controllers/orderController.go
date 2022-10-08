package controllers

import (
	"go-orders-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (inDB *InDB) GetAllOrders(ctx *gin.Context) {
	var orders []models.Order

	if err := inDB.db.Preload("Items").Find(&orders).Error; err != nil {
		resError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	resSuccess(ctx, http.StatusOK, orders)
}

func (inDB *InDB) CreateOrder(ctx *gin.Context) {
	var order models.Order

	if err := ctx.ShouldBindJSON(&order); err != nil {
		resError(ctx, http.StatusBadRequest, extractBindError(err))
		return
	}

	if err := inDB.db.Create(&order).Error; err != nil {
		resError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	resSuccess(ctx, http.StatusCreated, order)
}

func (inDB *InDB) UpdateOrderById(ctx *gin.Context) {
	var order models.Order
	var payloadOrder models.OrderUpdatePayload
	var orderId = ctx.Param("orderId")

	if err := ctx.ShouldBindJSON(&payloadOrder); err != nil {
		resError(ctx, http.StatusBadRequest, extractBindError(err))
		return
	}

	if err := inDB.db.First(&order, orderId).Error; err != nil {
		resError(ctx, http.StatusNotFound, err.Error())
		return
	}

	order.CustomerName = payloadOrder.CustomerName
	order.OrderedAt = payloadOrder.OrderedAt

	inDB.db.Save(&order)

	// for _, item := range payloadOrder.Items {
	// 	inDB.db.Model(&models.Item{}).Where("item_id = ?", item.LineItemID).
	// 		Where("order_id = ?", order.OrderID).
	// 		Updates(item)
	// }

	resSuccess(ctx, http.StatusOK, order)
}

func (inDB *InDB) DeleteOrderById(ctx *gin.Context) {
	var order models.Order
	var orderId = ctx.Param("orderId")

	if err := inDB.db.First(&order, orderId).Error; err != nil {
		resError(ctx, http.StatusNotFound, err.Error())
		return
	}

	inDB.db.Select("Items").Delete(&order)
	resSuccess(ctx, http.StatusOK, order)
}
