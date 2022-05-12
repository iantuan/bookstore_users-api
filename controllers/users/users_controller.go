package users

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iantuan/bookstore_users-api/domain/users"
	"github.com/iantuan/bookstore_users-api/services"
	"github.com/iantuan/bookstore_users-api/utils/errors"
)

var (
	counter int
)

func CreateUser(c *gin.Context) {

	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		fmt.Println(err)
		return
	}

	//fmt.Println(user)
	//bytes, err := ioutil.ReadAll(c.Request.Body)
	//if err != nil {
	//	//TODO: Handle error
	//	return
	//}
	//if err := json.Unmarshal(bytes, &user); err != nil {
	//	fmt.Println("json unmarshal error")
	//	fmt.Println(err.Error())
	//	//fmt.Println(string(bytes))
	//	//TODO: Handle json error
	//	return
	//}

	result, saveErr := services.CreateUser(user)

	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}
	//fmt.Println(result)
	c.JSON(http.StatusCreated, result)
}

func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement GetUser!")
}

//func SearchUser(c *gin.Context) {
//	c.String(http.StatusNotImplemented, "implement me!")
//}
