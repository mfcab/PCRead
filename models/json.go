package models
type RequestDirJson struct{
	BookName string `json:"BookName, omitempty"`
}
type RequestPageJson struct{
	BookName string `json:"BookName,omitempty"`
	ChapterName string `json:"ChapterName,omitempty"`
}
type RequestBookRandJson struct{
	BookType string `json:"BookType,omitempty"`
}
type LoginJson struct{
	Phone string `json:"Phone,omitempty"`
	Pwd   string `json:"Pwd,omitempty"`
}
