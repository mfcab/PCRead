package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"./controllers"
	"./tools"
	"./models"
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
	models.InitDB()
	r:=gin.Default()
	r.Use(Cors())
	v1:=r.Group("v1")
	v1.POST("/getRandBookList",controllers.GetRandBookList)
	v1.POST("getBookList",controllers.GetBookList)
	v2:=r.Group("v2")
	v1.GET("/hello")
	v2.POST("/getDirectory", controllers.GetDirectory)
	v2.POST("/getPage",controllers.GetPage)
	v2.POST("/getNextPage",controllers.GetNextPage)
	r.Run()

}
