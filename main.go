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

//func main(){
//	service.Info()
//	var  name,pass,passre,email string
//	fmt.Println("输入姓名")
//	fmt.Scan(&name)
//	fmt.Println("输入密码")
//	fmt.Scan(&pass)
//	fmt.Println("重复密码")
//	fmt.Scan(&passre)
//	fmt.Println("输入邮箱")
//	fmt.Scan(&email)
//	message := service.GetMessage(name,pass,passre,email)
//	fmt.Println(message)
//}
func main() {
	service.Info()
	var name, pass string
	fmt.Println("输入姓名")
	fmt.Scan(&name)
	fmt.Println("输入密码")
	fmt.Scan(&pass)
	cookie := service.Login(name, pass)
	fmt.Println(cookie)
	i := 0
	for i < 1 {
		role := service.GetRole(cookie)
		fmt.Println(role)
		time.Sleep(time.Duration(1) * time.Second)
	}

}
