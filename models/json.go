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
