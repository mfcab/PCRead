package controllers

import (
	"github.com/gin-gonic/gin"
	"../models"
	"../tools"
	"fmt"
	"time"
	"net/http"
	"../auth"
	jwtgo "github.com/dgrijalva/jwt-go"
)

func GetDirectory(c *gin.Context){
	claims := c.MustGet("claims").(*auth.CustomClaims)
	if claims == nil {
		c.JSON(404,gin.H{"status":"-1","msg":"token错误"})
		return
	}
	var reqInfo models.RequestDirJson
	err:=c.BindJSON(&reqInfo)
	if err!=nil{
		c.JSON(404,gin.H{"status":"-1","msg":"解析错误"})
		return
	}
	if reqInfo.BookName==""{
		c.JSON(404,gin.H{"status":"-1","msg":"信息有误"})
		return
	}
	a,err:=tools.GetDirectory(reqInfo.BookName)
	if err!=nil{
		c.JSON(404,gin.H{"status":"-1","msg":"信息有误。"})
		return
	}
	c.JSON(200, gin.H{
		"List": a,
})
	return
}
func GetPage(c *gin.Context){
	claims := c.MustGet("claims").(*auth.CustomClaims)
	if claims == nil {
		c.JSON(404,gin.H{"status":"-1","msg":"token错误"})
		return
	}
	var reqInfo models.RequestPageJson
	err:=c.BindJSON(&reqInfo)
	if err!=nil{
		fmt.Println(err)
		c.JSON(404,gin.H{"status":"-1","msg":"解析错误"})
		return
	}
	a,err:=tools.GetPage(reqInfo.BookName,reqInfo.ChapterName)
	if err!=nil{
		fmt.Println(err)
		c.JSON(404,gin.H{"status":"-1","msg":"信息有误"})
		return
	}
	c.JSON(200, gin.H{
		"text": a,
	})
	return
}
func GetNextPage(c *gin.Context){
	claims := c.MustGet("claims").(*auth.CustomClaims)
	if claims == nil {
		c.JSON(404,gin.H{"status":"-1","msg":"token错误"})
		return
	}
	var reqInfo models.RequestPageJson
	err:=c.BindJSON(&reqInfo)
	if err!=nil||reqInfo.BookName==""{
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
func GetSelfBook(c *gin.Context){
	claims := c.MustGet("claims").(*auth.CustomClaims)
	if claims == nil||claims.Phone=="" {
		c.JSON(404,gin.H{"status":"-1","msg":"token错误"})
		return
	}
	list,err:=models.GetSelfBook(claims.Phone)
	if err!=nil {
		c.JSON(404, gin.H{"status": "-1", "msg": "查询错误"})
		return
	}
	c.JSON(200, gin.H{
		"user":claims.Phone,
		"bookList": list,
	})
	return

}
func GetRandBookList(c *gin.Context){
	var reqInfo models.RequestBookRandJson
	err:=c.BindJSON(&reqInfo)
	if err!=nil||reqInfo.BookType==""{
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
	if err!=nil||reqInfo.BookType==""{
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
func GetBookInfo(c *gin.Context){
	var reqInfo models.RequestPageJson
	err:=c.BindJSON(&reqInfo)
	if err!=nil||reqInfo.BookName==""{
		fmt.Println(err)
		c.JSON(404,gin.H{"status":"-1","msg":"解析错误"})
		return
	}
	var bookInfo models.BookInfo
	err=bookInfo.GetBookInfo(reqInfo.BookName)
	if err!=nil{
		c.JSON(404,gin.H{})
		return
	}
	c.JSON(200, gin.H{
		"book": bookInfo,
	})
	return
}
func Register(c *gin.Context){
	var reqInfo models.LoginJson
	err:=c.BindJSON(&reqInfo)
	if err!=nil||reqInfo.Phone==""{
		c.JSON(200,gin.H{
			"status" :"-1",
			"err":"请求有问题",
		})
		return
	}
	err=models.RegisterCheck(reqInfo.Phone,reqInfo.Pwd)
	if err!=nil{
		c.JSON(200,gin.H{
			"status" :"-1",
			"err":err.Error(),
		})
		return
	}
	c.JSON(200,gin.H{
		"status" :"0",
	})
	return

}
func Login(c *gin.Context){
	var reqInfo models.LoginJson
	err:=c.BindJSON(&reqInfo)
	fmt.Println(reqInfo.Phone)
	if err!=nil||reqInfo.Phone==""{
		c.JSON(200,gin.H{
			"status" :"-1",
		})
		return
	}
	err=models.LoginCheck(reqInfo.Phone,reqInfo.Pwd)
	if err!=nil{
		c.JSON(200,gin.H{
			"status" :"-1",
		})
		return
	}
	j := auth.NewJWT()
	claims := auth.CustomClaims{
		reqInfo.Phone,
		jwtgo.StandardClaims{
			NotBefore: int64(time.Now().Unix() - 1000), // 签名生效时间
			ExpiresAt: int64(time.Now().Unix() + 360000), // 过期时间 一小时
			Issuer:    "GPC",                   //签名的发行者
		},
	}

	token, err := j.CreateToken(claims)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": -1,
			"msg":    err.Error(),
		})
		return
	}
	c.JSON(200,gin.H{
		"status" :"0",
		"token" : token,
	})
	return
}
func AddBook(c *gin.Context){
	claims := c.MustGet("claims").(*auth.CustomClaims)
	if claims == nil||claims.Phone=="" {
		c.JSON(404,gin.H{"status":"-1","msg":"token错误"})
		return
	}
	var reqInfo models.RequestDirJson
	err:=c.BindJSON(&reqInfo)
	if err!=nil||reqInfo.BookName==""{
		c.JSON(404,gin.H{"status":"-1","msg":"解析错误"})
		return
	}
	err=models.AddBook(reqInfo.BookName,claims.Phone)
	if err!=nil{
		c.JSON(404,gin.H{"status":"-1","msg":"添加错误"})
		return
	}
	c.JSON(200,gin.H{"status":"0","msg":"添加成功"})
	return
}
func DelBook(c *gin.Context){
	claims := c.MustGet("claims").(*auth.CustomClaims)
	if claims == nil||claims.Phone=="" {
		c.JSON(404,gin.H{"status":"-1","msg":"token错误"})
		return
	}
	var reqInfo models.RequestDirJson
	err:=c.BindJSON(&reqInfo)
	if err!=nil||reqInfo.BookName==""{
		c.JSON(404,gin.H{"status":"-1","msg":"解析错误"})
		return
	}
	err=models.DeleteBook(reqInfo.BookName,claims.Phone)
	if err!=nil{
		c.JSON(404,gin.H{"status":"-1","msg":"删除错误"})
		return
	}
	c.JSON(200,gin.H{"status":"0","msg":"删除成功"})
	return
}
