/*
@Time : 2021/8/4 9:33 上午
@Author : fushisanlang
@File : getmessage
@Software: GoLand
*/
package service

import (
	"Dao-client/app/model"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func Login(name string, pass string) (string, int) {
	//urlPath := "https://dao.fushisanlang.cn/user/login"
	urlPath := "http://127.0.0.1:8200/user/login"
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
	var respMessage model.GetMessage
	json.Unmarshal([]byte(respStr), &respMessage)

	if respMessage.Code != 200 {
		fmt.Println(404)
		return respMessage.Message, respMessage.Code

	} else {
		return resp.Header["Set-Cookie"][0], 200
	}

}
func GetRole(cookie string) string {
	//urlPath := "https://dao.fushisanlang.cn/role/get"
	urlPath := "http://127.0.0.1:8200/role/get"

	postString := "{}"
	postStringByte := []byte(postString)
	req, err := http.NewRequest("GET", urlPath, bytes.NewBuffer(postStringByte))
	req.Header.Set("Cookie", cookie)
	client := &http.Client{}
	resp, _ := client.Do(req)
	if err != nil {
		panic(err)
	}

	body, _ := ioutil.ReadAll(resp.Body)

	message := string(body)
	//fmt.Println(message)
	defer resp.Body.Close()
	return message
}
func AddUser(name string, pass string, passre string, email string, cookie string) string {

	//urlPath := "https://dao.fushisanlang.cn/user/add"
	urlPath := "http://127.0.0.1:8200/user/add"
	postString := "{\"name\":\"" + name + "\",\"pass\":\"" + pass + "\",\"passre\":\"" + passre + "\",\"mail\":\"" + email + "\"}"
	fmt.Println(postString)
	postStringByte := []byte(postString)
	req, err := http.NewRequest("POST", urlPath, bytes.NewBuffer(postStringByte))
	req.Header.Set("Cookie", cookie)
	client := &http.Client{}
	resp, _ := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	message := string(body)
	return message
}
