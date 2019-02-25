package controllers

import (
	"github.com/gin-gonic/gin"
	"../models"
	"../tools"
	"fmt"
)

func GetDirectory(c *gin.Context){
	var reqInfo models.RequestDirJson
	err:=c.BindJSON(&reqInfo)
	if err!=nil{
		c.JSON(404,gin.H{})
		return
	}
	if reqInfo.BookName==""{
		c.JSON(404,gin.H{})
		return
	}
	a,err:=tools.GetDirectory(reqInfo.BookName)
	if err!=nil{
		c.JSON(404,gin.H{})
		return
	}
	c.JSON(200, gin.H{
		"List": a,
})
	return
}
func GetPage(c *gin.Context){
	var reqInfo models.RequestPageJson
	err:=c.BindJSON(&reqInfo)
	if err!=nil{
		fmt.Println(err)
		c.JSON(404,gin.H{})
		return
	}
	a,err:=tools.GetPage(reqInfo.BookName,reqInfo.ChapterName)
	if err!=nil{
		fmt.Println(err)
		c.JSON(404,gin.H{})
		return
	}
	c.JSON(200, gin.H{
		"text": a,
	})
	return
}
func GetNextPage(c *gin.Context){
	var reqInfo models.RequestPageJson
	err:=c.BindJSON(&reqInfo)
	if err!=nil{
		c.JSON(404,gin.H{})
		return
	}
	text,title,err:=tools.GetNextPage(reqInfo.BookName,reqInfo.ChapterName)
	if err!=nil{
		c.JSON(404,gin.H{})
		return
	}
	c.JSON(200, gin.H{
		"text": text,
		"title": title,
	})
	return
}
/*func GetSelfBook(c *gin.Context){
	a:=c.MustGet("claims")
	a.name,redis>selfbook>returnjson({book.id,title,png,}{}{})
}*/
func GetRandBookList(c *gin.Context){
	var reqInfo models.RequestBookRandJson
	err:=c.BindJSON(&reqInfo)
	if err!=nil{
		c.JSON(404,gin.H{})
		return
	}
	a,err:=models.GetRandBookList(reqInfo.BookType)
	if err!=nil{
		c.JSON(404,gin.H{})
		return
	}
	c.JSON(200, gin.H{
		"bookList": a,
	})
	return
}
func GetBookList(c *gin.Context){
	var reqInfo models.RequestBookRandJson
	err:=c.BindJSON(&reqInfo)
	if err!=nil{
		c.JSON(404,gin.H{})
		return
	}
	a,err:=models.GetBookList(reqInfo.BookType)
	if err!=nil{
		c.JSON(404,gin.H{})
		return
	}
	c.JSON(200, gin.H{
		"bookList": a,
	})
	return
}
/*func GetSelfBook(c *gin.Context){
	a:=c.MustGet("claims")

}*/