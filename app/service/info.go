/*
@Time : 2021/8/3 4:44 下午
@Author : fushisanlang
@File : info.go
@Software: GoLand
*/
package service

import (
	"Dao-client/app/tools"
	"fmt"
)

//定义url
var UrlPre string

func SelectServerAdress() {
	var serverId int

	fmt.Println("服务器列表")
	//fmt.Println("0.潜龙在渊-内测服务器")
	fmt.Println("1.飞龙在天-大陆服务器")
	fmt.Println("2.龙战于野-国际服务器")
	fmt.Println("请输入服务器序号选择服务器")
	fmt.Scan(&serverId)
	tools.Clean()
	if serverId == 0 {
		UrlPre = "http://127.0.0.1:8200"
	} else if serverId == 1 {
		UrlPre = "https://daoserver.fushisanlang.cn"
		//panic("内测中，该服务器暂未开放。。。")
	} else if serverId == 2 {
		UrlPre = "http://43.128.67.66:8200"
		panic("内测中，该服务器暂未开放。。。")
	} else {
		panic("服务器选择错误")
	}
}
func Info() {

	generation, worldMp := getWorldStatus()
	fmt.Println("末法时代，元气消散。\n只有最优秀的人才能飞升上界，逃出生天。\n无法飞升的人只能化为虚无，归还元气给天地。\n如今，是元气复苏的" + generation + "。\n元气剩余" + worldMp + "份。\n努力修炼，争取飞升。")

}
