package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
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
func InitDB() (*gorm.DB, error) {
	db, err := gorm.Open("mysql", "readuser:p@s#0fPCR@/Blog?charset=utf8&parseTime=True&loc=Local")
	if err == nil {
		db.LogMode(true)
		db.SingularTable(true)
		DB=db
		return db, err
	}
	return nil, err
}
func GetRandBookList( s string) ([6]*BookInfo,error){
	var bookList [6]*BookInfo
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