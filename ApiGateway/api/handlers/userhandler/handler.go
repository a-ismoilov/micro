package userhandler

import (
	"fmt"
	"github.com/akbarshoh/microOLX/protos/userproto"
	"github.com/gin-gonic/gin"
	"net/http"
)

type GinUserHandler struct {
	stub userproto.UserServiceClient
}

func New(s userproto.UserServiceClient) GinUserHandler {
	return GinUserHandler{
		stub: s,
	}
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
func (g GinUserHandler) Log(c *gin.Context) {
	a := userproto.User{}
	if err := c.ShouldBindJSON(&a); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	ok, err := g.stub.Log(c, &a)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("%v", err)})
		return
	}
	c.JSON(http.StatusOK, ok)
}

func (g GinUserHandler) Payment(c *gin.Context) {
	p := userproto.PayRequest{}
	if err := c.ShouldBindJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	ok, err := g.stub.Payment(c, &p)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
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
// @Router       /admin/user-list [get]
func (g GinUserHandler) UserList(c *gin.Context) {
	a := userproto.Admin{}
	if err := c.ShouldBindJSON(&a); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	us, err := g.stub.UserList(c, &a)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
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
func (g GinUserHandler) LogAdmin(c *gin.Context) {
	a := userproto.Admin{}
	if err := c.ShouldBindJSON(&a); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("%v", err)})
		return
	}
	ok, err := g.stub.LogAdmin(c, &a)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, ok)
}
