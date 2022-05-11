package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	counter int
)

func CreateUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement CreateUser!")
}

func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement GetUser!")
}

//func SearchUser(c *gin.Context) {
//	c.String(http.StatusNotImplemented, "implement me!")
//}
