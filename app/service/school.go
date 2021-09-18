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

func GetSchoolList(cookie string) []string {
	urlPath := UrlPre + "/school/getschoollist"
	req, _ := http.NewRequest("GET", urlPath, nil)
	req.Header.Set("Cookie", cookie)
	resp, _ := (&http.Client{}).Do(req)

	//resp, _ := http.Get(urlPath)

	defer resp.Body.Close()

	schoolList, _ := ioutil.ReadAll(resp.Body)
	schoolListString := string(schoolList)
	SchoolList := tools.SliptJson(schoolListString)
	//fmt.Println(SchoolList)

	//var School model.SchoolList
	//var SchoolList2 []model.SchoolList
	//for i:=0;i<len(SchoolList);i++ {
	//json.Unmarshal([]byte(SchoolList[i]), &School)
	//
	//	SchoolList2=append(SchoolList2,School.SchoolInfo)
	//	fmt.Println(SchoolList2)
	//}
	return SchoolList
}
