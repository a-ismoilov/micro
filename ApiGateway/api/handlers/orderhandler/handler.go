package orderhandler

import (
	"github.com/akbarshoh/microOLX/protos/orderproto"
	_ "github.com/akbarshoh/microOLX/protos/orderproto"
	"github.com/akbarshoh/microOLX/protos/userproto"
	_ "github.com/akbarshoh/microOLX/protos/userproto"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type Error struct {
	err string "error"
	def error
}

type OrderStub struct {
}

type GinOrderHandler struct {
	stub orderproto.OrderServiceClient
}

func New(s orderproto.OrderServiceClient) GinOrderHandler {
	return GinOrderHandler{
		stub: s,
	}
}

// GetOrder
// @Summary      GetOrder
// @Description  get order by order_id
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        id  path  string true "order id"
// @Success      200 {object} orderproto.Order
// @Failure      400 {object} map[string]any
// @Failure      500 {object} map[string]any
// @Router       /user/get-order:id [get]
func (g GinOrderHandler) GetOrder(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	o, err := g.stub.GetOrder(c, &orderproto.Request{Id: int32(id)})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, o)
}

// Choose
// @Summary      Choose
// @Description  user chooses meal
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        meal_id  body  orderproto.MealChoice true "meal id and order id"
// @Success      200 {object} orderproto.OK
// @Failure      400 {object} map[string]any
// @Failure      500 {object} map[string]any
// @Router       /user/menu/choose [post]
func (g GinOrderHandler) Choose(c *gin.Context) {
	o := orderproto.MealChoice{}
	if err := c.ShouldBindJSON(&o); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	ok, err := g.stub.Choose(c, &o)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, ok)
}

// Payment
// @Summary      Payment
// @Description  user pays his order
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        order_id  body  userproto.PayRequest true "user id and order id"
// @Success      200 {object} orderproto.OK
// @Failure      400 {object} map[string]any
// @Failure      500 {object} map[string]any
// @Router       /user/pay [post]
func (g GinOrderHandler) Payment(c *gin.Context) {
	p := userproto.PayRequest{}
	if err := c.ShouldBindJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	ok, err := g.stub.Payment(c, &orderproto.Request{Id: p.OrderId})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, ok)
}

// MealList
// @Summary      MealList
// @Description  shows menu to user
// @Tags         user
// @Accept       json
// @Produce      json
// @Success      200 {object} orderproto.Meals
// @Failure      400 {object} map[string]any
// @Failure      500 {object} map[string]any
// @Router       /user/menu [get]
func (g GinOrderHandler) MealList(c *gin.Context) {
	ms, err := g.stub.MealList(c, &orderproto.Request{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, ms)
}

// OrderList
// @Summary      OrderList
// @Description  shows orders to admin
// @Tags         admin
// @Accept       json
// @Produce      json
// @Success      200 {object} orderproto.Orders
// @Failure      400 {object} map[string]any
// @Failure      500 {object} map[string]any
// @Router       /admin/order-list [get]
func (g GinOrderHandler) OrderList(c *gin.Context) {
	a := userproto.Admin{}
	if err := c.ShouldBindJSON(&a); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	os, err := g.stub.OrderList(c, &orderproto.Request{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, os)
}

// Cancel
// @Summary      Cancel
// @Description  cancels user order
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        order_id  path  string true "order id"
// @Success      200 {object} orderproto.OK
// @Failure      400 {object} map[string]any
// @Failure      500 {object} map[string]any
// @Router       /user/log/cancel:id [delete]
func (g GinOrderHandler) Cancel(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	ok, err := g.stub.Cancel(c, &orderproto.Request{Id: int32(id)})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, ok)
}

// NewMeal
// @Summary      New
// @Description  admin adds new meal
// @Tags         admin
// @Accept       json
// @Produce      json
// @Param        meal_struct  body  orderproto.Meal true "meal struct"
// @Success      200 {object} orderproto.OK
// @Failure      400 {object} map[string]any
// @Failure      500 {object} map[string]any
// @Router       /admin/new-meal [post]
func (g GinOrderHandler) NewMeal(c *gin.Context) {
	m := orderproto.Meal{}
	if err := c.ShouldBindJSON(&m); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	ok, err := g.stub.NewMeal(c, &m)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, ok)
}

// UpdateMeal
// @Summary      Update
// @Description  admin update existing meal
// @Tags         admin
// @Accept       json
// @Produce      json
// @Param        meal struct  body  orderproto.Meal true "meal struct"
// @Success      200 {object} orderproto.OK
// @Failure      400 {object} map[string]any
// @Failure      500 {object} map[string]any
// @Router       /admin/update-meal [put]
func (g GinOrderHandler) UpdateMeal(c *gin.Context) {
	m := orderproto.Meal{}
	if err := c.ShouldBindJSON(&m); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	ok, err := g.stub.UpdateMeal(c, &m)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, ok)
}
