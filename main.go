package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"./controllers"
	"./tools"
	"./models"
	"fmt"
	"./auth"
)
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		    method := c.Request.Method
			c.Header("Access-Control-Allow-Origin", "*")
		    c.Header("Access-Control-Allow-Headers", "Authorization, Content-Type")
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
			c.Header("Access-Control-Allow-Credentials", "true")
			c.Set("content-type", "application/json")

		if method == "OPTIONS" {
			c.JSON(http.StatusOK, "Options Request!")
		}

		c.Next()
	}
}
func main(){
	tools.Init()
	_,err1:=models.InitDB()
	_,err2:=models.InitRedis()
	if err1!=nil||err2!=nil{
		fmt.Println(err1,err2)
		return
	}
	r:=gin.Default()
	r.Use(Cors())
	v1:=r.Group("v1")
	v1.POST("/getRandBookList",controllers.GetRandBookList)
	v1.POST("/getBookList",controllers.GetBookList)
	v1.POST("/getBookInfo",controllers.GetBookInfo)
	v1.POST("/register",controllers.Register)
	v1.POST("/login",controllers.Login)
	v2:=r.Group("v2")
	v2.Use(auth.JWTAuth())
	v2.POST("/getDirectory", controllers.GetDirectory)
	v2.POST("/getPage",controllers.GetPage)
	v2.POST("/getNextPage",controllers.GetNextPage)
	v2.POST("/addBook",controllers.AddBook)
	v2.POST("/delBook",controllers.DelBook)
	v2.POST("/getSelfBook",controllers.GetSelfBook)
	r.Run()

}
