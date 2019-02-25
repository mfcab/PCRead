package models
type RequestDirJson struct{
	BookName string `json:"BookName"`
}
type RequestPageJson struct{
	BookName string `json:"BookName"`
	ChapterName string `json:"ChapterName"`
}
type RequestBookRandJson struct{
	BookType string `json:"BookType"`
}
