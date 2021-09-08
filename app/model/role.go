/*
@Time : 2021/9/7 14:31
@Author : zhangyin
@File : role
@Software: GoLand
*/
package model

type RoleStauts struct {
	Name       string
	Sex        string
	Skill      string
	Level      int
	Experience int
	School     string
	Hp         int
	Mp         int
	MpMax      int
	HpMax      int
}

type RoleCreate struct {
	Name      string `json:"name"`
	School_id int    `json:"school_id"`
	Skill_id  int    `json:"skill_Id"`
	Jin       int    `json:"jin"`
	Mu        int    `json:"mu"`
	Shui      int    `json:"shui"`
	Huo       int    `json:"huo"`
	Tu        int    `json:"tu"`
	Sex       int    `json:"sex"`
}
