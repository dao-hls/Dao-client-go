/*
@Time : 2021/9/7 15:43
@Author : zhangyin
@File : school
@Software: GoLand
*/
package service

import (
	"Dao-client/app/tools"
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetSkillList() []string {
	urlPath := UrlPre + "/skill/getskilllist"
	resp, _ := http.Get(urlPath)

	defer resp.Body.Close()

	schoolList, _ := ioutil.ReadAll(resp.Body)
	schoolListString := string(schoolList)
	SchoolList := tools.SliptJson(schoolListString)
	fmt.Println("--------------------")
	fmt.Println(SchoolList)
	fmt.Println("--------------------")
	return SchoolList
}
