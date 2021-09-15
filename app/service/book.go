/*
@Time : 2021/9/15 15:53
@Author : zhangyin
@File : book
@Software: GoLand
*/
package service

import (
	"Dao-client/app/model"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetBook(cookie string) {
	urlPath := UrlPre + "/book/getbook"

	req, _ := http.NewRequest("Get", urlPath, nil)
	req.Header.Set("Cookie", cookie)
	resp, _ := (&http.Client{}).Do(req)

	defer resp.Body.Close()
	bookList, _ := ioutil.ReadAll(resp.Body)
	respStr := string(bookList)

	var respMessage model.RespBooks
	json.Unmarshal([]byte(respStr), &respMessage)

	if respMessage.Code == 200 {
		len := respMessage.Len
		books := respMessage.Books[0:len]
		readBook(books, cookie)

	} else {
		fmt.Printf(respMessage.Books[0].Name)

	}
}

func readBook(bookList []model.Book, cookie string) {

	lenBookList := len(bookList)

	for i := 0; i < lenBookList; i++ {
		fmt.Printf("%v. ", i+1)
		fmt.Println(bookList[i].Name)

	}

	fmt.Println("请输入编号以修习相应功法：")
	var code int
	fmt.Scan(&code)
	fmt.Println(bookList[code-1].Name) //todo
}
