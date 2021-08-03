/*
@Time : 2021/8/3 4:27 下午
@Author : fushisanlang
@File : main.go
@Software: GoLand
*/
package main

import (
	"Dao-client/tools"
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

func getMessage(name string,pass string ,email string) string {
	//cookie := getCookie()
	//loginServer(cookie)
	urlPath := "http://127.0.0.1:8200/user/add"
	postString := "{\"name\":\""+ name + "\",\"pass\":\"" + pass + "\",\"mail\":\"" +email +"\"}"
	fmt.Println(postString)
	postStringByte := []byte(postString)
	req, err := http.NewRequest("POST", urlPath, bytes.NewBuffer(postStringByte))
	req.Header.Set("Host", "oasys.e-nci.com")
	req.Header.Set("Content-Length", "310")
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	//req.Header.Set("Cookie", cookie)
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
func main(){
	tools.Info()
	var  name,pass,email string
	fmt.Println("输入姓名")
	fmt.Scan(&name)
	fmt.Println("输入密码")
	fmt.Scan(&pass)
	fmt.Println("输入邮箱")
	fmt.Scan(&email)
	message := getMessage(name,pass,email)
	fmt.Println(message)
}