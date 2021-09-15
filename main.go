/*
@Time : 2021/8/3 4:27 下午
@Author : fushisanlang
@File : main.go
@Software: GoLand
*/
package main

import (
	"Dao-client/app/service"
	"Dao-client/app/tools"
	"fmt"
	"time"
)

func init() {

	service.SelectServerAdress()

	service.VersionDiff()
	service.Info()
	//
	var loginCode int
	//
	for true {
		fmt.Println("输入 1 登录，输入 2 注册 ")

		fmt.Scan(&loginCode)
		tools.Clean()
		//fmt.Println(loginCode)

		if loginCode == 1 {
			tools.Clean()
			break
		} else if loginCode == 2 {
			service.AddUser()

		} else {
			tools.Clean()
			fmt.Println("请勿输入 1 2 之外的内容")
		}
	}
}

func main() {

	name, pass := service.CliGetUser()
	cookie, code := service.Login(name, pass)

	i := 0
	for i < 1 {
		if code != 200 {

			name, pass = service.CliGetUser()
			cookie, code = service.Login(name, pass)
		} else {
			break
		}
		//
	}

	go service.PrintInfo(cookie)
	for true {
		time.Sleep(time.Duration(1) * time.Second)
	}

}
