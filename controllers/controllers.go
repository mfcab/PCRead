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
/*func GetSelfBook(c *gin.Context){
	a:=c.MustGet("claims")
	a.name,redis>selfbook>returnjson({book.id,title,png,}{}{})
}*/
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
/*func GetSelfBook(c *gin.Context){
	a:=c.MustGet("claims")

}*/
func Register(c *gin.Context){
	var reqInfo models.LoginJson
	err:=c.BindJSON(&reqInfo)
	if err!=nil||reqInfo.Phone==""{
		c.JSON(200,gin.H{
			"status" :"-1",
		})
		return
	}
	err=models.RegisterCheck(reqInfo.Phone,reqInfo.Pwd)
	if err!=nil{
		c.JSON(200,gin.H{
			"status" :"-1",
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
			Issuer:    "newtrekWang",                   //签名的发行者
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