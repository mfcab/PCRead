package models
type RequestDirJson struct{
	BookName string `json:"BookName,omitempy"`
}
type RequestPageJson struct{
	BookName string `json:"BookName,omitempy"`
	ChapterName string `json:"ChapterName,omitempy"`
}
type RequestBookRandJson struct{
	BookType string `json:"BookType,omitempy"`
}
type LoginJson struct{
	Phone string `json:"Phone,omitempy"`
	Pwd   string `json:"Pwd,omitempy"`
}
