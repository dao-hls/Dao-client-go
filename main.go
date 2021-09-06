/*
@Time : 2021/8/3 4:27 下午
@Author : fushisanlang
@File : main.go
@Software: GoLand
*/
package main

import (
	"Dao-client/app/service"
	"fmt"
	"time"
)

func main() {
	service.Info()

	var loginCode int

	fmt.Println("输入 1 登录")
	fmt.Println("输入 2 注册")
	fmt.Scan(&loginCode)
	fmt.Println(loginCode)
	//service.Clean()
	if loginCode == 1 {
		name, pass := service.CliGetUser()
		cookie, code := service.Login(name, pass)
		i := 0
		for i < 1 {
			if code != 200 {
				fmt.Println(cookie)
				name, pass = service.CliGetUser()
				cookie, code = service.Login(name, pass)
			} else {
				break
			}

		}
		for i < 1 {
			role := service.GetRole(cookie)
			fmt.Println(role)
			time.Sleep(time.Duration(1) * time.Second)
		}
	} else if loginCode == 2 {
		fmt.Println("注册功能马上开启。。。")
	} else {
		fmt.Println("请勿输入 1 2 之外的内容")
	}


}
