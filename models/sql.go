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
		err:=DB.Select("id,name,author").Limit(100).Find(&bookList).Error
		return bookList,err
	}
	err:=DB.Select("id,name,author").Where("type=?",s).Limit(100).Find(&bookList).Error
	return bookList,err
}
func GetSelfBook(phone string)([]*BookInfo,error){
	var bookList []*BookInfo
	list,err:=Clint.SMembers(phone).Result()
	if err!=nil{
		return nil,err
	}
	for _,name:=range list{
			if name==""{
				continue
			}
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
	a,err:=Clint.Exists(phone).Result()
	if err!=nil||a!=0{
		return errors.New("请直接登陆")
	}
	code,err:=GetCheckCode(phone)
	if err!=nil{
		return err
	}
	if pwd!=code{
		return errors.New("验证码错误")
	}
	err=Clint.SAdd(phone,"帝霸").Err()
	if err!=nil{
		return err
	}
return nil
}
func LoginCheck(phone string, pwd string) error{

	a,err:=Clint.Exists(phone).Result()
	if err!=nil||a!=1{
		return errors.New("Wrong")
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
func AddBook(book string,phone string)  error{
	a,err:=Clint.Exists(phone).Result()
	if err!=nil||a!=1{
		return errors.New("Wrong")
	}
	err=Clint.SAdd(phone,book).Err()
	if err!=nil{
		return err
	}
	return nil
}
func DeleteBook(book string,phone string) error{
	a,err:=Clint.Exists(phone).Result()
	if err!=nil||a!=1{
		return errors.New("Wrong")
	}
	isMember, err := Clint.SIsMember(phone, book).Result()
	if err!=nil||!isMember{
		return errors.New("Wrong")
	}
	err=Clint.SRem(phone,book).Err()
	if err!=nil{
		return err
	}
	return nil
}
func SearchBook(name string) ([]*BookInfo,error){
	var bookList []*BookInfo
	name2:=fmt.Sprintf("%%%s%%",name)
	fmt.Println(name2)
	err:=DB.Select("id","name").Where("name LIKE ?", name2).Find(bookList).Error
	if err!=nil{
		return nil,err
	}
	return bookList,nil
}