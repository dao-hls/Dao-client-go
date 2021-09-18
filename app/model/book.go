/*
@Time : 2021/9/15 15:54
@Author : zhangyin
@File : book
@Software: GoLand
*/
package model

type Book struct {
	Id   int
	Name string
	//Description string
}
type RespBooks struct {
	Code  int
	Books [5]Book
	Len   int
}
type RespReading struct {
	BookName string
	Code     int
	BookId   int
}

type SetReading struct {
	Id int
}
type BookPass struct {
	Pass int
}
