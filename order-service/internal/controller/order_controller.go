package controller

import (
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
    "github.com/Butternut01/order-service/internal/entity"
    "github.com/Butternut01/order-service/internal/usecase"
)

type OrderController struct {
    orderUseCase usecase.OrderUseCase
}

func NewOrderController(orderUseCase usecase.OrderUseCase) *OrderController {
    return &OrderController{
        orderUseCase: orderUseCase,
    }
}

func (c *OrderController) CreateOrder(ctx *gin.Context) {
    var order entity.Order
    if err := ctx.ShouldBindJSON(&order); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := c.orderUseCase.CreateOrder(&order); err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusCreated, order)
}

func (c *OrderController) GetOrder(ctx *gin.Context) {
    id := ctx.Param("id")

    order, err := c.orderUseCase.GetOrder(id)
    if err != nil {
        ctx.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
        return
    }

    ctx.JSON(http.StatusOK, order)
}

func (c *OrderController) UpdateOrderStatus(ctx *gin.Context) {
    id := ctx.Param("id")

    var request struct {
        Status entity.OrderStatus `json:"status"`
    }
    if err := ctx.ShouldBindJSON(&request); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := c.orderUseCase.UpdateOrderStatus(id, request.Status); err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, gin.H{"message": "Order status updated"})
}

func (c *OrderController) ListOrders(ctx *gin.Context) {
    var filter entity.OrderFilter

    // Parse query parameters
    if userID := ctx.Query("user_id"); userID != "" {
        filter.UserID = userID
    }
    if status := ctx.Query("status"); status != "" {
        filter.Status = entity.OrderStatus(status)
    }
    if page := ctx.Query("page"); page != "" {
        if val, err := strconv.Atoi(page); err == nil && val > 0 {
            filter.Page = val
        } else {
            filter.Page = 1
        }
    } else {
        filter.Page = 1
    }
    if limit := ctx.Query("limit"); limit != "" {
        if val, err := strconv.Atoi(limit); err == nil && val > 0 {
            filter.Limit = val
        } else {
            filter.Limit = 10
        }
    } else {
        filter.Limit = 10
    }

    orders, err := c.orderUseCase.ListOrders(filter)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, orders)
}