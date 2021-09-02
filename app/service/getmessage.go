/*
@Time : 2021/8/4 9:33 上午
@Author : fushisanlang
@File : getmessage
@Software: GoLand
*/
package service

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

func Login(name string, pass string) string {
	urlPath := "https://dao.fushisanlang.cn/user/login"
	postString := "{\"name\":\"" + name + "\",\"pass\":\"" + pass + "\"}"
	postStringByte := []byte(postString)
	req, err := http.NewRequest("POST", urlPath, bytes.NewBuffer(postStringByte))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	return resp.Header["Set-Cookie"][0]

}
func GetRole(cookie string) string {
	urlPath := "https://dao.fushisanlang.cn/role/get"
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

	urlPath := "https://dao.fushisanlang.cn/user/add"
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
