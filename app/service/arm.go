/*
@Time : 2021/9/7 15:43
@Author : zhangyin
@File : school
@Software: GoLand
*/
package service

import (
	"Dao-client/app/tools"
	"io/ioutil"
	"net/http"
)

func GetArmList(cookie string) []string {
	urlPath := UrlPre + "/arm/getarmlist"

	req, _ := http.NewRequest("Get", urlPath, nil)
	req.Header.Set("Cookie", cookie)
	resp, _ := (&http.Client{}).Do(req)

	//resp, _ := http.Get(urlPath)

	defer resp.Body.Close()
	schoolList, _ := ioutil.ReadAll(resp.Body)
	schoolListString := string(schoolList)
	SchoolList := tools.SliptJson(schoolListString)

	return SchoolList
}
