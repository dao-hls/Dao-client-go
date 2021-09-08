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

func CliGetUser() (string, string) {
	var name, pass string
	fmt.Println("登录中：请输入姓名,输入 2 进入注册")
	fmt.Scan(&name)
	if name == "2" {
		name, pass = AddUser()
		fmt.Println("稍后将自动登录")
	} else {
		fmt.Println("登录中：请输入密码")
		fmt.Scan(&pass)
	}
	return name, pass
}
func cliGetAddUser() (string, string, string, string) {
	var name, pass, passre, email string
	fmt.Println("注册中：请输入姓名")
	fmt.Scan(&name)
	fmt.Println("注册中：请输入密码")
	fmt.Scan(&pass)
	fmt.Println("注册中：请确认密码")
	fmt.Scan(&passre)
	fmt.Println("注册中：请输入邮箱")
	fmt.Scan(&email)
	return name, pass, passre, email
}

func AddUser() (string, string) {
	var user, pass, passre, email string

	for true {
		user, pass, passre, email = cliGetAddUser()
		for true {
			if pass == passre {
				break

			} else {
				tools.Clean()
				fmt.Println("两次密码不匹配")
				user, pass, passre, email = cliGetAddUser()
			}
		}
		//urlPath := "https://dao.fushisanlang.cn/user/add"
		urlPath := UrlPre + "/user/add"
		postString := "{\"user\":\"" + user + "\",\"pass\":\"" + pass + "\",\"passre\":\"" + passre + "\",\"mail\":\"" + email + "\"}"

		postStringByte := []byte(postString)
		req, err := http.NewRequest("POST", urlPath, bytes.NewBuffer(postStringByte))

		client := &http.Client{}
		resp, _ := client.Do(req)
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
			fmt.Println("注册失败")
			fmt.Println(respMessage.Message)
			fmt.Println("请重新注册")

		} else {
			tools.Clean()
			fmt.Println(respMessage.Message)

			break
		}

	}
	return user, pass
}
