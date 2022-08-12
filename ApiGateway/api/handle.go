package api

import (
	_ "github.com/akbarshoh/microOLX/docs"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Online Ordering Micro
// @version         1.0
// @description     This is a new online ordering app with go.

// @host      localhost:8080

func Handle(u UserHandler, o OrderHandler) {
	e := gin.Default()
	user := e.Group("/user")
	user.POST("/log", u.Log)
	user.POST("/log/pay", u.Payment, o.Payment)
	user.DELETE("/log/cancel:id", o.Cancel)
	user.GET("/get-order:id", o.GetOrder)
	user.GET("/menu", o.MealList)
	user.POST("/menu/choose", o.Choose)
	admin := e.Group("/admin")
	admin.POST("/log", u.LogAdmin)
	admin.GET("/user-list", u.UserList)
	admin.POST("/new-meal", o.NewMeal)
	admin.GET("/order-list", o.OrderList)
	admin.PUT("/update-meal", o.UpdateMeal)

	e.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	e.Run(":8080")
}
