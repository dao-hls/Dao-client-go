/*
@Time : 2021/9/9 11:06
@Author : zhangyin
@File : print
@Software: GoLand
*/
package service

import (
	"Dao-client/app/tools"
	"fmt"
	"time"
)

func PrintInfo(cookie string) {

	for true {

		fmt.Println("1.获取角色信息")
		fmt.Println("2.读书")
		fmt.Println("3.奇遇")
		fmt.Println("0.退出")

		var opCode int
		fmt.Scan(&opCode)
		tools.Clean()
		if opCode == 1 {
			role, status := GetRole(cookie)
			tools.Clean()

			if status == true {
				PrintRole(role)
			}
		} else if opCode == 2 {
			ReadBook(cookie)

		} else if opCode == 3 {
			FinishtEvent(cookie)

		} else if opCode == 0 {
			fmt.Println("客户端将在10秒后关闭")
			time.Sleep(time.Duration(10) * time.Second)
			panic("bye")
		} else {
			fmt.Println("输入错误。。。")
		}
	}
}
