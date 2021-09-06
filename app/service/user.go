package service

import "fmt"

func CliGetUser() (string, string) {
	var name, pass string
	fmt.Println("输入姓名")
	fmt.Scan(&name)
	fmt.Println("输入密码")
	fmt.Scan(&pass)
	return name, pass
}
