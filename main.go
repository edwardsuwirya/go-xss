package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/microcosm-cc/bluemonday"
	"net/http"
)

type User struct {
	UserName  string `json:"username"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

func NewUser(user User) *User {
	p := bluemonday.UGCPolicy()
	user.UserName = p.Sanitize(user.UserName)
	user.FirstName = p.Sanitize(user.FirstName)
	user.LastName = p.Sanitize(user.LastName)
	return &user
}

func main() {
	listOfUsers := make([]User, 0)
	apiHost := "localhost"
	apiPort := "8888"
	r := gin.Default()
	route := r.Group("/enigma")

	route.POST("/user", func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT")

		var newUser User
		if err := c.ShouldBindJSON(&newUser); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		listOfUsers = append(listOfUsers, *NewUser(newUser))

		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	route.GET("/user", func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT")

		c.JSON(200, listOfUsers)
	})
	listenAddress := fmt.Sprintf("%s:%s", apiHost, apiPort)
	err := r.Run(listenAddress)
	if err != nil {
		panic(err)
	}
}

// Sample request body
//{
//"username":"jon",
//"firstName":"<button onclick='alert(\"sss\")'>Jhon</button>",
//"lastName":"Key"
//}
