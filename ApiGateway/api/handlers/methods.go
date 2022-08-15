package handlers

import (
	"fmt"
	"github.com/akbarshoh/microOLX/protos/orderproto"
	"github.com/akbarshoh/microOLX/protos/userproto"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type Handler struct {
	OS orderproto.OrderServiceClient
	US userproto.UserServiceClient
}

func New(os orderproto.OrderServiceClient, us userproto.UserServiceClient) Handler {
	return Handler{
		OS: os,
		US: us,
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
// @Router       /user/get-order/{id} [get]
func (h Handler) GetOrder(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("%v, %s", err, c.Param("id"))})
		return
	}
	o, err := h.OS.GetOrder(c, &orderproto.Request{Id: int32(id)})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("%v", err)})
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
func (h Handler) Choose(c *gin.Context) {
	o := orderproto.MealChoice{}
	if err := c.ShouldBindJSON(&o); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("%v", fmt.Sprintf("%v", err))})
		return
	}
	ok, err := h.OS.Choose(c, &o)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("%v", fmt.Sprintf("%v", err))})
		return
	}
	c.JSON(http.StatusOK, ok)
}

// PaymentOrder
// @Summary      Payment
// @Description  user pays his order
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        order_id  body  userproto.PayRequest true "user id and order id"
// @Success      200 {object} orderproto.OK
// @Failure      400
// @Failure      500
// @Router       /user/log/pay [post]
func (h Handler) PaymentOrder(c *gin.Context) {
	p := userproto.PayRequest{}
	if err := c.ShouldBindJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("%v", fmt.Sprintf("%v", err))})
	}
	okp, err := h.OS.Payment(c, &orderproto.Request{Id: p.OrderId})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("%v", fmt.Sprintf("%v", err))})
	}
	//c.JSON(http.StatusOK, ok)
	p.Price = okp.OK
	ok, err := h.US.Payment(c, &p)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("%v, 3", err)})
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
func (h Handler) MealList(c *gin.Context) {
	ms, err := h.OS.MealList(c, &orderproto.Request{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("%v", err)})
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
func (h Handler) OrderList(c *gin.Context) {
	os, err := h.OS.OrderList(c, &orderproto.Request{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("%v, 2", err)})
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
// @Router       /user/log/cancel/:id [delete]
func (h Handler) Cancel(c *gin.Context) {
	//id, err := strconv.Atoi(c.Param("id"))
	//if err != nil {
	//	c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("%v", err)})
	//	return
	//}
	ok, err := h.OS.Cancel(c, &orderproto.Request{Id: 4})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("%v", err)})
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
func (h Handler) NewMeal(c *gin.Context) {
	m := orderproto.Meal{}
	if err := c.ShouldBindJSON(&m); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("%v", fmt.Sprintf("%v", err))})
		return
	}
	ok, err := h.OS.NewMeal(c, &m)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("%v", fmt.Sprintf("%v", err))})
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
func (h Handler) UpdateMeal(c *gin.Context) {
	m := orderproto.Meal{}
	if err := c.ShouldBindJSON(&m); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("%v", err)})
		return
	}
	ok, err := h.OS.UpdateMeal(c, &m)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("%v", err)})
		return
	}
	c.JSON(http.StatusOK, ok)
}

// Log
// @Summary      UserLogin
// @Description  user should log to get access
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        user  body  userproto.User true "user struct"
// @Success      200 {object} userproto.Response
// @Failure      400
// @Failure      500
// @Router       /user/log [post]
func (h Handler) Log(c *gin.Context) {
	a := userproto.User{}
	if err := c.ShouldBindJSON(&a); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	ok, err := h.US.Log(c, &a)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("%v", err)})
		return
	}
	c.JSON(http.StatusOK, ok)
}

// UserList
// @Summary      users
// @Description  admin gets users list
// @Tags         admin
// @Accept       json
// @Produce      json
// @Param        admin body userproto.Admin true "admin with id"
// @Success      200 {object} userproto.Users
// @Failure      400 {object} map[string]string
// @Failure      500 {object} map[string]string
// @Router       /admin/user-list [post]
func (h Handler) UserList(c *gin.Context) {
	a := userproto.Admin{}
	if err := c.ShouldBindJSON(&a); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("%v", err)})
		return
	}
	us, err := h.US.UserList(c, &a)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("%v", err)})
		return
	}
	c.JSON(http.StatusOK, us)
}

// LogAdmin
// @Summary      login
// @Description  admin should log to get access
// @Tags         admin
// @Accept       json
// @Produce      json
// @Param        password&id  body  userproto.Admin true "admin password and id"
// @Success      200 {object} userproto.Response
// @Failure      400 {object} map[string]string
// @Failure      500 {object} map[string]string
// @Router       /admin/log [post]
func (h Handler) LogAdmin(c *gin.Context) {
	a := userproto.Admin{}
	if err := c.ShouldBindJSON(&a); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("%v", err)})
		return
	}
	ok, err := h.US.LogAdmin(c, &a)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("%v", err)})
		return
	}
	c.JSON(http.StatusOK, ok)
}
