package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/go-redis/redis"
	"errors"
	"fmt"
)
type BookInfo struct{
	Id int
	Name string
	Author  string
	Info string
	Png string
	Path string
	Type string
	Dn int
}
var DB *gorm.DB
var Clint *redis.Client
func InitDB() (*gorm.DB, error) {
	db, err := gorm.Open("mysql", "readuser:p@s#0fPCR@/PCRead?charset=utf8&parseTime=True&loc=Local")
	if err == nil {
		db.LogMode(true)
		db.SingularTable(true)
		DB=db
		return db, err
	}
	return nil, err
}
func InitRedis()(*redis.Client,error){
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
	})
	_, err := client.Ping().Result()
	if err!=nil{
		return nil,err
	}
	Clint=client
	return client,err
}
func GetRandBookList( s string) ([]*BookInfo,error){
	var bookList []*BookInfo
	if s=="热门图书"{
		err:=DB.Select("id,name,author,png").Order("rand()").Limit(6).Find(&bookList).Error
		return bookList,err
	}
	err:=DB.Select("id,name,author,png").Where("type=?",s).Order("rand()").Limit(6).Find(&bookList).Error
	return bookList,err
}
func GetBookList(s string) ([]*BookInfo,error){
	var bookList []*BookInfo
	if s=="热门图书"{
		err:=DB.Select("id,name,author,png").Limit(100).Find(&bookList).Error
		return bookList,err
	}
	err:=DB.Select("id,name,author,png").Where("type=?",s).Limit(100).Find(&bookList).Error
	return bookList,err
}
func GetSelfBook(list []string)([]*BookInfo,error){
	var bookList []*BookInfo
	for _,name:=range list{
			var book BookInfo
			err:=DB.Select("id,name,png").Where("name=?",name).Find(&book).Error
			if err!=nil{
				return nil,err
			}
			bookList=append(bookList,&book)
	}
	return bookList,nil
}
func (book *BookInfo) GetBookInfo(s string) error{
	err:=DB.Select("id,name,author,info,type,png,dn").Where("name=?",s).Find(book).Error
	return err
}
func RegisterCheck(phone string, pwd string) error{
	code,err:=GetCheckCode(phone)
	if err!=nil{
		return err
	}
	if pwd!=code{
		return errors.New("验证码错误")
	}
	err=Clint.Set(phone,"",0).Err()
	if err!=nil{
		return err
	}
return nil
}
func LoginCheck(phone string, pwd string) error{

	a:=Clint.Exists(phone)
	err1:=a.String()
	fmt.Println("this,is a",err1)
	b,err:=a.Result()
	fmt.Println(a.Err(),a.Name(),a.Val(),a.Args(),b,err)
	err2:=a.Err()
	fmt.Println(err2)
	if err2!=nil{
		fmt.Println(err2)
		return err2
	}
	code,_:=GetCheckCode(phone)
	if pwd!=code{
		return errors.New("验证码错误")
	}

	return nil
}
func GetCheckCode(phone string) (string,error){
	//短信网关
	code:=phone
	code="123456"
	return code,nil
}
