/*
@Time : 2021/9/7 14:31
@Author : zhangyin
@File : role
@Software: GoLand
*/
package model

type RoleStauts struct {
	Arm           string
	Name          string
	Level         int
	Experience    int
	School        string
	Jin           int
	Mu            int
	Shui          int
	Huo           int
	Tu            int
	Sex           int
	Skill         string
	MaxMp         int
	Mp            int
	MaxHp         int
	Hp            int
	HpBak         int
	MpBak         int
	SkillCD       int
	Attack        int //攻击
	Defense       int //防御
	ViolentRate   int //暴击几率
	ViolentAttack int //暴击伤害
}

type RoleCreate struct {
	Name      string `json:"name"`
	School_id int    `json:"school_id"`
	Skill_id  int    `json:"skill_Id"`
	Arm_id    int    `json:"arm_id"`
	Jin       int    `json:"jin"`
	Mu        int    `json:"mu"`
	Shui      int    `json:"shui"`
	Huo       int    `json:"huo"`
	Tu        int    `json:"tu"`
	Sex       int    `json:"sex"`
}
