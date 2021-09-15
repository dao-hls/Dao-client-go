/*
@Time : 2021/9/7 14:48
@Author : zhangyin
@File : getRoleStableStatus
@Software: GoLand
*/
package tools

import (
	"math/rand"
	"time"
)

func GetRoleStableStatus() (int, int, int, int, int) {
	i := 20
	//time.Sleep(1 * time.Second)
	rand.Seed(time.Now().Unix())

	//fmt.Println(time.Now().Unix())
	a := rand.Intn(i)
	b := rand.Intn(i - a)
	c := rand.Intn(i - a - b)
	d := rand.Intn(i - a - b - c)
	e := i - a - b - c - d
	return a, b, c, d, e
}
