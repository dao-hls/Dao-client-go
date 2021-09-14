/*
@Time : 2021/9/14 15:46
@Author : zhangyin
@File : event
@Software: GoLand
*/
package model

type Event struct {
	Id          int
	Name        string
	Description string
}
type FinishEvent struct {
	Id int
}
