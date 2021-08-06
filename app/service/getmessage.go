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
	req.Header.Set("Host", "oasys.e-nci.com")
	req.Header.Set("Content-Length", "102")

	req.Header.Set("Cache-Control", "max-age=0")
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	req.Header.Set("Accept-Encoding", "gzip, deflate")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	return resp.Header["Set-Cookie"][0]

}
func GetRole(cookie string) {
	urlPath := "https://dao.fushisanlang.cn/role/get"
	postString := "{}"
	postStringByte := []byte(postString)
	req, err := http.NewRequest("GET", urlPath, bytes.NewBuffer(postStringByte))
	req.Header.Set("Host", "oasys.e-nci.com")
	req.Header.Set("Content-Length", "310")
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	req.Header.Set("Cookie", cookie)
	client := &http.Client{}
	resp, _ := client.Do(req)
	if err != nil {
		panic(err)
	}

	body, _ := ioutil.ReadAll(resp.Body)

	message := string(body)
	fmt.Println(message)
	defer resp.Body.Close()
}
func AddUser(name string, pass string, passre string, email string, cookie string) string {

	urlPath := "https://dao.fushisanlang.cn/user/add"
	postString := "{\"name\":\"" + name + "\",\"pass\":\"" + pass + "\",\"passre\":\"" + passre + "\",\"mail\":\"" + email + "\"}"
	fmt.Println(postString)
	postStringByte := []byte(postString)
	req, err := http.NewRequest("POST", urlPath, bytes.NewBuffer(postStringByte))
	req.Header.Set("Host", "oasys.e-nci.com")
	req.Header.Set("Content-Length", "310")
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
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
