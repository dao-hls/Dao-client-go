/*
@Time : 2021/9/15 15:53
@Author : zhangyin
@File : book
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

func ReadBook(cookie string) {
	urlPath := UrlPre + "/book/readbook"

	req, _ := http.NewRequest("GET", urlPath, nil)
	req.Header.Set("Cookie", cookie)
	resp, _ := (&http.Client{}).Do(req)

	defer resp.Body.Close()
	bookList, _ := ioutil.ReadAll(resp.Body)
	respStr := string(bookList)

	var respMessage model.RespReading
	json.Unmarshal([]byte(respStr), &respMessage)
	if respMessage.Code == 404 {
		GetBook(cookie, 1)

	} else {
		//在读
		tools.Clean()
		if respMessage.BookName != "" {

			fmt.Println("您正在修炼" + respMessage.BookName + "中。")
			fmt.Println("1. 继续修炼。")
			fmt.Println("2. 更换心法，进度不保存。")
			fmt.Println("请输入序号选择 继续修炼 还是 更换心法：")
			var a int
			fmt.Scan(&a)
			if a == 1 {
				readBook(cookie, respMessage.BookId)
			} else if a == 2 {
				GetBook(cookie, 2)
			} else {
				tools.Clean()
				fmt.Println("输入错误，继续修炼当前心法。")
				readBook(cookie, respMessage.BookId)
			}
		} else {

			GetBook(cookie, 2)
		}
	}
}

func GetBook(cookie string, status int) {
	urlPath := UrlPre + "/book/readnewbook"

	req, _ := http.NewRequest("GET", urlPath, nil)
	req.Header.Set("Cookie", cookie)
	resp, _ := (&http.Client{}).Do(req)

	defer resp.Body.Close()
	bookList, _ := ioutil.ReadAll(resp.Body)
	respStr := string(bookList)

	var respMessage model.RespBooks
	json.Unmarshal([]byte(respStr), &respMessage)

	if respMessage.Code == 200 { //开启新书
		lenNum := respMessage.Len

		books := respMessage.Books[0:lenNum]

		var i int
		for i = 0; i < lenNum; i++ {
			fmt.Printf("%v. ", i+1)
			fmt.Println(books[i].Name)

		}

		fmt.Println("请输入编号以修习相应功法：")
		var code int
		fmt.Scan(&code)
		if i > 0 || i < lenNum {
			setReading(cookie, books[code-1].Id, status)
			readBook(cookie, books[code-1].Id)
		} else {
			fmt.Println("输入错误，开始修炼" + books[0].Name)
			readBook(cookie, books[0].Id)

		}

	} else {
		fmt.Printf(respMessage.Books[0].Name)

	}
}

func setReading(cookie string, bid int, status int) {
	var urlPath string
	if status == 1 {
		urlPath = UrlPre + "/book/setreading"
	} else if status == 2 {
		urlPath = UrlPre + "/book/setnewreading"
	}
	var setReading model.SetReading
	setReading.Id = bid
	postString, _ := json.Marshal(setReading)
	//fmt.Println(postString)
	req, err := http.NewRequest("POST", urlPath, bytes.NewBuffer(postString))
	req.Header.Set("Cookie", cookie)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	respStr := string(body)
	//fmt.Println(respStr)
	var respMessage model.RespMessage
	json.Unmarshal([]byte(respStr), &respMessage)
	tools.Clean()

	fmt.Println(respMessage.Message)

}

func readBook(cookie string, bid int) {

	code := GetReadInfo(cookie)
	for code != 405 && code != 0 {

		//tools.Clean()

		fmt.Println("1.开始修炼")
		fmt.Println("2.退出修炼")
		var a int
		fmt.Scan(&a)
		if a == 1 {
			code = GetReadInfo(cookie)

		} else if a == 2 {
			fmt.Println("退出")
			break
		} else {
			fmt.Println("输入错误，退出中")
			break
		}
	}
}

func GetReadInfo(cookie string) int {
	pass := getBookInfo()
	if pass > 18 {
		return 0
	} else {
		urlPath := UrlPre + "/book/bookinfo"
		var BookPass model.BookPass
		BookPass.Pass = pass
		postString, _ := json.Marshal(setReading)
		//fmt.Println(postString)
		req, err := http.NewRequest("POST", urlPath, bytes.NewBuffer(postString))
		req.Header.Set("Cookie", cookie)
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

		body, _ := ioutil.ReadAll(resp.Body)

		respStr := string(body)
		//fmt.Println(respStr)
		var respMessage model.RespMessage
		json.Unmarshal([]byte(respStr), &respMessage)
		tools.Clean()

		fmt.Println(respMessage.Message)
		return respMessage.Code
	}
}

func getBookInfo() int {
	var JingMai = [8]string{"任脉", "督脉", "冲脉", "带脉", "阴跷脉", "阳跷脉", "阴维脉", "阳维脉"}

	fmt.Println("修炼到关键时刻，一股磅礴精气自丹田而生，可借此冲击奇经八脉。\n此时正式突破的好时机，你想借此冲击哪条经脉？")
	for i := 0; i < 8; i++ {
		fmt.Printf("%v. ", i+1)
		fmt.Println(JingMai[i])

	}
	fmt.Println("输入其他则退出修炼。")
	var a, b int
	fmt.Scan(&a)
	if a < 1 || a > 8 {

		fmt.Println("退出修炼")
		a = 100
		b = 100
	} else {
		fmt.Println("开始冲击" + JingMai[a-1])
		a = a - 1

		tools.Clean()
		fmt.Println("天地有阴阳，静脉有起始，功法有顺逆。\n你想正向冲击经脉还是反向冲击静脉？")
		fmt.Println("1. 顺练")
		fmt.Println("2. 逆练")
		fmt.Println("输入其他则退出修炼。")
		fmt.Scan(&b)

		if b == 1 {
			fmt.Println("开始顺练")
			b = 0
		} else if b == 2 {
			fmt.Println("开始逆练")
			b = 1
		} else {
			fmt.Println("退出修炼。")

			b = 100
		}
	}
	return 8*b + a
}
