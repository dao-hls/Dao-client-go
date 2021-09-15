/*
@Time : 2021/8/4 9:33 上午
@Author : fushisanlang
@File : getmessage
@Software: GoLand
*/
package service

import (
	"Dao-client/app/model"
	"Dao-client/app/tools"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func Login(name string, pass string) (string, int) {
	//urlPath := "https://dao.fushisanlang.cn/user/login"
	urlPath := UrlPre + "/user/login"
	postString := "{\"name\":\"" + name + "\",\"pass\":\"" + pass + "\"}"
	postStringByte := []byte(postString)
	req, err := http.NewRequest("POST", urlPath, bytes.NewBuffer(postStringByte))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	respStr := string(body)
	var respMessage model.RespMessage
	json.Unmarshal([]byte(respStr), &respMessage)

	if respMessage.Code != 200 {
		tools.Clean()
		fmt.Println("登陆失败")
		fmt.Println(respMessage.Message)
		fmt.Println("请重新登录")
		return respMessage.Message, respMessage.Code

	} else {
		return resp.Header["Set-Cookie"][0], 200
	}

}
