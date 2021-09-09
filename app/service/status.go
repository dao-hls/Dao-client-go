/*
@Time : 2021/9/7 9:11
@Author : zhangyin
@File : status
@Software: GoLand
*/
package service

import (
	"Dao-client/app/model"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func getServiceVersion() string {

	urlPath := UrlPre + "/status/version"
	resp, _ := http.Get(urlPath)
	defer resp.Body.Close()

	version, _ := ioutil.ReadAll(resp.Body)

	respStr := string(version)
	var serverVersion model.Version
	json.Unmarshal([]byte(respStr), &serverVersion)

	return serverVersion.Version
}
func getWorldStatus() (string, string) {
	urlPath := UrlPre + "/status/worldstatus"
	resp, _ := http.Get(urlPath)
	defer resp.Body.Close()

	version, _ := ioutil.ReadAll(resp.Body)

	respStr := string(version)
	var serverVersion model.WorldStatus
	json.Unmarshal([]byte(respStr), &serverVersion)

	return serverVersion.Generation, serverVersion.WorldMp
}
func VersionDiff() {
	clientVersion := "0.0.2"

	serverVersion := getServiceVersion()

	if clientVersion != serverVersion {
		fmt.Println("客户端版本异常，请重新下载客户端。\n客户端下载链接为： https://dao.fushisanlang.cn/download.html ")
		fmt.Println("客户端将在20秒后自动关闭")
		time.Sleep(20 * time.Second)
		panic("bye")

	}

}
