package tools

import (
	"os"
	"bufio"
	"io"
	"regexp"
	"errors"
	"fmt"
)
var Reg *regexp.Regexp
func Init(){
	Reg=regexp.MustCompile(`^第.{1,5}[章回节苏].{1,16}`)
}
func GetDirectory(name string) ([]string,error) {
	a := []string{}
	text, err := os.Open("/root/Book/"+name)
	if err != nil {
		return a, err
	}
	defer text.Close()
	r := bufio.NewReader(text)
	for {
		s, _, err := r.ReadLine()
		if err != nil && err == io.EOF {
			return a, nil
		}
		str:=string(s)
		Title:=Reg.FindString(str)
		if Title==""{
			continue
		}
		a=append(a,Title)
	}
}

func GetPage(name string,chapter string) (string,error){
	page:=""
	tag:=false
	if name==""{
		return page, errors.New("XXX")
	}
	if chapter==""{
		tag=true
	}
	reg2:=regexp.MustCompile("^"+chapter)
	text, err := os.Open("/root/Book/"+name)
	if err != nil {
		return page, err
	}
	defer text.Close()
	r := bufio.NewReader(text)
	for {
		s,_, err := r.ReadLine()
		if err != nil&&err==io.EOF {
			fmt.Println(err)
			return page,nil
		}
		str:=string(s)
		if !tag{
			a:=reg2.MatchString(str)
			if !a{
				continue
			} else {
				tag=true
				page=page+str+"\n"
				continue

			}}else {
			nextTitle:=Reg.MatchString(str)
			if !nextTitle{
				page=page+str+"\n"
				continue
			} else {
				return page,nil
			}
		}

	}
	return page,nil
}
func GetNextPage(name string,chapter string) (string,string,error){
	nextChapter:=""
	findTag:=false
	page:=""
	tag:=false
	if name==""{
		return page,nextChapter, errors.New("XXX")
	}
	text, err := os.Open("/root/Book/"+name)
	if err != nil {
		return page,nextChapter, err
	}
	defer text.Close()
	r := bufio.NewReader(text)
	if chapter==""{
		for {
			s,_, err := r.ReadLine()
			if err != nil&&err==io.EOF {
				fmt.Println(err)
				return page,nextChapter,nil
			}
			str:=string(s)
			if !tag{
				a:=Reg.MatchString(str)
				if !a{
					continue
				} else {
					tag=true
					nextChapter=str
					page=page+str+"\n"
					continue

				}}else {
				nextTitle:=Reg.MatchString(str)
				if !nextTitle{
					page=page+str+"\n"
					continue
				} else {
					return page,nextChapter,nil
				}
			}

		}
	} else {
		reg2:=regexp.MustCompile("^"+chapter)
		for {
			s, _, err := r.ReadLine()
			if err != nil && err == io.EOF {
				fmt.Println(err)
				return page, nextChapter, nil
			}
			str := string(s)
			if !tag {
				if !findTag {
					a := reg2.MatchString(str)
					if a {
						findTag = true
					}
					continue
				} else {
					b := Reg.MatchString(str)
					if !b {
						continue
					} else {
						tag = true
						nextChapter = str
						page = page + str + "\n"
						continue

					}
				}
			} else {
				nextTitle := Reg.MatchString(str)
				if !nextTitle {
					page = page + str + "\n"
					continue
				} else {
					return page, nextChapter, nil
				}

			}
		}
		}
	return page,nextChapter,nil
	}
