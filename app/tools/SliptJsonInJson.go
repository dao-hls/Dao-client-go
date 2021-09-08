/*
@Time : 2021/9/7 22:07
@Author : zhangyin
@File : SliptJsonInJson
@Software: GoLand
*/
package tools

import (
	"regexp"
	"strings"
)

func SliptJson(NeedSliptString string) []string {
	listFirst := strings.Split(NeedSliptString, "[")[1]
	listSecond := strings.Split(listFirst, "]")[0]

	reg := regexp.MustCompile(`},{`)
	listThird := reg.ReplaceAllString(listSecond, `}},{{`)
	List := strings.Split(listThird, "},{")
	return List
}
