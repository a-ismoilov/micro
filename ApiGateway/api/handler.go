package api

import (
	"github.com/gin-gonic/gin"
)

type Handler interface {
	Log(ctx *gin.Context)
	PaymentOrder(ctx *gin.Context)
	UserList(ctx *gin.Context)
	LogAdmin(ctx *gin.Context)
	GetOrder(ctx *gin.Context)
	Choose(ctx *gin.Context)
	MealList(ctx *gin.Context)
	OrderList(ctx *gin.Context)
	Cancel(ctx *gin.Context)
	NewMeal(ctx *gin.Context)
	UpdateMeal(ctx *gin.Context)
}
